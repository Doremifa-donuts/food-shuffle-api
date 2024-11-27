package dto

import "time"

type CreateUrgentCampaign struct {
	CampaignUuid   string
}

type GetUrgentCampaigns struct {
	CampaignUuid string
	RestaurantUuid string
	StartAt        time.Time
	EndAt          time.Time
	Description    string
	DiscountOffer  string
	CreatedAt      time.Time
}