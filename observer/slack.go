package observer

import (
	"fmt"
	"github.com/slack-go/slack"
)

type SlackWebHookObserver struct {
	WebHookUrl string
	LogLevel   int
}

func (s SlackWebHookObserver) Level() int {
	return s.LogLevel
}

func NewSlackWebHookObserver(webHookUrl string, Level int) SlackWebHookObserver {
	return SlackWebHookObserver{WebHookUrl: webHookUrl, LogLevel: Level}
}

func (s SlackWebHookObserver) Initialize() {
}

func (s SlackWebHookObserver) Process(log string) {
	message := &slack.WebhookMessage{
		Text: log,
	}
	err := slack.PostWebhook(s.WebHookUrl, message)
	if nil != err {
		fmt.Println(err)
	}
}
