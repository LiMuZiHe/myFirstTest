package dao

import "time"

type Student struct {
	Person            Person `json:"person" form:"person" gorm:"person;foreignKey:IdNumber"`
	IdNumber          string `json:"id_number" form:"id_number" gorm:"id_number"`
	StudentID         string `json:"student_id" form:"student_id" gorm:"student_id;primaryKey"`
	EnterYear         int    `json:"enter_year" form:"enter_year" gorm:"enter_year"`
	Dormitory         string `json:"dormitory" form:"dormitory" gorm:"dormitory"`
	DormitoryNum      string `json:"dormitory_num" form:"dormitory_num" gorm:"dormitory_num"`
	Class             string `json:"class" form:"class" gorm:"class"`
	WhetherBackSchool *bool  `json:"whether_back_school" form:"whether_back_school" gorm:"whether_back_school"`
	Collega           string `json:"collega" form:"collega" gorm:"collega"`
}

type Person struct {
	IdNumber    string   `gorm:"id_number;primaryKey"`
	Name        string   `json:"name" gorm:"name"`
	Avatar      string   `json:"avatar" gorm:"photo"`
	LastAddress string   `json:"last_address" gorm:"last_address"`
	Hometown    string   `gorm:"hometown"`
	Gender      string   `gorm:"gender"`
	Reports     []Report `gorm:"reports;foreignKey:IdNumber"`
}

type Journey struct {
	Id                uint   `gorm:"id" form:"id"`
	BeforeMoving      string `gorm:"before_moving" form:"before_moving"`
	AfterMoving       string `gorm:"after_moving" form:"after_moving"`
	StudentID         string `gorm:"student_id" form:"student_id"`
	JourneyCard       string `gorm:"journey_card" form:"journey_card"`
	DetailedItinerary string `gorm:"detailed_itinerary" form:"detailed_itinerary"`
}

type Report struct {
	ID                   int       `form:"id" gorm:"column:id;primaryKey"`
	StudentID            string    `form:"student_id" gorm:"column:student_id"`
	IdNumber             string    `form:"id_number" gorm:"column:id_number"`
	ReportScreenshot     string    `form:"report_screenshot" gorm:"column:screenshot"`
	ReportResult         bool      `form:"report_result" gorm:"column:result"`
	SampleCollectionTime time.Time `form:"sample_collection_time" gorm:"column:time"`
}
type Counselor struct {
	Person    Person `json:"person" form:"person" gorm:"person;foreignKey:IdNumber"`
	IdNumber  string `json:"id_number" form:"id_number"`
	College   string `json:"college" form:"college"`
	Grade     string `json:"grade" form:"grade"`
	JobNumber string `json:"job_number" form:"job_number"`
}
