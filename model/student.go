package model

import "wxproject1/dao"

type Student dao.Student
type Journey dao.Journey
type Report dao.Report

func GetStudentInfo(studentId string) (studentInfo dao.Student) {
	studentInfo = dao.GetStudentInfo(studentId)
	return
}

func (student *Student) UpdataInfo() {
	dao.UpdataStudent((*dao.Student)(student))
}
func (journer *Journey) Create() {
	dao.CreateJourney((*dao.Journey)(journer))
}

func (report *Report) Create() {
	dao.CreateReport((*dao.Report)(report))
}
