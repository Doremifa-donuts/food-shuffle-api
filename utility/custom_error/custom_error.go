package custom_error

import (
	"fmt"
)

/*
	カスタムエラーとは
	アクセス権限がないエンドポイントにアクセスした、必要なパラメータが存在しない、パスワードが一致しないなどのエラーを自身で定義するもの
	HTTPのリクエストが正常に処理されたかどうかを判断するために使用される
	Serviceで定義されたエラーを発生させ、Controllerでエラーハンドリングを行う
*/

// カスタムエラー構造体
type CustomError struct {
	statusCode int
	message    string
}

// カスタムエラーを作成する関数
func NewError(code int, message string) *CustomError {
	// error.logにエラーを記録する

	return &CustomError{code, message}
}

// エラーメッセージを取得するメソッド
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error Code: %d, Message: %s", e.statusCode, e.message)
}
func (e *CustomError) StatusCode() int {
	return e.statusCode
}
