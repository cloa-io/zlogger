package zlogger

import (
	"fmt"
	"github.com/cloa-io/zlogger/observer"
	"go.uber.org/zap"
)

type ZLogger struct {
	obs *observer.Observer
	zap *zap.SugaredLogger
}

func (s ZLogger) Zap() *zap.Logger {
	return zap.L()
}

func (s *ZLogger) Sync() error {
	return s.zap.Sync()
}

func (s *ZLogger) Debug(args ...interface{}) {
	s.zap.Debug(args...)
	if s.obs != nil {
		s.obs.Notify(observer.DEBUG, printf(args...))
	}
}

func printf(args ...interface{}) string {
	if args != nil && len(args) >= 1 {
		s := args[0]
		return fmt.Sprintf("%s", s)
	}
	return fmt.Sprintf("%v", args...)
}

func (s *ZLogger) Info(args ...interface{}) {
	s.zap.Info(args...)
	if s.obs != nil {
		s.obs.Notify(observer.INFO, printf(args...))
	}
}

func (s *ZLogger) Warn(args ...interface{}) {
	s.zap.Warn(args...)
	if s.obs != nil {
		s.obs.Notify(observer.WARN, printf(args...))
	}
}

func (s *ZLogger) Error(args ...interface{}) {
	s.zap.Error(args...)
	if s.obs != nil {
		s.obs.Notify(observer.ERROR, printf(args...))
	}
}

func (s *ZLogger) Fatal(args ...interface{}) {
	s.zap.Fatal(args...)
	if s.obs != nil {
		s.obs.Notify(observer.FATAL, printf(args...))
	}
}

func (s *ZLogger) Debugf(template string, args ...interface{}) {
	s.zap.Debugf(template, args...)
	if s.obs != nil {
		s.obs.Notify(observer.DEBUG, printf(args...))
	}
}

func (s *ZLogger) Infof(template string, args ...interface{}) {
	s.zap.Infof(template, args...)
	if s.obs != nil {
		s.obs.Notify(observer.INFO, printf(args...))
	}
}

func (s *ZLogger) Warnf(template string, args ...interface{}) {
	s.zap.Warnf(template, args...)
	if s.obs != nil {
		s.obs.Notify(observer.WARN, printf(args...))
	}
}

func (s *ZLogger) Errorf(template string, args ...interface{}) {
	s.zap.Errorf(template, args...)
	if s.obs != nil {
		s.obs.Notify(observer.ERROR, printf(args...))
	}
}

func (s *ZLogger) Fatalf(template string, args ...interface{}) {
	s.zap.Fatalf(template, args...)
	if s.obs != nil {
		s.obs.Notify(observer.FATAL, printf(args...))
	}
}
