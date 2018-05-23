package main

import (
	"fmt"
	"time"
)

type SlackMessage struct {
	Text        string        `json:"text,omitempty"`
	Mrkdwn      bool          `json:"mrkdwn,omitempty"`
	Attachments []Attachments `json:"attachments,omitempty"`
}

type Attachments struct {
	Color      string `json:"color,omitempty"`
	Pretext    string `json:"pretext,omitempty"`
	AuthorName string `json:"author_name,omitempty"`
	AuthorLink string `json:"author_link,omitempty"`
	AuthorIcon string `json:"author_icon,omitempty"`
	Title      string `json:"title,omitempty"`
	TitleLink  string `json:"title_link,omitempty"`
	Text       string `json:"text,omitempty"`
	ThumbUrl   string `json:"thumb_url,omitempty"`
	Footer     string `json:"footer,omitempty"`
	FooterIcon string `json:"footer_icon,omitempty"`
	Ts         int64  `json:"ts,omitempty"`
}

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

func (this *Message) ToSlackMessage(event string) *SlackMessage {

	formatList, ok := eventFormat[event]

	if !ok {
		return nil
	}

	actionFmt, ok := formatList[this.Action]

	if !ok {
		return nil
	}

	var message string

	var a Attachments
	a.Title = fmt.Sprintf("%s %s", eventToString[event], actionToString[this.Action])
	a.Footer = this.Data.Project.Name
	a.Ts = time.Now().Unix()
	a.Pretext = fmt.Sprintf("Tower.im 项目通知 [%s]", this.Data.Project.Name)

	switch event {
	case "todolists":

		switch this.Action {
		case "commented":
			message = fmt.Sprintf(actionFmt, this.Data.TodoList.Title, this.Data.Comment.Content)
			break
		default:
			message = fmt.Sprintf(actionFmt, this.Data.TodoList.Title)
		}

		a.AuthorName = this.Data.TodoList.Handler.Nickname
		a.AuthorLink = fmt.Sprintf("https://tower.im/members/%s/", this.Data.TodoList.Handler.Guid)
		a.TitleLink = fmt.Sprintf("https://tower.im/projects/%s/lists/%s/show/", this.Data.Project.Guid, this.Data.TodoList.Guid)
		a.Text = message

		break
	case "todos":

		switch this.Action {
		case "assigned":
			message = fmt.Sprintf(actionFmt, this.Data.Todo.Handler.Nickname, this.Data.Todo.Title, this.Data.Todo.Assignee.Nickname)
			break
		case "unassigned":
			message = fmt.Sprintf(actionFmt, this.Data.Todo.Handler.Nickname, this.Data.Todo.Assignee.Nickname, this.Data.Todo.Title)
			break
		case "deadline_changed":
			message = fmt.Sprintf(actionFmt, this.Data.Todo.Handler.Nickname, this.Data.Todo.Title, this.Data.Todo.DueAt)
			break
		case "moved":
			message = fmt.Sprintf(actionFmt, this.Data.Todo.Handler.Nickname, this.Data.Todo.Title, this.Data.TodoList.Title)
			break
		case "commented":
			message = fmt.Sprintf(actionFmt, this.Data.Todo.Handler.Nickname, this.Data.Todo.Title, this.Data.Comment.Content)
			break
		default:
			message = fmt.Sprintf(actionFmt, this.Data.Todo.Handler.Nickname, this.Data.Todo.Title)
		}

		a.AuthorName = this.Data.Todo.Handler.Nickname
		a.AuthorLink = fmt.Sprintf("https://tower.im/members/%s/", this.Data.Todo.Handler.Guid)
		a.TitleLink = fmt.Sprintf("https://tower.im/projects/%s/todos/%s/", this.Data.Project.Guid, this.Data.Todo.Guid)
		a.Text = message

		break
	case "check_items":

		switch this.Action {
		case "assigned":
			message = fmt.Sprintf(actionFmt, this.Data.Todo.Title, this.Data.TodoCheckitem.Title, this.Data.TodoCheckitem.Assignee.Nickname)
			break
		case "unassigned":
			message = fmt.Sprintf(actionFmt, this.Data.Todo.Title, this.Data.TodoCheckitem.Assignee.Nickname, this.Data.TodoCheckitem.Title)
			break
		case "deadline_changed":
			message = fmt.Sprintf(actionFmt, this.Data.Todo.Title, this.Data.TodoCheckitem.Title, this.Data.TodoCheckitem.DueAt)
			break
		default:
			message = fmt.Sprintf(actionFmt, this.Data.Todo.Title, this.Data.TodoCheckitem.Title)
		}

		a.AuthorName = this.Data.TodoCheckitem.Handler.Nickname
		a.AuthorLink = fmt.Sprintf("https://tower.im/members/%s/", this.Data.TodoCheckitem.Handler.Guid)
		a.TitleLink = fmt.Sprintf("https://tower.im/projects/%s/todos/%s/", this.Data.Project.Guid, this.Data.Todo.Guid)
		a.Text = message

		break
	case "topics":

		switch this.Action {
		case "commented":
			message = fmt.Sprintf(actionFmt, this.Data.Topic.Title, this.Data.Comment.Content)
			break
		default:
			message = fmt.Sprintf(actionFmt, this.Data.Topic.Title)
			break
		}

		a.AuthorName = this.Data.Topic.Handler.Nickname
		a.AuthorLink = fmt.Sprintf("https://tower.im/members/%s/", this.Data.Topic.Handler.Guid)
		a.TitleLink = fmt.Sprintf("https://tower.im/projects/%s/messages/%s/", this.Data.Project.Guid, this.Data.Topic.Guid)
		a.Text = message

		break
	case "documents":

		switch this.Action {
		case "commented":
			message = fmt.Sprintf(actionFmt, this.Data.Document.Title, this.Data.Comment.Content)
			break
		default:
			message = fmt.Sprintf(actionFmt, this.Data.Document.Title)
			break
		}

		a.AuthorName = this.Data.Document.Handler.Nickname
		a.AuthorLink = fmt.Sprintf("https://tower.im/members/%s/", this.Data.Document.Handler.Guid)
		a.TitleLink = fmt.Sprintf("https://tower.im/projects/%s/docs/%s/", this.Data.Project.Guid, this.Data.Document.Guid)
		a.Text = message

		break
	case "attachments":

		switch this.Action {
		case "commented":
			message = fmt.Sprintf(actionFmt, this.Data.Attachment.Title, this.Data.Comment.Content)
			break
		default:
			message = fmt.Sprintf(actionFmt, this.Data.Attachment.Title)
			break
		}

		a.AuthorName = this.Data.Attachment.Handler.Nickname
		a.AuthorLink = fmt.Sprintf("https://tower.im/members/%s/", this.Data.Attachment.Handler.Guid)
		a.TitleLink = fmt.Sprintf("https://tower.im/projects/%s/uploads/%s/", this.Data.Project.Guid, this.Data.Attachment.Guid)
		a.Text = message

		break
	}

	var slackMessage SlackMessage

	aList := make([]Attachments, 1)
	aList = append(aList, a)

	slackMessage.Attachments = aList

	return &slackMessage
}
