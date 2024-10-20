package logging

import (
	"fmt"
	"food-shuffle-api/utility/custom_error"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var errorLog *zap.Logger

func InitLogging() {
	// エラーログを初期化する
	errorLogFile, err := os.OpenFile("log/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil { // エラーログを開けなかった場合
		fmt.Println("Error opening error.log", err)
	}

	errorLog, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}

	errorLog = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(errorLogFile),
		zap.ErrorLevel, // ログのレベルをエラーに設定
	))
}

// エラーログを簡潔に記録するヘルパー関数
func LogError(message string, err error) {
	errorLog.Error(message, zap.Error(custom_error.NewError(custom_error.UncategorizedError)))
}
