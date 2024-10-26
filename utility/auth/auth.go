package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"gorm.io/gorm"

	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/repository"
	"food-shuffle-api/utility/custom_error"
)

// ログイン時にjwtトークンを発行する関数　生成したトークンは返し、トークンの検証に必要なjtiはデータベースに格納する
func GenerateToken(user *model.User) (string, error) {
	// トークンの一意性を確保するために、ランダムなUUIDを生成 時間でソートする必要はないためUUIDv4を採用
	jtiToken, err := uuid.NewRandom()
	if err != nil {
		logging.LogError("failed to generate jti token", err)
		return "", err
	}

	// トークンを文字列型にキャスト
	jtiTokenString := jtiToken.String()

	// 対象ユーザーのjtiトークンを更新する
	if err := repository.SaveJtiByUserUuid(repository.GetDB(), user.UserUuid, jtiTokenString); err != nil {
		logging.LogError("failed to save jti", err)
		return "", err
	}

	// クレームに使用するjtiとUUIDを設定
	claims := jwt.MapClaims{
		"jti": jtiTokenString,
		"id":  user.UserUuid,
		"exp": time.Now().Add(time.Second * time.Duration(Expire)).Unix(), // トークンの有効期限を1年に設定
	}

	// JWTトークンを生成、シークレットキーを使用して署名
	jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SecretKey))
	if err != nil {
		logging.LogError("failed to generate token", err)
		return "", err
	}

	return jwtToken, nil
}

// JWTトークンを検証する関数
func ValidateToken(tokenString string) (string, error) {

	// Bearerプレフィックスをチェック
	if len(tokenString) > 7 && strings.HasPrefix(tokenString, "Bearer ") {
		// プレフィックスを外してトークンを取得
		tokenString = tokenString[7:]
	} else {	// Bearerプレフィックスがなければエラーを返す
		logging.LogError("token doesn't have bearer prefix", nil)
		return "", custom_error.NewError(custom_error.ResourceNotFoundError)
	}

	// 受信したトークンを検証する
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logging.LogError("unexpected signing method", nil)
			return nil, custom_error.NewError(custom_error.UnauthorizedError)
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		logging.LogError("failed to parse token", err)
		return "", err
	}

	// トークンをアサーション
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		logging.LogError("invalid token", nil)
		return "", custom_error.NewError(custom_error.AssertionFailedError)
	}

	// トークンからjtiを取得
	jti, ok := claims["jti"].(string)
	if !ok {
		logging.LogError("jti not found in token", nil)
		return "", custom_error.NewError(custom_error.UnauthorizedError)
	} else {

		// jtiを検証
		// クレームのuuidとjtiの組み合わせを検証
		uuid := claims["id"].(string)

		// トークンから得られたuuidとjtiの組み合わせを検証
		err := repository.CheckJtiUser(repository.GetDB(), uuid, jti)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {	// 一致したレコードが存在しなかった場合
				logging.LogError("pair of uuid and jti not found", nil)
				return "", custom_error.NewError(custom_error.UnauthorizedError)
			}
			return "", err
		}

		// クレームからexpを取得
		exp, ok := claims["exp"].(float64)
		if !ok {
			logging.LogError("exp not found in token", nil)
			return "", custom_error.NewError(custom_error.AssertionFailedError)
		}
		// Unixtimeを日時に変換
		expTime := time.Unix(int64(exp), 0)

		// トークンが有効期限内かを検証
		if time.Now().After(expTime) {
			logging.LogError("token expired", nil)
			return "",  custom_error.NewError(custom_error.TokenExpiredError)
		}

		// チェックを通過したユーザーのUUIDを返す
		return uuid, nil
	}
}
