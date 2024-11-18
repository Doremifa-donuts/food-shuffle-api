package model

import (
	"time"
)

type PopupGroupSharedReviews struct {
	PopupGroupUuid string    `gorm:"type:char(36);foreignKey:PopupGroupUuid;primary_key;"`
	ReviewUuid     string    `gorm:"type:char(36);foreignKey:ReviewUuid;primary_key;"`
	CreatedAt      time.Time `gorm:"not null"`
}
