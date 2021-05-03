package observer

import "github.com/getsentry/sentry-go"

type SentryObserver struct {
	Address string
	Level   int
}

func NewSentryObserver(address string, level int) *SentryObserver {
	return &SentryObserver{Address: address, Level: level}
}

func (s SentryObserver) Initialize() {

	err := sentry.Init(sentry.ClientOptions{
		Dsn: s.Address,
	})
	if err != nil {
		panic("sentry initialize fail cause=" + err.Error())
	}
}

func (s SentryObserver) LogLevel() int {
	return s.Level
}

func (s SentryObserver) Process(log string) {
	sentry.CaptureMessage(log)
}
