package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/repository"
	"food-shuffle-api/utility/custom_error"
)

// ログイン時にjwtトークンを発行する関数　生成したトークンは返し、トークンの検証に必要なjtiはデータベースに格納する
func GenerateToken(entity interface{}) (string, error) {
	// トークンの一意性を確保するために、ランダムなUUIDを生成 時間でソートする必要はないためUUIDv4を採用
	jti, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	// jtiにプレフィックスを付与
	jtiPrefix := jti.String()

	// トークンを検証するために必要なjtiをDBに格納
	// 引数の型によってJTIを保存する
	var UUID string
	switch e := entity.(type) { // 型アサーションを使用
	case *model.User:
		// UUIDを取り出す
		UUID = e.UserUuid
		// jtiにプレフィックスを付与
		jtiPrefix = "u-" + jtiPrefix
		if err := repository.SaveJtiByUserUuid(repository.GetDB(), e.UserUuid, jtiPrefix); err != nil {
			return "", err
		}
	case *model.RestoUser:
		// UUIDを取り出す
		UUID = e.RestoUuid
		// jtiにプレフィックスを付与
		jtiPrefix = "r-" + jtiPrefix
		if err := repository.SaveJtiByRestoUuid(repository.GetDB(), e.RestoUuid, jtiPrefix); err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("unsupported entity type")
	}

	fmt.Println("生成したjti", jtiPrefix)

	// クレームに使用するjtiとUUIDを設定
	claims := jwt.MapClaims{
		"jti": jtiPrefix,
		"id":  UUID,
		"exp": time.Now().Add(time.Second * time.Duration(Expire)).Unix(), // トークンの有効期限を1年に設定
	}

	// JWTトークンを生成、シークレットキーを使用して署名
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SecretKey))
}

// JWTトークンを検証する関数
func ValidateToken(tokenString string) (string, error) {

	// Bearerプレフィックスをチェック
	if len(tokenString) > 7 && strings.HasPrefix(tokenString, "Bearer ") {
		// プレフィックスを外してトークンを取得
		tokenString = tokenString[7:]
	}

	// 受信したトークンを検証する
	fmt.Println(tokenString)
	// トークンを検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// logging.LogError("unexpected signing method: %v", token.Header["alg"].)
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("署名は検証できた")

	// トークンをアサーション
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		logging.LogError("invalid token", custom_error.NewError(custom_error.UncategorizedError))
		return "", fmt.Errorf("invalid token")
	}
	fmt.Println(claims)

	// トークンからjtiを取得
	jti, ok := claims["jti"].(string)
	if !ok {
		return "", fmt.Errorf("invalid token")
	} else {

		// jtiを検証
		// クレームのUuidとjtiの組み合わせを検証
		Uuid := claims["id"].(string)
		var isMatch bool
		var err error
		switch {
		case jti[:2] == "u-":
			// id と割り当てられたjtiが一致したレコードがあるかを検証
			isMatch, err = repository.CheckJtiUser(repository.GetDB(), Uuid, jti)
		case jti[:2] == "r-":
			// id と割り当てられたjtiが一致したレコードがあるかを検証
			isMatch, err = repository.CheckJtiResto(repository.GetDB(), Uuid, jti)
		default:
			return "", fmt.Errorf("invalid token")
		}

		if err != nil {
			logging.LogError("Error checking jti", err)
			return "", err
		} else if !isMatch {
			return "", fmt.Errorf("invalid token")
		}

		// トークンが有効期限内かを検証
		exp, ok := claims["exp"].(float64)
		if !ok {
			return "", fmt.Errorf("invalid token")
		}
		// Unixtimeを日時に変換
		expTime := time.Unix(int64(exp), 0)

		if time.Now().After(expTime) {
			return "", fmt.Errorf("token expired")
		}

		// チェックを通過したユーザーのUUIDを返す
		return Uuid, nil
	}
}
