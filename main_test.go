package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/gorilla/mux"
)

func Test_server(t *testing.T) {

	testapi := os.Getenv("TESTAPIPATH")

	if testapi == "" {
		t.Fatal("Env `TESTAPIPATH` must be set (example: /services/XXX/XXX/XXXXXX)")
	}

	towerSecret = "weixinhost-test"

	router := mux.NewRouter()
	router.HandleFunc("/services/{appid}/{firstid}/{secondid}", tower).Methods("POST")

	server := httptest.NewServer(router)
	defer server.Close()

	// 创建 httpexpect 实例
	e := httpexpect.New(t, server.URL)

	// 测试api是否工作

	messageList := []map[string]interface{}{
		map[string]interface{}{

			"X-Tower-Event":     "todolists",
			"X-Tower-Signature": "weixinhost-test",

			"message": map[string]interface{}{

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
			},
		},
		map[string]interface{}{

			"X-Tower-Event":     "todolists",
			"X-Tower-Signature": "weixinhost-test",

			"message": map[string]interface{}{

				"action": "commented",
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
					"comment": map[string]interface{}{
						"guid":    "863d3b8560244638aab37b86763d83f1",
						"content": "单元测试",
					},
				},
			},
		},
		map[string]interface{}{

			"X-Tower-Event":     "topics",
			"X-Tower-Signature": "weixinhost-test",

			"message": map[string]interface{}{

				"action": "created",
				"data": map[string]interface{}{
					"project": map[string]interface{}{
						"guid": "fd06e021daba4fe294f6372075e478ee",
						"name": "单元测试",
					},
					"topic": map[string]interface{}{
						"guid":       "4e9d48368be447829422a0c87ef45ace",
						"title":      "单元测试",
						"updated_at": "2018-05-22T07:57:05.000Z",
						"handler": map[string]interface{}{
							"guid":     "a4fd9aa255f744b0a736f5f3ec379154",
							"nickname": "李博",
						},
					},
				},
			},
		},
		map[string]interface{}{

			"X-Tower-Event":     "check_items",
			"X-Tower-Signature": "weixinhost-test",

			"message": map[string]interface{}{

				"action": "created",
				"data": map[string]interface{}{
					"project": map[string]interface{}{
						"guid": "fd06e021daba4fe294f6372075e478ee",
						"name": "单元测试",
					},
					"todos::checkitem": map[string]interface{}{
						"guid":       "4e9d48368be447829422a0c87ef45ace",
						"title":      "单元测试",
						"updated_at": "2018-05-22T07:57:05.000Z",
						"handler": map[string]interface{}{
							"guid":     "a4fd9aa255f744b0a736f5f3ec379154",
							"nickname": "李博",
						},
						"assignee": map[string]interface{}{
							"guid":     "88677e863488476a8f6a5b6ddb906cec",
							"nickname": "李博",
						},
					},

					"todo": map[string]interface{}{
						"guid":  "e4030c2116f947ca97855dc822bfaf09",
						"title": "单元测试",
					},
				},
			},
		},
	}

	for _, message := range messageList {

		e.POST(testapi).WithHeader("X-Tower-Event", message["X-Tower-Event"].(string)).WithHeader("X-Tower-Signature", message["X-Tower-Signature"].(string)).WithJSON(message["message"]).Expect().Status(http.StatusOK)
	}
}
