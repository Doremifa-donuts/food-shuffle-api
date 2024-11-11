package model

import (
	"time"
)

type PopupGroupSharedReviews struct {
	PopupGroupUuid string    `gorm:"type:char(36);primary_key;"`
	ReviewUuid     string    `gorm:"type:char(36);primary_key;"`
	CreatedAt      time.Time `gorm:"not null"`
	PopupGroup     []PopupGroup	`gorm:"foreignKey:PopupGroupUuid;references:PopupGroupUuid"`
	Review         []Review	`gorm:"foreignKey:ReviewUuid;references:ReviewUuid"`
}
