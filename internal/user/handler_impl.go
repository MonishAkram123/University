package user

import (
	"University/model"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type HandlerImpl struct {
	controller Controller
}

func NewHandler(controller Controller) Handler {
	return &HandlerImpl{controller: controller}
}

func (handler *HandlerImpl) Add(request *http.Request) (response http.Response) {
	logrus.Debug("user.handler.Add.called")

	user, err := readUserRequestBody(request.Body)
	if err != nil {
		logrus.WithError(err).Error("user.handler.Add.error")
		return http.Response{StatusCode: http.StatusBadRequest}
	}

	log := logrus.WithField("user", user)
	err = handler.controller.AddUser(user)
	if err != nil {
		log.Error("user.handler.Add.error")
		return http.Response{StatusCode: http.StatusServiceUnavailable}
	}

	log.Info("user.handler.Add.success")

	return http.Response{StatusCode: http.StatusCreated}
}

func (handler *HandlerImpl) Delete(request *http.Request) (response http.Response) {
	logrus.Debug("user.handler.Delete.called")

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		logrus.WithError(err).Error("user.handler.Delete.error")
		return http.Response{StatusCode: http.StatusBadRequest}
	}

	log := logrus.WithField("id", id)

	err = handler.controller.DeleteUser(id)
	if err != nil {
		log.Error("user.handler.Add.error")
		return http.Response{StatusCode: http.StatusServiceUnavailable}
	}

	log.Debug("user.handler.Delete.success")

	return http.Response{StatusCode: http.StatusNoContent}
}

func (handler *HandlerImpl) Get(request *http.Request) (response http.Response) {
	logrus.Debug("user.handler.Get.called")

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		logrus.WithError(err).Error("user.handler.Get.error")
		return http.Response{StatusCode: http.StatusBadRequest}
	}

	log := logrus.WithField("id", id)

	user, err := handler.controller.GetUser(id)
	if err != nil {
		log.Error("user.handler.Get.error")
		return http.Response{StatusCode: http.StatusServiceUnavailable}
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		log.WithField("user", user).WithError(err).Error("user.handler.Get.error")
		return http.Response{StatusCode: http.StatusServiceUnavailable}
	}

	log.WithField("user", user).Debug("user.handler.Get.success")

	return http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBuffer(userBytes)),
	}
}

func readUserRequestBody(body io.ReadCloser) (user model.User, err error) {
	if err = json.NewDecoder(body).Decode(&user); err != nil {
		err = errors.Wrap(err, "readUserRequestBody.Decode.error")
		return
	}

	if err = user.Validate(); err != nil {
		err = errors.Wrap(err, "readUserRequestBody.Validate.error")
		return
	}

	return
}
