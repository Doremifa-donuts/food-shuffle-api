package model

import "time"

// グループに所属するユーザーを管理するテーブル
type PopupGroupSubmission struct {
	PopupGroupUuid string    `gorm:"type:char(36);primary_key"` // グループUUID
	UserUuid       string    `gorm:"type:char(36);primary_key"`       // ユーザーUUID
	CreatedAt      time.Time `gorm:"not null"`                                            // ユーザーはグループに所属した日時に基づいてグループをソートできる
	PopupGroup     []PopupGroup	`gorm:"foreignKey:PopupGroupUuid;references:PopupGroupUuid"`
	User           []User		`gorm:"foreignKey:UserUuid;references:UserUuid"`
}
