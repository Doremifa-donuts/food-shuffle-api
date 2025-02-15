package dto

type LoginUser struct {
	JtiToken string
}

type GetCourses struct {
	CourseUuid     string
	RestaurantUuid string
	CourseName     string
	Description    string
	Images         []string
	Price          int
	Minimum        int
}

type RestaurantPlace struct {
	Latitude  float64
	Longitude float64
}
