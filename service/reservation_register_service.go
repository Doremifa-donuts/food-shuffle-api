package service

import (
	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationService struct{}

// 予約の登録を行い、トークンを返す
func (service *ReservationService) ResevationRegister(reservation model.Reservation) (string, error) {

	// レスポンスの型を初期化する
	var reservationUuid string

	// トランザクションを開始する
	err := repository.Transaction(func(tx *gorm.DB) error {

		// UUIDを生成する
		uuid, err := uuid.NewV7()
		if err != nil {
			logging.LogError("failed to generate user uuid", err)
			return err
		}

		//データを挿入する
		reservation.ReservationUuid = uuid.String()

		reservationUuid = uuid.String()

		//予約テーブルに追加情報を追加する
		err = repository.CreateReservation(tx, reservation)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	return reservationUuid, nil
}
