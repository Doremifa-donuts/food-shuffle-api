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
		Tell:        "07012341234",
		UserType:    General,
	},
	{
		UserUuid:    "a0adb027-0f54-4c1a-9ed3-86041c863340",
		MailAddress: "poster2@test.com",
		Password:    "$2a$10$UkrQfUmAsPJ35cw5TVzJeOuoLySOWpMHN/b2zN561eixU0abBSCpq", // test
		Tell:        "07012341234",
		UserType:    General,
	},
	{
		UserUuid:    "91a78381-f472-496b-90e3-2c66a33391d1",
		MailAddress: "viewer@test.com",
		Password:    "$2a$10$UkrQfUmAsPJ35cw5TVzJeOuoLySOWpMHN/b2zN561eixU0abBSCpq", // test
		Tell:        "08012341234",
		UserType:    General,
	},

	{ // 路地裏ビストロ816
		UserUuid:    "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		MailAddress: "resto1@test.com",
		Password:    "$2a$10$UkrQfUmAsPJ35cw5TVzJeOuoLySOWpMHN/b2zN561eixU0abBSCpq", // test
		Tell:        "08056785678",
		UserType:    Restaurant,
	},
	{ // 炭焼 高田屋
		UserUuid:    "a80499ae-eb6c-1305-a5cc-e1510c52744a",
		MailAddress: "resto2@test.com",
		Password:    "$2a$10$UkrQfUmAsPJ35cw5TVzJeOuoLySOWpMHN/b2zN561eixU0abBSCpq", // test
		Tell:        "08012341234",
		UserType:    Restaurant,
	},
	{ // おにぎりごりちゃん 中崎町本店
		UserUuid:    "0bf97fc8-019e-421b-85f5-84818aab19d8",
		MailAddress: "resto3@test.com",
		Password:    "$2a$10$UkrQfUmAsPJ35cw5TVzJeOuoLySOWpMHN/b2zN561eixU0abBSCpq", // test
		Tell:        "08012341234",
		UserType:    Restaurant,
	},
	{ // Cafe de paris 大丸心斎橋店
		UserUuid:    "d61aed9f-68b0-4efd-af77-98d7e061526d",
		MailAddress: "resto4@test.com",
		Password:    "$2a$10$UkrQfUmAsPJ35cw5TVzJeOuoLySOWpMHN/b2zN561eixU0abBSCpq", // test
		Tell:        "08012341234",
		UserType:    Restaurant,
	},
	{ // かごのや
		UserUuid:    "6d7c3625-a1fa-4d63-8600-39f538dcac87",
		MailAddress: "resto5@test.com",
		Password:    "$2a$10$UkrQfUmAsPJ35cw5TVzJeOuoLySOWpMHN/b2zN561eixU0abBSCpq", // test
		Tell:        "08012341234",
		UserType:    Restaurant,
	},
	{ // ラーメン大王 西中島店
		UserUuid:    "5923b6b8-a4d6-4419-acf1-b1410480b0b5",
		MailAddress: "resto6@test.com",
		Password:    "$2a$10$UkrQfUmAsPJ35cw5TVzJeOuoLySOWpMHN/b2zN561eixU0abBSCpq", // test
		Tell:        "08012341234",
		UserType:    Restaurant,
	},
}
