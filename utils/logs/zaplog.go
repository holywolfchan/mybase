// Package ginzap provides sugarlogger handling using zap package.
// Code structure based on ginrus package.
package logs

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var userlogger *zap.SugaredLogger
var accesslogger *zap.SugaredLogger

func init() {
	var err error
	encoder_cfg := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	custom_cfg1 := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    encoder_cfg,
		OutputPaths:      []string{"stdout", "./temp_log/syslog.log"},
		ErrorOutputPaths: []string{"stderr"},
	}
	custom_cfg2 := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    encoder_cfg,
		OutputPaths:      []string{"stdout", "./temp_log/userlog.log"},
		ErrorOutputPaths: []string{"stderr"},
	}
	custom_cfg3 := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    encoder_cfg,
		OutputPaths:      []string{"stdout", "./temp_log/accesslog.log"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err = custom_cfg1.Build()
	if err != nil {
		panic(err)
	}

	logger2, err2 := custom_cfg2.Build()
	if err2 != nil {
		panic(err2)
	}
	userlogger = logger2.Sugar()

	logger3, err3 := custom_cfg3.Build()
	if err3 != nil {
		panic(err3)
	}
	accesslogger = logger3.Sugar()
}

// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
//
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
//
// It receives:
//   1. A time package format string (e.g. time.RFC3339).
//   2. A boolean stating whether to use UTC time zone or local.
func Ginlog(timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		visitor := c.GetHeader("Authorization")
		if visitor == "" {
			visitor = c.DefaultQuery("token", "UNKNOWN GUEST")
		}
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				fmt.Printf("ERROR:[%s]\n", e)
				logger.Error(e)
			}
		} else {
			logger.Info(path,
				zap.Duration("latency", latency),
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("visitor", visitor),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("endtime", end.Format(timeFormat)),
			)
		}
	}
}

// Error logs a message at error level.
func Error(v ...interface{}) {
	userlogger.Error(v...)
}

// Warn compatibility alias for Warning()
func Warn(v ...interface{}) {
	userlogger.Warn(v...)
}

// Info compatibility alias for Warning()
func Info(v ...interface{}) {
	userlogger.Info(v...)
}

// Debug logs a message at debug level.
func Debug(v ...interface{}) {
	userlogger.Debug(v...)
}

// Trace logs a message at trace level.
// compatibility alias for Warning()
func Fatal(v ...interface{}) {
	userlogger.Fatal(v...)
}

func Panic(v ...interface{}) {
	userlogger.Panic(v...)
}

func Infof(template string, args ...interface{}) {
	userlogger.Infof(template, args...)
}
func Errorf(template string, args ...interface{}) {
	userlogger.Errorf(template, args...)
}
func Infow(msg string, args ...interface{}) {
	userlogger.Infow(msg, args...)
}
func Errorw(msg string, args ...interface{}) {
	userlogger.Errorw(msg, args...)
}

//Access intents to log important,sensible resource's access behavior
func Access(v ...interface{}) {
	accesslogger.Info(v...)
}
func Accessf(template string, args ...interface{}) {
	accesslogger.Infof(template, args...)
}
