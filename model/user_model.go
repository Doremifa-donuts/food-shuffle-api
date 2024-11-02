package model

type UserType string

const (
	General    UserType = "General"
	Restaurant UserType = "Restaurant"
)

type User struct {
	UserUuid    string   `gorm:"type:char(36);primary_key;"`                  // ユーザーのUUID
	MailAddress string   `gorm:"type:varchar(255);not null"`                  // メールアドレス
	Password    string   `gorm:"type:varchar(255);not null"`                  // ハッシュ化されたパスワード
	Tell        string   `gorm:"type:varchar(20);not null"`                   // 電話番号
	JtiToken    string   `gorm:"type:varchar(255);not null"`                  // jtiトークン　JWTの解析に使う
	UserType    UserType `gorm:"type:enum('General', 'Restaurant');not null"` // ユーザーの種類　一般利用者、レストラン利用者
	User *GeneralUser `gorm:"foreignKey:UserUuid"`
	Restaurant *RestaurantUser `gorm:"foreignKey:RestaurantUuid;references:UserUuid"`
}

// サンプルデータ
var SampleUsers = []User{
	{
		UserUuid:    "a0adb027-0f54-4c1a-9ed3-86041c863344",
		MailAddress: "general@test.com",
		Password:    "$2a$10$UkrQfUmAsPJ35cw5TVzJeOuoLySOWpMHN/b2zN561eixU0abBSCpq", // test
		Tell:        "08012341234",
		JtiToken:    "general_jti_token",
		UserType:    General,
	},
	{
		UserUuid:    "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		MailAddress: "resto@test.com",
		Password:    "$2a$10$UkrQfUmAsPJ35cw5TVzJeOuoLySOWpMHN/b2zN561eixU0abBSCpq", // test
		Tell:        "08056785678",
		JtiToken:    "resto_jti_token",
		UserType:    Restaurant,
	},
}
