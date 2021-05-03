# ZLogger Library

Uber Zap 라이브러리를 이용한 Log Library

## Features

  - 레벨 별 로깅
  - Console Log
  - File Log
      - 로케이션 기능 지원
  - Log Formatter 
  - 레벨 별 Observer
    
## How to Use

### Installation
**Go 1.16** or later is required


### Console Log Example

```
func main() {
    logger := zlogger.NewBuilder().Build()
    defer logger.Sync() // Sync flushes any buffered log entries.
    
    logger.Debug("DEBUG MESSAGE")
}
```
```
{"level":"DEBUG","ts":"2021-05-03T10:22:37.223+0900","msg":"DEBUG MESSAGE"}
```

### File Log Example

```
func main() {
    logger := zlogger.NewBuilder().
                FileLogOn(true).
                LogPath("./log/zlogger_test.log"). //log path
                RotationLayout("%Y-%m-%d").     //e.g) ./log/zlogger_test.log.2021-05-03
                RotationLimit(7).  //로테이션 최대 갯수 지정(7개가 넘어갈 경우 자동 삭제)
                RotationCycle(24 * time.Hour). //로테이션 주기
                Build()
    defer logger.Sync()
     
    logger.Info("INFO MESSAGE")                   
}   
```
```
{"level":"INFO","ts":"2021-05-03T10:34:46.199+0900","msg":"INFO MESSAGE"}
```
```
/mnt/c/Users/user/IdeaProjects/zlogger/log » ll                                                                                                                                                                                                                                                 pak1627@LAPTOP-3FALA0UI
total 0
lrwxrwxrwx 1 pak1627 pak1627  27 May  3 10:41 zlogger_test.log -> zlogger_test.log.2021-05-03
-rwxrwxrwx 1 pak1627 pak1627 148 May  3 10:41 zlogger_test.log.2021-05-03
```

### Observer Example
```
func main() {
    hookUrl := "https://hooks.slack.com/services/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"
    DEBUG_LEVEL = 0
    
    slackObserver := observer.NewSlackWebHookObserver(hookUrl, DEBUG_LEVEL)
    logger := NewBuilder().AddObserver(slackObserver).Build()
    defer logger.Sync()
    
    logger.DEBUG("DEBUG MESSAGE")
}
```

### Encoder Example
```
func main() {
    config := zapcore.EncoderConfig{
        MessageKey:    "M",
        LevelKey:      "L",
        TimeKey:       "T",
        // CapitalLevelEncoder serializes a Level to an all-caps string. For example,
        // InfoLevel is serialized to "INFO".
        EncodeLevel: zapcore.CapitalLevelEncoder,
        // ISO8601TimeEncoder serializes a time.Time to an ISO8601-formatted string
        // with millisecond precision.
        //
        // If enc supports AppendTimeLayout(t time.Time,layout string), it's used
        // instead of appending a pre-formatted string value.
        EncodeTime:       zapcore.EpochMillisTimeEncoder,
    }
    encoder := zapcore.NewJSONEncoder(config)
    
    logger := NewBuilder().Encoder(encoder).Build()
    logger.Info("INFO MESSAGE")
}
```
```
{"L":"INFO","T":1620014473482.8857,"M":"INFO MESSAGE"}
```


### Options
zlogger 생성시 Builder를 통해 옵션 지원 
- NewBuilder().Level(): 로깅 레벨 지정
- NewBuilder().FileLogOn(): file 로그를 사용할지 여부 (기본 false)
  - NewBuilder.LogPath(): file 로그를 남길 path
  - NewBuilder.RotationLayout(): rotation layout [Strftime](https://man7.org/linux/man-pages/man3/strftime.3.html) 지원
  - NewBuilder.RotationLimit(): 로테이션 최대 갯수 지정(7개가 넘어갈 경우 자동 삭제)
  - NewBuilder.RotationCycle(): 로테이션 주기(기본:24시간)
- NewBuilder().Encoder(): 로그가 어떤 형태로 남을지 설정되는 Encoder 값 설정(기본:JSON 포맷)
- NewBuilder().AddObserver(): 레벨별 옵저버 설정