package model

type ShareSettingReview struct {
	UserUuid        string  `gorm:"type:char(36);foreignKey:UserUuid;primary_key"`
	FirstReviewUuid *string `gorm:"type:char(36)"`
	// SecondReviewUuid *string `gorm:"type:char(36)"`
	// ThirdReviewUuid  *string `gorm:"type:char(36)"`
	FirstReview *Review `gorm:"foreignKey:FirstReviewUuid"`
	// SecondReview     *Review `gorm:"foreignKey:SecondReviewUuid"`
	// ThirdReview      *Review `gorm:"foreignKey:ThirdReviewUuid"`
}

var sampleShareSettingReview = &[]ShareSettingReview{
	{
		UserUuid:        "a0adb027-0f54-4c1a-9ed3-86041c863344",
		FirstReviewUuid: stringPointer("39f93b17-c378-46f4-b55e-0c65642d99b0"),
		// SecondReviewUuid: stringPointer("e08505ac-eb06-43ea-a29b-b206367f7b8d"),
		// ThirdReviewUuid:  stringPointer("573fa1e4-1510-4eaf-9f1f-9df903bbd020"),
	},
}
