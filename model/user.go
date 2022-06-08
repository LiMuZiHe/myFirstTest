package model

import "wxproject1/dao"

type User struct {
	ID       string `json:"id" form:"id"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

func (user *User) SelectSelf() (res *User, err error) {
	result := dao.DB.Table("users").Where("id = ? AND role = ?", user.ID, user.Role).First(&res)
	if result.Error != nil {
		err = result.Error
	}
	return
}
func (user *User) SelectCounselor() (counselor *dao.Counselor, err error) {
	result := dao.DB.Preload("Person").Table("counselors").Where("job_number = ?", user.ID).First(&counselor)
	if result.Error != nil {
		err = result.Error
	}
	return
}
