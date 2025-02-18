package zap

import (
	contrib "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func NewLumberjack(c *Config) (*lumberjack.Logger, func()) {
	_conf := c.File
	if !_conf.Enable {
		return nil, func() {}
	}
	w := &lumberjack.Logger{
		Filename:   _conf.Filename,
		MaxSize:    int(_conf.Maxage), // megabytes
		MaxBackups: int(_conf.Maxbackups),
		MaxAge:     int(_conf.Maxage), //days
		Compress:   _conf.Compress,    // disabled by default
	}
	return w, func() {
		_ = w.Close()
	}
}

func NewZap(_conf *Config, ll *lumberjack.Logger) (*zap.Logger, func(), error) {
	level, err := zap.ParseAtomicLevel(_conf.Level)
	if err != nil {
		return nil, nil, err
	}
	enc := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "ts",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
	syncers := make([]zapcore.WriteSyncer, 0, 2)
	if _conf.Stdout.Enable {
		syncers = append(syncers, zapcore.AddSync(os.Stdout))
	}
	if _conf.File.Enable {
		syncers = append(syncers, &zapcore.BufferedWriteSyncer{
			WS: zapcore.AddSync(ll),
		})
	}

	core := zapcore.NewCore(
		enc,
		zapcore.NewMultiWriteSyncer(syncers...),
		level,
	)
	stacktrace := zap.AddStacktrace(zapcore.ErrorLevel)
	if level.Level() == zapcore.DebugLevel {
		stacktrace = zap.AddStacktrace(zapcore.WarnLevel)
	}
	z := zap.New(core,
		zap.AddCaller(),
		stacktrace,
		zap.ErrorOutput(zapcore.AddSync(os.Stderr)),
		zap.AddCallerSkip(3),
	)
	return z, func() {
		_ = z.Sync()
	}, nil
}

type KV []any

func NewLogger(kv KV, logger *zap.Logger) log.Logger {
	l := log.With(contrib.NewLogger(logger), kv...)
	helper := log.NewHelper(l)
	helper.Info("日志启动")
	return l
}

func NewKV(c *Config) KV {
	_conf := c.KV
	return []any{
		"service.id", _conf.ID,
		"service.name", _conf.Name,
		"service.version", _conf.Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	}
}
