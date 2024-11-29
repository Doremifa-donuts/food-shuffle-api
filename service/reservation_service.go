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

// 　予約詳細のレスポンスの構造体
type ReservationDetail struct {
	ReservationDate     time.Time // 予約日
	NumberOfPeople      int       // 予約人数
	UserName            string    // ユーザー名
	CourseName          *string   // コース名
	CampaignDescription *string   // キャンペーン
	DiscountOffer       *string   // 割引
	ReservationStatus   bool      // 予約が通ってるかどうかのフラグ
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

// 　予約詳細のレスポンスの構造体
func (s ReservationService) GetReservationDetailByReservation(uuid string, reservation_uuid string) (ReservationDetail, error) {
	// 返り値を定義
	var reservationDetailResponse ReservationDetail

	// トランザクションを開始する
	err := repository.Transaction(func(tx *gorm.DB) error {
		// 予約UUIDから予約を取得する
		reservation, err := repository.GetReservationByReservationUuid(tx, uuid, reservation_uuid)
		if err != nil {
			return err
		}

		// 予約情報を格納する
		reservationDetailResponse.ReservationDate = reservation.ReservationDate
		reservationDetailResponse.NumberOfPeople = reservation.NumberOfPeople
		reservationDetailResponse.ReservationStatus = reservation.ReservationStatus

		if reservation.CampaignUuid != nil {
			// キャンペーンUUIDからdescriptionを取得する
			description, err := repository.GetDescriptionByCampaignUuid(tx, *reservation.CampaignUuid)
			if err != nil {
				return err
			}
			// descriptionを格納する
			reservationDetailResponse.CampaignDescription = &description

			discountOffer, err := repository.GetDiscountOfferByCampaignUuid(tx, *reservation.CampaignUuid)
			if err != nil {
				return err
			}
			// discountを格納する
			reservationDetailResponse.DiscountOffer = &discountOffer
		}

		// 予約されたユーザーUUIDからユーザー情報を取得する
		user, err := repository.GetGeneralUserByUserUuid(tx, reservation.UserUuid)
		if err != nil {
			return err
		}

		// ユーザー名を格納する
		reservationDetailResponse.UserName = user.UserName

		if reservation.CourseUuid != nil {

			// コースUUIDからコース情報を取得する
			CourseName, err := repository.GetCourseNameByCourseUuid(tx, *reservation.CourseUuid)
			if err != nil {
				return err
			}

			// コース名を格納する
			reservationDetailResponse.CourseName = &CourseName
		}

		return nil
	})

	// 返り値を返す
	return reservationDetailResponse, err
}

func (s ReservationService) ApproveReservation(uuid string, reservation_uuid string) error {
	// トランザクションを開始する
	err := repository.Transaction(func(tx *gorm.DB) error {
		// 予約UUIDから予約を取得する
		reservation, err := repository.GetReservationByReservationUuid(tx, uuid, reservation_uuid)
		if err != nil {
			return err
		}
		// 予約を承認する
		reservation.ReservationStatus = true
		return repository.UpdateReservationStatus(tx, uuid, reservation.ReservationUuid, reservation.ReservationStatus)
	})

	// 返り値を返す
	return err
}
