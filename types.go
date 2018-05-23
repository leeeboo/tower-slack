package main

import (
	"fmt"
)

type Message struct {
	Action string  `json:"action"`
	Data   Payload `json:"data"`
}

type Payload struct {
	Project       Project    `json:"project"`
	TodoList      TodoList   `json:"todolist"`
	Todo          Todo       `json:"todo"`
	TodoCheckitem Todo       `json:"todos::checkitem"`
	Comment       Comment    `json:"comment"`
	Topic         Topic      `json:"topic"`
	Document      Document   `json:"document"`
	Attachment    Attachment `json:"attachment"`
}

type Project struct {
	Guid string `json:"guid"`
	Name string `json:"name"`
}

type TodoList struct {
	Guid      string `json:"guid"`
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
	Handler   Person `json:"handler"`
}

type Todo struct {
	Guid      string `json:"guid"`
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
	Handler   Person `json:"handler"`
	DueAt     string `json:"due_at"`
	Assignee  Person `json:"assignee"`
}

type Comment struct {
	Guid    string `json:"guid"`
	Content string `json:"content"`
}

type Topic struct {
	Guid      string `json:"guid"`
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
	Handler   Person `json:"handler"`
}

type Document struct {
	Guid      string `json:"guid"`
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
	Handler   Person `json:"handler"`
}

type Attachment struct {
	Guid      string `json:"guid"`
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
	Handler   Person `json:"handler"`
}

type Person struct {
	Guid     string `json:"guid"`
	Nickname string `json:"nickname"`
}

func (this *Message) ToSlackMessage(event string) string {

	actionList, ok := actionToString[event]

	if !ok {
		return ""
	}

	actionFmt, ok := actionList[this.Action]

	if !ok {
		return ""
	}

	var message string

	switch event {
	case "todolists":

		switch this.Action {
		case "commented":
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.TodoList.Handler.Nickname, this.Data.TodoList.Title, this.Data.Comment.Content)
			break
		default:
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.TodoList.Handler.Nickname, this.Data.TodoList.Title)
		}

		break
	case "todos":

		switch this.Action {
		case "assigned":
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.TodoList.Title, this.Data.Todo.Handler.Nickname, this.Data.Todo.Title, this.Data.Todo.Assignee.Nickname)
			break
		case "unassigned":
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.TodoList.Title, this.Data.Todo.Handler.Nickname, this.Data.Todo.Assignee.Nickname, this.Data.Todo.Title)
			break
		case "deadline_changed":
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.TodoList.Title, this.Data.Todo.Handler.Nickname, this.Data.Todo.Title, this.Data.Todo.DueAt)
			break
		case "moved":
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.TodoList.Title, this.Data.Todo.Handler.Nickname, this.Data.Todo.Title, this.Data.TodoList.Title)
			break
		case "commented":
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.TodoList.Title, this.Data.Todo.Handler.Nickname, this.Data.Todo.Title, this.Data.Comment.Content)
			break
		default:
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.TodoList.Title, this.Data.Todo.Handler.Nickname, this.Data.Todo.Title)
		}

		break
	case "check_items":

		switch this.Action {
		case "assigned":
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.TodoCheckitem.Handler.Nickname, this.Data.Todo.Title, this.Data.TodoCheckitem.Title, this.Data.TodoCheckitem.Assignee.Nickname)
			break
		case "unassigned":
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.TodoCheckitem.Handler.Nickname, this.Data.Todo.Title, this.Data.TodoCheckitem.Assignee.Nickname, this.Data.TodoCheckitem.Title)
			break
		case "deadline_changed":
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.TodoCheckitem.Handler.Nickname, this.Data.Todo.Title, this.Data.TodoCheckitem.Title, this.Data.TodoCheckitem.DueAt)
			break
		default:
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.TodoCheckitem.Handler.Nickname, this.Data.Todo.Title, this.Data.TodoCheckitem.Title)
		}

		break
	case "topics":

		switch this.Action {
		case "commented":
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.Topic.Handler.Nickname, this.Data.Topic.Title, this.Data.Comment.Content)
			break
		default:
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.Topic.Handler.Nickname, this.Data.Topic.Title)
			break
		}

		break
	case "documents":

		switch this.Action {
		case "commented":
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.Document.Handler.Nickname, this.Data.Document.Title, this.Data.Comment.Content)
			break
		default:
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.Document.Handler.Nickname, this.Data.Document.Title)
			break
		}

		break
	case "attachments":

		switch this.Action {
		case "commented":
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.Attachment.Handler.Nickname, this.Data.Attachment.Title, this.Data.Comment.Content)
			break
		default:
			message = fmt.Sprintf(actionFmt, this.Data.Project.Name, this.Data.Attachment.Handler.Nickname, this.Data.Attachment.Title)
			break
		}

		break
	}

	return message
}
