package redis

import (
	"food-shuffle-api/utility/custom_error"
	"net/http"

	"github.com/redis/go-redis/v9"
)

// 位置情報をredisに登録する
func SetGeoLocation(userUuid string, latitude float64, longitude float64) error {
	r := *Redis
	key := "locations"
	err := r.client.GeoAdd(r.ctx, key, &redis.GeoLocation{
		Name:      userUuid,
		Longitude: longitude,
		Latitude:  latitude,
	}).Err()
	if err != nil {
		return err
	}
	return nil
}

// 自身からレビュー共有範囲内のユーザーのUUIDを取得
func GetUserUuidsByReviewShareRadius(userUuid string, reviewShareRadius int64) ([]string, error) {
	r := *Redis
	key := "locations"

	var userUuids []string
	// ユーザーの位置情報を取得

	userGeo := r.client.GeoPos(r.ctx, key, userUuid).Val()[0]
	if userGeo == nil {
		return userUuids, custom_error.NewError(http.StatusInternalServerError, "can not get user's geo data")
	}

	//TODO: 適切のレビュー範囲の調整をする
	results, err := r.client.GeoRadius(r.ctx, key, userGeo.Longitude, userGeo.Latitude, &redis.GeoRadiusQuery{Radius: float64(reviewShareRadius), Unit: "m"}).Result()
	if err != nil {
		return userUuids, err
	}

	for _, result := range results {
		//　自分以外のユーザーのUUIDを配列として格納する
		if result.Name != userUuid {
			userUuids = append(userUuids, result.Name)
		}
	}

	return userUuids, nil
}

// 店舗の付近にいるユーザーのリストを取得する
func GetUserUuidsByRestaurantBoostRadius(latitude float64, longitude float64, reviewShareRadius int64) ([]string, error) {
	r := *Redis
	key := "locations"

	var userUuids []string

	results, err := r.client.GeoRadius(r.ctx, key, longitude, latitude, &redis.GeoRadiusQuery{Radius: float64(reviewShareRadius), Unit: "m"}).Result()
	if err != nil {
		return userUuids, err
	}

	for _, result := range results {
		userUuids = append(userUuids, result.Name)
	}

	return userUuids, nil
}
