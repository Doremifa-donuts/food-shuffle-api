package logging

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var errorLog *zap.Logger

func InitLogging() error {
	// エラーログを初期化する
	errorLogFile, err := os.OpenFile("log/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil { // エラーログを開けなかった場合
		fmt.Println("Error opening error.log", err)
		return err
	}

	errorLog, err = zap.NewProduction()
	if err != nil {
		return err
	}

	errorLog = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder, // ログレベルを小文字で
			EncodeTime:     zapcore.ISO8601TimeEncoder,    // 時間をISO8601形式で
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder, // 短いファイルパスと行数
		}),
		zapcore.Lock(errorLogFile),
		zap.ErrorLevel, // ログのレベルをエラーに設定
	))

	return nil
}

// エラーログを簡潔に記録するヘルパー関数
func LogError(message string, err error) {
	fmt.Println(message, err)
	errorLog.Error(message, zap.Error(err))
}
