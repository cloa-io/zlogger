package observer

const (
	DEBUG = iota
	WARN
	INFO
	ERROR
	FATAL
)

type Observer struct {
	obs []LogObserver
}

func NewObserver() *Observer {
	return &Observer{obs: make([]LogObserver, 0)}
}

func (l *Observer) Add(entry LogObserver) {
	l.obs = append(l.obs, entry)
}

func (l *Observer) Notify(level int, msg string) {
	for _, entry := range l.obs {
		if entry.Level() <= level {
			entry.Process(msg)
		}
	}
}

type LogObserver interface {
	Initialize()
	Level() int
	Process(log string)
}
