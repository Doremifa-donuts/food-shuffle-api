package model

import (
	"time"
)

type GroupSharedReviews struct {
	GroupUuid  string    `gorm:"type:char(36);foreignkey:GroupUuid;primary_key;"`
	ReviewUuid string    `gorm:"type:char(36);foreignkey:ReviewUuid;primary_key;"`
	CreatedAt  time.Time `gorm:"not null"`
}
