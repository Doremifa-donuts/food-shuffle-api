package service

import (
	"food-shuffle-api/dto/response"
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

func (s ReservationService) GetReservationsByRestaurant(uuid string) ([]response.ReservationsByRestaurant, error) {
	// 返り値を定義
	var ReservationListResponses []response.ReservationsByRestaurant

	// トランザクションを開始する
	err := repository.Transaction(func(tx *gorm.DB) error {
		// レストランUUIDから予約を取得する
		reservations, err := repository.GetReservationsByRestaurantUuid(tx, uuid)
		if err != nil {
			return err
		}

		//
		for _, reservation := range reservations {

			// 予約情報を格納する
			ReservationListResponse := response.ReservationsByRestaurant{
				ReservationUuid:   reservation.ReservationUuid,
				ReservationDate:   reservation.ReservationDate,
				NumberOfPeople:    reservation.NumberOfPeople,
				ReservationStatus: reservation.ReservationStatus,
			}

			// 予約されたユーザーUUIDからユーザー情報を取得する
			user, err := repository.GetGeneralUserByUserUuid(tx, reservation.UserUuid)
			if err != nil {
				return err
			}

			// ユーザー名を格納する
			ReservationListResponse.UserName = user.UserName

			// 予約情報を格納する
			ReservationListResponses = append(ReservationListResponses, ReservationListResponse)
		}

		return nil
	})

	// 返り値を返す
	return ReservationListResponses, err
}
