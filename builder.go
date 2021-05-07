package zlogger

import (
	"fmt"
	"github.com/cloa-io/zlogger/observer"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

type ZLogBuilder struct {
	option  Options
	encoder zapcore.Encoder
	obs     *observer.Observer
}

func NewBuilder() ZLogBuilder {
	return ZLogBuilder{option: Options{}, obs: observer.NewObserver()}
}

func (z ZLogBuilder) Level(level string) ZLogBuilder {
	z.option.LogLevel = strings.ToUpper(level)
	return z
}

func (z ZLogBuilder) FileLogOn(isOn bool) ZLogBuilder {
	z.option.FileLogOn = isOn
	return z
}

func (z ZLogBuilder) LogPath(logPath string) ZLogBuilder {
	z.option.LogPath = logPath
	return z
}

func (z ZLogBuilder) RotationLayout(layout string) ZLogBuilder {
	z.option.RotationLayout = layout
	return z
}

func (z ZLogBuilder) RotationLimit(limit uint) ZLogBuilder {
	z.option.RotationLimit = limit
	return z
}

func (z ZLogBuilder) RotationCycle(cycle time.Duration) ZLogBuilder {
	z.option.RotationCycle = cycle
	return z
}

func (z ZLogBuilder) Encoder(encoder zapcore.Encoder) ZLogBuilder {
	z.encoder = encoder
	return z
}

func (z ZLogBuilder) AddObserver(obs observer.LogObserver) ZLogBuilder {
	z.obs.Add(obs)
	return z
}

func (z ZLogBuilder) DisableConsole(disable bool) ZLogBuilder {
	z.option.DisableConsole = disable
	return z
}

func (z ZLogBuilder) Build() *ZLogger {
	z.option.init()

	if z.encoder == nil {
		z.encoder = newJSONEncoder()
	}

	rotator := getRotator(z)
	syncer := getWriteSyncer(rotator, z.option.DisableConsole)

	core := zapcore.NewCore(z.encoder, syncer, getLogLevel(z.option.LogLevel))
	logger := zap.New(core)

	return &ZLogger{
		obs: z.obs,
		zap: logger.Sugar(),
	}
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.InfoLevel
	case "WARN":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	case "FATAL":
		return zapcore.FatalLevel
	}

	panic(fmt.Sprintf("unknown log level [%s] found", level))
}

func getWriteSyncer(rotator *rotatelogs.RotateLogs, disableConsole bool) zapcore.WriteSyncer {
	if nil == rotator && disableConsole {
		panic("console & file all options are off")
	}

	if !disableConsole && rotator != nil {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(rotator), os.Stdout)
	} else if disableConsole && rotator != nil {
		return zapcore.AddSync(rotator)
	} else {
		return os.Stdout
	}
}

func getRotator(z ZLogBuilder) *rotatelogs.RotateLogs {
	if !z.option.FileLogOn {
		return nil
	}

	layout := z.option.RotationLayout
	limit := z.option.RotationLimit
	cycle := z.option.RotationCycle

	return newRotator(z.option.LogPath, layout, limit, cycle)
}

func newRotator(logFile, layout string, rotationLimit uint, cycle time.Duration) *rotatelogs.RotateLogs {
	rotator, err := rotatelogs.New(
		logFile+"."+layout,
		rotatelogs.WithRotationTime(cycle),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithRotationCount(rotationLimit),
		rotatelogs.WithLinkName(logFile))

	if err != nil {
		panic(err)
	}
	return rotator
}

func newJSONEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "ts",
		NameKey:        "name",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}
