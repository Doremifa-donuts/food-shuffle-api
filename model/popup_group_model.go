package model

import "time"

type PopupGroup struct {
	PopupGroupUuid string                    `gorm:"char(36);primary_key"`            // グループのUUID
	PopupGroupName string                    `gorm:"type:varchar(50);not null"`       // グループ名
	ExpirationDate time.Time                 `gorm:"type:date;not null"`              // グループの有効期限は23：59固定にするため日付情報をみを格納する　期限が切れたグループはバッチ処理によって削除する。グループの有効期限は作成日時から一週間
	InviteCode     string                    `gorm:"type:varchar(6);not null;unique"` // グループの招待コード　トリガーをDB側で作成しておき、重複を防げるすごいシステムになる予定
	Users          []PopupGroupSubmission    `gorm:"foreignKey:PopupGroupUuid"`
	SharedReviews  []PopupGroupSharedReviews `gorm:"foreignKey:PopupGroupUuid"`
}
