package log

import (
	"net/url"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZaprLogger(cfg Config) (logr.Logger, error) {
	zapLogger, err := newZapLogger(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "new logger failed")
	}

	logger := zapr.NewLogger(zapLogger)

	return logger, nil
}

func NewZapLogger(cfg Config) (*zap.Logger, error) {
	return newZapLogger(cfg)
}

func newZapLogger(cfg Config) (*zap.Logger, error) {
	cfgOutputPaths := []string{}
	if runtime.GOOS == "windows" {
		zap.RegisterSink("winfile", func(u *url.URL) (zap.Sink, error) {
			return os.OpenFile(u.Path[1:], os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		})
		for _, outputPath := range cfg.OutputPaths {
			cfgOutputPaths = append(cfgOutputPaths, func(cfgRootDir string) string {
				if cfgRootDir == "stdout" {
					return filepath.ToSlash(filepath.Join(cfgRootDir, ""))
				}

				return "winfile:///" + filepath.ToSlash(filepath.Join(cfgRootDir, ""))
			}(outputPath))
		}
	} else {
		cfgOutputPaths = cfg.OutputPaths
	}

	var zapLogLevel zapcore.Level
	zapLogLevel.UnmarshalText([]byte(cfg.OutputLevel))
	zapAtomicLevel := zap.NewAtomicLevelAt(zapLogLevel)

	zapCfg := zap.Config{
		Level:         zapAtomicLevel,
		OutputPaths:   cfgOutputPaths,
		Encoding:      "console",
		DisableCaller: true,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "message",
			LevelKey:      "level",
			TimeKey:       "timestamp",
			NameKey:       "name",
			CallerKey:     "caller",
			StacktraceKey: "stacktrace",
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeName:    zapcore.FullNameEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}

	zapLogger, err := zapCfg.Build()
	if err != nil {
		return nil, errors.Errorf("new zap logger failed: %s", err.Error())
	}
	defer zapLogger.Sync()

	return zapLogger, nil
}
