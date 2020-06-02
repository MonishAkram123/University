package user

import "University/model"

type Dao interface {
	Add(user model.User) (err error)
	DeleteById(id int) (err error)
	GetById(id int) (user model.User, err error)
}
