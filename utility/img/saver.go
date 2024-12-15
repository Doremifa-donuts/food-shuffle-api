package img

import (
	"fmt"
	"food-shuffle-api/utility/custom_error"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	logging "food-shuffle-api/log"

	"github.com/google/uuid"
)

// 画像を保存する
func SaveImages(dirPath string, images []*multipart.FileHeader) ([]string, error) {
	// 画像パスのスライスを宣言
	var imagesPath []string

	//TODO: 画像の保存が途中で失敗した場合にロールバックする

	// ディレクトリが存在しない場合は作成する
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		logging.LogError("failed to create dir", err)
		return nil, err
	}

	// 画像を保存	//TODO: 画像を軽量なフォーマットに変換する
	for _, image := range images {
		// 保存するファイルを開く
		src, err := image.Open()
		if err != nil {
			logging.LogError("colud not open image file", err)
			return nil, err
		}
		defer src.Close() // 終了後リソースを破棄

		// 画像が適切なものであるかを確認
		// ファイルのMIMEタイプを取得
		buffer := make([]byte, 512) // バッファを作成	最初の数バイトの情報でmimeTypeは確認できるので512で足りる
		_, err = src.Read(buffer)   // ファイルを読み込む
		if err != nil {
			return nil, err
		}
		// バッファの情報からmimeTypeを取得
		mimeType := http.DetectContentType(buffer)

		// ファイルシークの位置をリセット バッファに読み込んだ続きから読み込む状態になっており、正常に保存するために読み込みの開始位置を元に戻しておく
		_, err = src.Seek(0, 0)
		if err != nil {
			logging.LogError("could not seek to start of file", err)
			return nil, err
		}

		// MIMEタイプが画像のものであるかを確認
		if !(mimeType == "image/jpeg" || mimeType == "image/png" || mimeType == "image/gif") {
			logging.LogError("mime type is not image type", nil)
			return nil, custom_error.NewError(http.StatusUnsupportedMediaType, "mime type is not image type")
		}

		// mimeTypeから拡張子を取得
		ext, err := mime.ExtensionsByType(mimeType)
		if err != nil {
			logging.LogError("could not get ext for mime type", err)
			return nil, err
		}

		// ファイルパスを生成
		uuid, err := uuid.NewV7()
		if err != nil {
			logging.LogError("failed generate image file path uuid", err)
		}

		// 保存するファイル名を生成
		uniqueFileName := fmt.Sprintf("%s%s", uuid.String(), ext[0])

		// 保存されるパスを生成
		savePath := filepath.Join(dirPath, uniqueFileName)

		// 保存するファイルを作成
		dst, err := os.Create(savePath)
		if err != nil {
			logging.LogError("Error saving image:", err)
			return nil, err
		}
		defer dst.Close() // 終了後リソースを破棄

		// 保存する
		_, err = io.Copy(dst, src)
		if err != nil {
			return nil, err
		}

		// ファイル名をスライスに追加
		imagesPath = append(imagesPath, uniqueFileName)
	}
	return imagesPath, nil
}
