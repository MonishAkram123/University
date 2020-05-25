package user

import (
	"University/model"
)

type Controller interface {
	AddUser(user model.User) (err error)
	DeleteUser(id int) (err error)
	GetUser(id int) (user model.User, err error)
}
