package model

import "time"

// グループに所属するユーザーを管理するテーブル
type GroupSubmission struct {
	GroupUuid string    `gorm:"type:char(36);foreignkey:GroupUuid;primary_key"` // グループUUID
	UserUuid  string    `gorm:"type:char(36);foreignkey:UserUuid;primary_key"`  // ユーザーUUID
	CreatedAt time.Time `gorm:"not null"`                                       // ユーザーはグループに所属した日時に基づいてグループをソートできる
}
