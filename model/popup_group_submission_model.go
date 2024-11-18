package model

import "time"

// グループに所属するユーザーを管理するテーブル
type PopupGroupSubmission struct {
	PopupGroupUuid string    `gorm:"type:char(36);foreignKey:PopupGroupUuid;primary_key"` // グループUUID
	UserUuid       string    `gorm:"type:char(36);foreignKey:UserUuid;primary_key"`       // ユーザーUUID
	CreatedAt      time.Time `gorm:"not null"`                                            // ユーザーはグループに所属した日時に基づいてグループをソートできる
}

// サンプルデータ
var samplePopupGroupSubmissions = []PopupGroupSubmission{}
