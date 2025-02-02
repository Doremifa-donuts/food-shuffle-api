package service

import (
	"food-shuffle-api/dto"
	logging "food-shuffle-api/log"
	"food-shuffle-api/repository/model"
	"food-shuffle-api/repository/orm"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationService struct{}

// 予約の登録を行い、トークンを返す
func (service *ReservationService) ResevationRegister(bReservation model.Reservation) (res dto.PostReservation, err error) {
	// トランザクションを開始する
	err = orm.Transaction(func(tx *gorm.DB) error {

		// UUIDを生成する
		reservationUuid, err := uuid.NewV7()
		if err != nil {
			logging.LogError("failed to generate user uuid", err)
			return err
		}
		// 作成した予約UUIDを格納する
		res.ReservationUuid = reservationUuid.String()

		//データを挿入する
		bReservation.ReservationUuid = reservationUuid.String()

		//予約テーブルに追加情報を追加する
		err = orm.CreateReservation(tx, bReservation)
		if err != nil {
			logging.LogError("failed to create user", err)
			return err
		}

		// トランザクションを終了する
		return nil
	})

	return
}

// レストランの予約リストを取得
func (s ReservationService) GetReservationsByRestaurant(uuid string) (res []dto.ReservationsByRestaurant, err error) {
	// トランザクションを開始する
	err = orm.Transaction(func(tx *gorm.DB) error {
		// レストランUUIDから予約を取得する
		reservations, err := orm.GetReservationsByRestaurantUuid(tx, uuid)
		if err != nil {
			logging.LogError("failed to get reservations", err)
			return err
		}

		// 予約情報を格納する
		for _, reservation := range reservations {

			// 予約情報を格納する
			ReservationListResponse := dto.ReservationsByRestaurant{
				ReservationUuid:   reservation.ReservationUuid,
				ReservationDate:   reservation.ReservationDate,
				NumberOfPeople:    reservation.NumberOfPeople,
				ReservationStatus: reservation.ReservationStatus,
			}

			// 予約されたユーザーUUIDからユーザー情報を取得する
			user, err := orm.GetGeneralUserByUserUuid(tx, reservation.UserUuid)
			if err != nil {
				logging.LogError("failed to get user", err)
				return err
			}

			// ユーザー名を格納する
			ReservationListResponse.UserName = user.UserName

			// 予約情報を格納する
			res = append(res, ReservationListResponse)
		}

		// トランザクションを終了する
		return nil
	})

	return
}

// ユーザーが自身の予約状況を取得する
func (s ReservationService) GetUpcomingReservation(userUuid string) (res []dto.UserReservation, err error) {
	// トランザクションを開始する
	err = orm.Transaction(func(tx *gorm.DB) error {
		// 現在自国以降の予約情報を取得する
		reservations, err := orm.ListUpcomingReservationByUserUuid(tx, userUuid)
		if err != nil {
			return err
		}

		// 予約に不足している情報を取得する
		for _, reservation := range reservations {
			// レストラン名
			restaurantName, err := orm.GetRestaurantNameByRestaurantUuid(tx, reservation.RestaurantUuid)
			if err != nil {
				return err
			}

			// TODO: 予約状況の取得内容が不足している
			// コース名と価格
			// if reservation.CourseUuid != nil {
			// 	course, err := orm.GetSpecificCourse(tx, *reservation.CourseUuid)
			// 	if err != nil {
			// 		return err
			// 	}
			// }

			// お助けブースト
			// if reservation.CampaignUuid != nil {
			// }

			res = append(res, dto.UserReservation{
				RestaurantUuid:    reservation.RestaurantUuid,
				RestaurantName:    restaurantName,
				ReservationDate:   reservation.ReservationDate,
				NumberOfPeople:    reservation.NumberOfPeople,
				ReservationStatus: reservation.ReservationStatus,
			})

		}
		return nil
	})
	return
}
