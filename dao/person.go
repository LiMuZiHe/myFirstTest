package dao

func SelectPersonMin(idNumber string) (person Person) {
	DB.Preload("Reports").Where("id_number = ?", idNumber).Take(&person)
	//DB.Table("reports").Where("id_number = ?", idNumber).Find(&personMin.Reports)
	return
}
