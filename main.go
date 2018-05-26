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
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var towerSecret string
var port string

var interrupt = make(chan os.Signal, 1)
var shutdown = make(chan struct{}, 1)

func init() {

	flag.StringVar(&towerSecret, "secret", "", "The secret of your tower webhook.")
	flag.StringVar(&port, "port", "8080", "Port to listen (default 8080).")

	flag.Parse()

	towerSecret = strings.TrimSpace(towerSecret)

	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
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

	slackMessage := message.ToSlackMessage(event)

	if slackMessage == nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	apiUrl := fmt.Sprintf("https://hooks.slack.com%s", r.URL.Path)

	err = sendToSlack(apiUrl, *slackMessage)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func sendToSlack(api string, slackMsg SlackMessage) error {

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
	body, _ := ioutil.ReadAll(resp.Body)

	if string(body) != "ok" {
		return errors.New(fmt.Sprintf("Slack api response error `%s`", string(body)))
	}

	return nil
}

func server() *http.Server {

	r := mux.NewRouter()
	r.HandleFunc("/services/{appid}/{firstid}/{secondid}", tower).Methods("POST")

	handler := handlers.CORS()(r)
	handler = handlers.CombinedLoggingHandler(os.Stdout, handler)
	handler = gziphandler.GzipHandler(handler)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	return srv
}

func main() {

	srv := server()

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			shutdown <- struct{}{}
			log.Println(err)
		}
	}()

	select {
	case killSignal := <-interrupt:
		switch killSignal {
		case os.Interrupt:
			log.Println("Got SIGINT...")
		case syscall.SIGTERM:
			log.Println("Got SIGTERM...")
		}
	case <-shutdown:
		log.Println("Got an error...")
	}

	log.Println("The service is shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	} else {
		log.Println("Done")
	}
}
