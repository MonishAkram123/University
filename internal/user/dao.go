package user

import "University/model"

type Dao interface {
	Add(user model.User) (err error)
	DeleteById(id int) (err error)
	GetByReg(regNo string) (user model.User, err error)
}
