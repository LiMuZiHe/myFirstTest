package dao

func GetStudentInfo(studentId string) (studentInfo Student) {
	DB.Preload("Person").Where("student_id = ?", studentId).Take(&studentInfo)
	return
}
func UpdataStudent(studentInfo *Student) {
	DB.Model(&Student{}).Omit("student_id").Where("student_id = ?", studentInfo.StudentID).Updates(studentInfo)
}
func CreateJourney(journey *Journey) {
	DB.Create(journey)
}
func CreateReport(report *Report) {
	DB.Create(report)
}
