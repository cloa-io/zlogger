package zlogger

import (
	"github.com/cloa-io/zlogger/observer"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
	"time"
)

func TestZLogger(t *testing.T) {
	//TestLog := "./logs/zlogger.log"

	//t.Run("로그 파일 생성 테스트", func(t *testing.T) {
	//	_, err := os.Open(TestLog)
	//	if os.IsExist(err) {
	//		err := os.Remove(TestLog)
	//		assert.Nil(t, err)
	//	}
	//
	//	logger := NewBuilder().
	//		FileLogOn(true).
	//		LogPath(TestLog).
	//		RotationLayout("%Y-%m-%d-%H-%S").
	//		RotationCycle(1 * time.Second).
	//		Build()
	//
	//	now := time.Now()
	//	format := now.Format("2006-01-02-15-05")
	//
	//	logger.Infof("sexy now")
	//
	//	assert.Truef(t, Exists(TestLog), "로그 파일이 존재해야된다")
	//	assert.Truef(t, Exists(TestLog + "." + format), "로그 파일이 존재해야된다")
	//
	//	os.Remove(TestLog)
	//	os.Remove(TestLog + "." + format)
	//})

	t.Run("로그 파일 로테이션 테스트", func(t *testing.T) {

		logger := NewBuilder().
			FileLogOn(true).
			LogPath("zlogger.log").
			RotationLayout("%Y-%m-%d-%H-%S").
			RotationLimit(2).
			RotationCycle(1 * time.Second).
			Build()

		now := time.Now()
		format := now.Format("2006-01-02-15-05")
		logger.Info("hello")

		expectLogFileName := "zlogger.log." + format

		assert.True(t, Exists(expectLogFileName), "로그 파일이 존재해야된다")

		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			logger.Infof("hello=%d", i)
		}
	})

}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func TestSlack(t *testing.T) {

	slackObserver := observer.NewSlackWebHookObserver("https://hooks.slack.com/services/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", 0)
	logger := NewBuilder().AddObserver(slackObserver).Build()
	//logger := NewBuilder().Build()

	logger.Debug("희희희", 123, " ㅂㅈㄷㅂㅈㄷ")

}

func TestSimpleConsoleLog(t *testing.T) {
	logger := NewBuilder().Build()
	defer logger.Sync()
	logger.Debug("DEBUG MESSAGE")
}

func TestFileLogPrint(t *testing.T) {
	logger := NewBuilder().
		FileLogOn(true).
		LogPath("./log/zlogger_test.log").
		RotationLayout("%Y-%m-%d").
		RotationLimit(7).
		RotationCycle(24 * time.Hour).
		Build()
	defer logger.Sync()

	logger.Info("INFO MESSAGE")
}

func TestEncoderResultPrint(t *testing.T) {
	config := zapcore.EncoderConfig{
		MessageKey: "M",
		LevelKey:   "L",
		TimeKey:    "T",
		NameKey:    "N",
		// CapitalLevelEncoder serializes a Level to an all-caps string. For example,
		// InfoLevel is serialized to "INFO".
		EncodeLevel: zapcore.CapitalLevelEncoder,
		// ISO8601TimeEncoder serializes a time.Time to an ISO8601-formatted string
		// with millisecond precision.
		//
		// If enc supports AppendTimeLayout(t time.Time,layout string), it's used
		// instead of appending a pre-formatted string value.
		EncodeTime:     zapcore.EpochMillisTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     nil,
	}
	consoleEncoder := zapcore.NewJSONEncoder(config)

	logger := NewBuilder().Encoder(consoleEncoder).Build()
	logger.Info("INFO MESSAGE")
}

func TestConsoleLogPrint(t *testing.T) {
	//rawJSON := []byte(`{
	//  "level": "debug",
	//  "encoding": "json",
	//  "outputPaths": ["stdout"],
	//  "errorOutputPaths": ["stderr"],
	//  "initialFields": {"foo": "bar"},
	//  "encoderConfig": {
	//    "messageKey": "message",
	//    "levelKey": "level",
	//    "levelEncoder": "lowercase"
	//  }
	//}`)
	//
	//var cfg zap.Config
	//if err := json.Unmarshal(rawJSON, &cfg); err != nil {
	//	panic(err)
	//}
	//logger, err := cfg.Build()
	//if err != nil {
	//	panic(err)
	//}
	//defer logger.Sync()
	//
	//logger.Info("logger construction succeeded")
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(consoleEncoder, os.Stdout, lowPriority)
	logger := zap.New(core)
	logger.Debug("message")
}
