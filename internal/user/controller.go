package user

import (
	"University/model"
)

type Controller interface {
	AddUser(user model.User) (err error)
	DeleteUser(id int) (err error)
	GetUser(regNo string) (user model.User, err error)
}
