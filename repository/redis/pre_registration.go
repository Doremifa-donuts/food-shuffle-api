package redis

import (
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

// 1時間の間仮登録ユーザーのユーザー情報を保存する
func CachePreRegistrationUser(key string, value []byte) error {
	r := *Redis
	// 1時間だけ仮登録データを保存する
	err := r.client.Set(r.ctx, key, value, 1*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

// ユーザーの仮登録データを取得する
func GetPreRegistrationUser(key string) ([]byte, error) {

	r := *Redis
	// ユーザーデータを取得
	value, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		// キーに一致する値が見つからなかった場合
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	return []byte(value), nil
}
