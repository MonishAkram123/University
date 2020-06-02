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
	return ctrl.dao.Add(user)
}

func (ctrl *ControllerImpl) DeleteUser(id int) (err error) {
	return ctrl.dao.DeleteById(id)
}

func (ctrl *ControllerImpl) GetUser(regNo string) (user model.User, err error) {
	return ctrl.dao.GetByReg(regNo)
}
