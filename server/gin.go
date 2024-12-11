package server

import (
	"fmt"
	logging "food-shuffle-api/log"
	"io"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitGin() (*gin.Engine, error) {

	// 画像を格納するディレクトリを作成する
	if err := os.MkdirAll("public/images", 0755); err != nil {
		logging.LogError("Error creating images directory", err)
		return nil, err
	}

	// ginのインスタンスを作成
	engine := gin.New()

	// ginのログの書き込み先であるaccess.logを開く
	AccessLogFile, err := os.OpenFile("log/access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil { // アクセスログを開けなかった場合
		// エラーをログに書き込む
		logging.LogError("Error opening access.log", err)
		return nil, err
	}

	engine.Use(gin.Recovery(), gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/metrics"},
		Output:    io.MultiWriter(AccessLogFile, os.Stdout),
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - [%s] \"%s %s %d %s\" \"%s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC3339),
				param.Method,
				param.Path,
				param.StatusCode,
				param.Latency,
				param.ErrorMessage,
			)
		},
	}))

	// CORSを許可する
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins:     true, // 許可するオリジン
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE",}, // 許可するHTTPメソッド
		AllowHeaders:     []string{"Content-Type", "Authorization"}, // 許可するヘッダー
		AllowCredentials: true, // Cookieの使用を許可
	}))
	

	// マルチパートフォームが利用できるメモリの制限を設定する(デフォルトは 32 MiB)
	engine.MaxMultipartMemory = 8 << 20 // 8 MiB

	// 接続確認用のwebページを読み込む
	checkConnectionRoute(engine)

	// APIサーバのルーティングを読み込む
	routing(engine)
	
	return engine, nil
}
