package custom_error

import (
	"fmt"
	"strings"
)

/*
	カスタムエラーとは
	アクセス権限がないエンドポイントにアクセスした、必要なパラメータが存在しない、パスワードが一致しないなどのエラーを自身で定義するもの
	HTTPのリクエストが正常に処理されたかどうかを判断するために使用される
	Serviceで定義されたエラーを発生させ、Controllerでエラーハンドリングを行う
*/

// カスタムエラーのエラーコードの型を定義
type Code int

// カスタムエラー構造体
type CustomError struct {
	code    Code
	message string
}

// カスタムエラーの一覧を定義
const ( // カスタムエラーの名前を追加していく
	UncategorizedError Code = iota // 未分類のエラー
	UnauthorizedError              // 認証を失敗したエラー
)

// 各エラーに対応するエラーメッセージを定義
var errorMessages = map[Code]string{
	UncategorizedError: "An unexpected error has occurred",
	UnauthorizedError:  "failed to authorize",
}

// カスタムエラーを作成する関数
func NewError(code Code, messages ...string) *CustomError {
	// エラーメッセージのコードを引数に実行されると、デフォルトメッセージと共にエラーを返す
	var message string
	if len(messages) > 0 {
		// 配列を引数に渡された場合,で区切り、文字列として結合する
		message = strings.Join(messages, ", ")
	} else {
		// 配列を引数に渡されなかった場合,デフォルトメッセージを使用する
		message = errorMessages[code]
	}
	return &CustomError{
		code:    code,
		message: message,
	}
}

// エラーメッセージを取得するメソッド
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error Code: %d, Message: %s", e.code, e.message)
}

// エラーコードを取得するメソッド
func (e *CustomError) Code() Code {
	return e.code
}
