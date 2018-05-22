package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/NYTimes/gziphandler"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var towerSecret string
var port string

func init() {
	flag.StringVar(&towerSecret, "secret", "", "The secret of your tower webhook.")
	flag.StringVar(&port, "port", "8080", "Port to listen (default 8080).")

	flag.Parse()

	towerSecret = strings.TrimSpace(towerSecret)
}

func tower(w http.ResponseWriter, r *http.Request) {

	if towerSecret != "" {
		if r.Header.Get("X-Tower-Signature") != towerSecret {
			w.WriteHeader(http.StatusForbidden)
			return
		}
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var message Message

	err = json.Unmarshal(body, &message)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	event := r.Header.Get("X-Tower-Event")

	log.Println(event)
	log.Println(string(body))

	text := message.ToSlackMessage(event)

	log.Println(text)

	if text == "" {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	apiUrl := fmt.Sprintf("https://hooks.slack.com%s", r.URL.Path)

	err = sendToSlack(apiUrl, text)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func sendToSlack(api string, text string) error {

	type SlackMessage struct {
		Text string `json:"text"`
	}

	var slackMsg SlackMessage
	slackMsg.Text = text

	buf := new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	err := enc.Encode(slackMsg)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	resp, err := http.Post(api, "application/json", strings.NewReader(buf.String()))

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if string(body) != "ok" {
		return errors.New(fmt.Sprintf("Slack api response error `%s`", string(body)))
	}

	return nil
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/services/{appid}/{firstid}/{secondid}", tower).Methods("POST")

	//https://hooks.slack.com/services/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

	handler := handlers.CORS()(r)
	handler = handlers.CombinedLoggingHandler(os.Stdout, r)
	handler = gziphandler.GzipHandler(handler)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	// this channel is for graceful shutdown:
	// if we receive an error, we can send it here to notify the server to be stopped
	shutdown := make(chan struct{}, 1)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			shutdown <- struct{}{}
			logs.Info(err)
		}

	}()
	logs.Info("The service is ready to listen and serve.", port)

	select {
	case killSignal := <-interrupt:
		switch killSignal {
		case os.Interrupt:
			logs.Info("Got SIGINT...")
		case syscall.SIGTERM:
			logs.Info("Got SIGTERM...")
		}
	case <-shutdown:
		logs.Error("Got an error...")
	}

	logs.Info("The service is shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		logs.Error("HTTP server Shutdown: %v", err)
	} else {
		logs.Info("Done")
	}
}
