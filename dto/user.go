package dto

type LoginUser struct {
	JtiToken string
}

type GetCourses struct {
	CourseUuid     string
	RestaurantUuid string
	CourseName     string
	Discription    string
	Images         []string
	Price          int
	Minimum        int
}
