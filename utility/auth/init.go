package auth

import (
	"errors"
	logging "food-shuffle-api/log"
	"os"
	"strconv"
)

// JWTトークンに使用するパラメータを定義
var (
	SecretKey string
	Expire    int
)

// jwtトークンに使用するパラメータを
func InitAuth() error {
	// 環境変数からJWTシークレットキーを取得
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		logging.LogError("JWT_SECRET_KEY is not set", nil)
		return errors.New("JWT_SECRET_KEY is not set")
	}
	// トークンの有効期限を取得
	expiration, err := strconv.Atoi(os.Getenv("JWT_TOKEN_LIFETIME"))
	if err != nil || expiration == 0 {
		logging.LogError("JWT_TOKEN_LIFETIME is not set", nil)
		return errors.New("JWT_TOKEN_LIFETIME is not set")
	}

	// JWTトークンに使用するパラメータを定義
	SecretKey = secretKey
	Expire = expiration

	return nil
}
