package user

import (
	"University/model"
)

type ControllerImpl struct {
	dao Dao
}

func NewController(dao Dao) Controller {
	return &ControllerImpl{dao: dao}
}

func (ctrl *ControllerImpl) AddUser(user model.User) (err error) {
	return
}

func (ctrl *ControllerImpl) DeleteUser(id int) (err error) {
	return
}

func (ctrl *ControllerImpl) GetUser(id int) (user model.User, err error) {
	return
}
