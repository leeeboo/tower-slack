package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/gorilla/mux"
)

func Test_server(t *testing.T) {

	s := server()
	if s == nil {
		t.Fatal("server init error.")
	}

	testapi := os.Getenv("TESTAPIPATH")

	if testapi == "" {
		t.Fatal("Env `TESTAPIPATH` must be set (example: /services/XXX/XXX/XXXXXX)")
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
		t.Log(event, mlist)

		for _, msg := range mlist {
			e.POST(testapi).WithHeader("X-Tower-Event", event).WithHeader("X-Tower-Signature", towerSecret).WithJSON(msg).Expect().Status(http.StatusOK)
		}
	}

	return

}
