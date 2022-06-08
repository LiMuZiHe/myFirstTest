package model

import (
	"wxproject1/dao"
)

func SelectPersonMin(idNumber string) (personMin dao.Person) {
	personMin = dao.SelectPersonMin(idNumber)
	return
}
