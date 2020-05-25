package model

import (
	"University/utils"
	"errors"
)

type User struct {
	Id    int
	RegNo string
	Name  string
	Phone string
}

type Users []User

func (user *User) Validate() error {
	if utils.IsEmptyString(user.RegNo) {
		return errors.New("field RegNo is empty")
	}

	if utils.IsEmptyString(user.Name) {
		return errors.New("field Name is empty")
	}

	if utils.IsEmptyString(user.Name) {
		return errors.New("field Phone is empty")
	}

	return nil
}
