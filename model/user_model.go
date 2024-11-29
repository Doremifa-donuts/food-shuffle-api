package model

type UserType string

const (
	General    UserType = "General"
	Restaurant UserType = "Restaurant"
)

type User struct {
	UserUuid    string          `gorm:"type:char(36);primary_key;"`                  // ユーザーのUUID
	MailAddress string          `gorm:"type:varchar(255);not null"`                  // メールアドレス
	Password    string          `gorm:"type:varchar(255);not null"`                  // ハッシュ化されたパスワード
	Tell        string          `gorm:"type:varchar(20);not null"`                   // 電話番号
	JtiToken    string          `gorm:"type:char(36);not null"`                      // jtiトークン　JWTの解析に使う
	UserType    UserType        `gorm:"type:enum('General', 'Restaurant');not null"` // ユーザーの種類　一般利用者、レストラン利用者
	User        *GeneralUser    `gorm:"foreignKey:UserUuid"`
	Restaurant  *RestaurantUser `gorm:"foreignKey:RestaurantUuid;references:UserUuid"`
}

// サンプルデータ
var sampleUsers = []User{
	{
		UserUuid:    "a0adb027-0f54-4c1a-9ed3-86041c863344",
		MailAddress: "poster@test.com",
		Password:    "$2a$10$UkrQfUmAsPJ35cw5TVzJeOuoLySOWpMHN/b2zN561eixU0abBSCpq", // test
		Tell:        "08012341234",
		UserType:    General,
	},
	{
		UserUuid:    "91a78381-f472-496b-90e3-2c66a33391d1",
		MailAddress: "viewer@test.com",
		Password:    "$2a$10$UkrQfUmAsPJ35cw5TVzJeOuoLySOWpMHN/b2zN561eixU0abBSCpq", // test
		Tell:        "08012341234",
		UserType:    General,
	},
	{
		UserUuid:    "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		MailAddress: "restaurant@test.com",
		Password:    "$2a$10$UkrQfUmAsPJ35cw5TVzJeOuoLySOWpMHN/b2zN561eixU0abBSCpq", // test
		Tell:        "08056785678",
		UserType:    Restaurant,
	},
	{
		UserUuid:    "a80499ae-eb6c-1305-a5cc-e1510c52744a",
		MailAddress: "resto@test.com",
		Password:    "$2a$10$UkrQfUmAsPJ35cw5TVzJeOuoLySOWpMHN/b2zN561eixU0abBSCpq", // test
		Tell:        "08012341234",
		UserType:    Restaurant,
	},
}
