package service

import (
	"food-shuffle-api/repository"
	"time"

	"gorm.io/gorm"
)

type ReservationService struct{}

// 予約一覧のレスポンスの構造体
type ReservationListResponse struct {
	ReservationUuid   string    // 予約UUID
	ReservationDate   time.Time // 予約日
	NumberOfPeople    int       // 予約人数
	UserName          string    // ユーザー名
	ReservationStatus bool      // 予約が通ってるかどうかのフラグ
}

func (s ReservationService) GetReservationsByRestaurant(uuid string) ([]ReservationListResponse, error) {
	// 返り値を定義
	var ReservationListResponses []ReservationListResponse

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
			ReservationListResponse := ReservationListResponse{
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
