package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/gorilla/mux"
)

var testapi string

func init() {

	testapi = os.Getenv("TESTAPIPATH")

	if testapi == "" {
		panic("Env `TESTAPIPATH` must be set (example: /services/XXX/XXX/XXXXXX)")
	}
}

func Test_server(t *testing.T) {

	s := server()
	if s == nil {
		t.Fatal("server init error.")
	}

	towerSecret = "test"

	router := mux.NewRouter()
	router.HandleFunc("/services/{appid}/{firstid}/{secondid}", tower).Methods("POST")

	server := httptest.NewServer(router)

	defer server.Close()

	// 创建 httpexpect 实例
	e := httpexpect.New(t, server.URL)

	// 测试api是否工作

	e.POST(testapi).WithHeader("X-Tower-Event", "todos").WithHeader("X-Tower-Signature", "fake").Expect().Status(http.StatusForbidden)
	e.POST(testapi).WithHeader("X-Tower-Event", "todos").WithHeader("X-Tower-Signature", towerSecret).Expect().Status(http.StatusInternalServerError)
	e.POST(testapi).WithHeader("X-Tower-Event", "todos").WithHeader("X-Tower-Signature", towerSecret).WithBytes([]byte("test")).Expect().Status(http.StatusInternalServerError)

	o := map[string]interface{}{

		"action": "created",
		"data": map[string]interface{}{
			"project": map[string]interface{}{
				"guid": "fd06e021daba4fe294f6372075e478ee",
				"name": "单元测试",
			},
			"todolist": map[string]interface{}{
				"guid":       "4e9d48368be447829422a0c87ef45ace",
				"title":      "单元测试",
				"updated_at": "2018-05-22T07:57:05.000Z",
				"handler": map[string]interface{}{
					"guid":     "a4fd9aa255f744b0a736f5f3ec379154",
					"nickname": "李博",
				},
			},
		},
	}

	e.POST(testapi).WithHeader("X-Tower-Event", "fake").WithHeader("X-Tower-Signature", towerSecret).WithJSON(o).Expect().Status(http.StatusInternalServerError)

	e.POST("/services/fake/fake/fake").WithHeader("X-Tower-Event", "todos").WithHeader("X-Tower-Signature", towerSecret).WithJSON(o).Expect().Status(http.StatusInternalServerError)

	o["action"] = "fake"

	e.POST(testapi).WithHeader("X-Tower-Event", "todos").WithHeader("X-Tower-Signature", towerSecret).WithJSON(o).Expect().Status(http.StatusInternalServerError)

	data, err := ioutil.ReadFile("./test.msg")

	if err != nil {
		t.Fatal(err)
	}

	var list map[string][]interface{}

	err = json.Unmarshal(data, &list)

	if err != nil {
		t.Fatal(err)
	}

	for event, mlist := range list {

		for _, msg := range mlist {
			e.POST(testapi).WithHeader("X-Tower-Event", event).WithHeader("X-Tower-Signature", towerSecret).WithJSON(msg).Expect().Status(http.StatusOK)
		}
	}

	return
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestTower(t *testing.T) {

	testRequest := httptest.NewRequest(http.MethodPost, testapi, errReader(0))
	testRequest.Header.Set("X-Tower-Signature", towerSecret)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(tower)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, testRequest)

	// Check the status code is what we expect.
	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("handler returned OK status: got %v want %v", rr.Code, http.StatusInternalServerError)
	}

	// Check the response body is what we expect.
}

func TestSendToSlack(t *testing.T) {

	var slackMsg SlackMessage
	slackMsg.Text = "test"

	err := sendToSlack("foo", slackMsg)

	if err == nil {
		t.Fatal("handler returned OK status: got OK want error")
	}
}
