package zlogger

import "time"

type Options struct {
	LogLevel       string //TRACE, DEBUG, WARN, INFO, ERROR, FATAL
	FileLogOn      bool
	LogPath        string
	RotationLayout string
	RotationLimit  uint
	RotationCycle  time.Duration
}

func (o *Options) init() {
	if o.LogLevel == "" {
		o.LogLevel = "DEBUG"
	}

	if o.FileLogOn {
		if o.LogPath == "" {
			o.LogPath = "DEBUG.log"
		}

		if o.RotationLimit == 0 {
			o.RotationLimit = 7
		}

		if o.RotationLayout == "" {
			o.RotationLayout = "%Y-%m-%d"
		}

		if o.RotationCycle == 0 {
			o.RotationCycle = 24 * time.Hour
		}
	}
}
