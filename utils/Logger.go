package utils

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugar = zap.NewExample().Sugar()

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = zapcore.DebugLevel
	// InfoLevel is the default logging priority.
	InfoLevel = zapcore.InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = zapcore.WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = zapcore.ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel = zapcore.DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel = zapcore.PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = zapcore.FatalLevel
)

func Logger(c echo.Context, ReqRes string, level zapcore.Level) {
	//pool, err := pgxpool.New(context.Background(), Registration.ConnectString)
	//if err != nil {
	//	fmt.Println(err)
	//	fmt.Println("Connection Failed")
	//}
	//
	//errPing := pool.Ping(context.Background())
	//if errPing != nil {
	//	fmt.Println(err)
	//	fmt.Println("")
	//}
}
