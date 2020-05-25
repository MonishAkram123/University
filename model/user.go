package model

import (
	"University/utils"
	"errors"
)

type User struct {
	Id    int    `json:"id"`
	RegNo string `json:"reg_no"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Users []User

func (user *User) Validate() error {
	if utils.IsEmptyString(user.RegNo) {
		return errors.New("field RegNo is empty")
	}

	if utils.IsEmptyString(user.Name) {
		return errors.New("field Name is empty")
	}

	if utils.IsEmptyString(user.Phone) {
		return errors.New("field Phone is empty")
	}

	return nil
}
