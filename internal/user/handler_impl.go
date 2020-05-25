package user

import (
	"net/http"
)

type HandlerImpl struct {
	controller Controller
}

func NewHandler(controller Controller) Handler {
	return &HandlerImpl{controller: controller}
}

func (handler *HandlerImpl) Add(request *http.Request) (response http.Response) {
	return
}

func (handler *HandlerImpl) Delete(request *http.Request) (response http.Response) {
	return
}

func (handler *HandlerImpl) Get(request *http.Request) (response http.Response) {
	return
}

func (handler *HandlerImpl) GetAll(request *http.Request) (response http.Response) {
	return
}
