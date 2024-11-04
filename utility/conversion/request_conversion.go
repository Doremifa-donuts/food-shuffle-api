package conversion

import (
	"encoding/json"
	"fmt"
	"food-shuffle-api/utility/custom_error"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"reflect"

	logging "food-shuffle-api/log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ファイルサイズを1MiB以内に制限する
// const maxFileSize = 1 << 20 // 1MiB

// 保存先ディレクトリを設定する
const saveDir = "public/images"

// 最大画像数を設定する
const maxImages = 10

// 画像ファイルのMIMEタイプを確認する関数
func isImage(fileHeader *multipart.FileHeader) (bool, error) {
	// ファイルを開く
	file, err := fileHeader.Open()
	if err != nil {
		return false, err
	}
	defer file.Close()

	// ファイルのMIMEタイプを取得
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return false, err
	}
	mimeType := http.DetectContentType(buffer)

	// MIMEタイプが画像のものであるかを確認
	return mimeType == "image/jpeg" || mimeType == "image/png" || mimeType == "image/gif", nil
}

// 構造体に"Images"フィールドがあるかを確認する
func hasImagesField(review interface{}) bool {
	// 引数をリフレクションで取得
	v := reflect.ValueOf(review)

	// 引数がポインタの場合、元の値を取得
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// 構造体でない場合は false
	if v.Kind() != reflect.Struct {
		return false
	}

	// "Images"フィールドを確認
	field := v.FieldByName("Images")
	return field.IsValid() && field.Kind() == reflect.Slice
}

// "Images"フィールドに値を追加する
func addImage(review interface{}, images []string) error {
	// 引数をリフレクションで取得
	v := reflect.ValueOf(review).Elem()

	// "Images"フィールドを取得
	field := v.FieldByName("Images")

	// 値を追加する
	newImages := reflect.ValueOf(images)
	field.Set(newImages)
	return nil
}

// multipart/form-dataのリクエストの画像保存と引数として与えられた構造体に値の格納を行う
func RequestSaveImagesAndBindJSON(ctx *gin.Context, target interface{}) error {

	// multipart/form-dataであることを確認する
	if ctx.ContentType() != "multipart/form-data" {
		logging.LogError("invalid content type", nil)
		return custom_error.NewError(http.StatusBadRequest, "Invalid content type")
	}

	// 引数が適切に設定されているかを確認する
	// ここに引っかかった場合は引数に設定する構造体が間違えているまたは、構造体の定義ミスがある
	if !hasImagesField(target) {
		err := custom_error.NewError(http.StatusBadRequest, "Images field is missing or invalid")
		logging.LogError("Images field is missing or invalid:", err)
		return err
	}

	// jsonデータを取得
	jsonData := ctx.PostForm("data")

	// jsonデータが空の場合はエラーを返す
	if jsonData == "" {
		err := custom_error.NewError(http.StatusBadRequest, "Invalid JSON data: Empty")
		logging.LogError("Invalid JSON data:", err)
		return err
	} else {
		// JSONデータを構造体にパース
		if err := json.Unmarshal([]byte(jsonData), &target); err != nil {
			logging.LogError("Error parsing JSON data:", err)
			return err
		}
	}

	// multipart/form-dataを受け取る
	form, err := ctx.MultipartForm()
	if err != nil {
		logging.LogError("Error getting form:", err)
		return err
	}

	// キーがimage[]であるものを取得する
	files := form.File["images[]"]
	if len(files) > maxImages {
		err := custom_error.NewError(http.StatusRequestEntityTooLarge, "Too many images")
		logging.LogError("Too many images:", err)
		return err
	}

	// 画像ファイルのMIMEタイプを確認
	for _, file := range files {
		isImg, err := isImage(file)
		if err != nil {
			logging.LogError("Uploaded file is not a valid image:", err)
			return err
		}
		if !isImg {
			err := custom_error.NewError(http.StatusUnsupportedMediaType, "Uploaded file is not a valid image")
			logging.LogError("Uploaded file is not a valid image:", err)
			return err
		}
	}

	// 画像パスのスライスを宣言
	var images []string

	// 画像を保存	//TODO: 画像を軽量なフォーマットに変換する
	for _, file := range files {
		// UUIDベースのファイル名を生成
		ext := filepath.Ext(file.Filename)
		uniqueFileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)
		savePath := filepath.Join(saveDir, uniqueFileName)

		err = ctx.SaveUploadedFile(file, savePath)
		if err != nil {
			logging.LogError("Error saving image:", err)
			return err
		}
		fmt.Println("Image saved:", uniqueFileName)
		// ファイル名をスライスに追加
		images = append(images, uniqueFileName)
	}

	// "Images"フィールドに値を追加
	err = addImage(target, images)
	if err != nil {
		return err
	}

	return nil
}
