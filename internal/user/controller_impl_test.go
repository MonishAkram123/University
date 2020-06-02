package user

import (
	"University/internal/user/mocks"
	"University/model"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

func TestNewControllerImpl(t *testing.T) {
	suite.Run(t, new(TestControllerImplSuite))
}

type TestControllerImplSuite struct {
	suite.Suite
	mocker *gomock.Controller
}

func (suite *TestControllerImplSuite) SetupTest() {
	suite.mocker = gomock.NewController(suite.T())
}

func (suite *TestControllerImplSuite) TearDownTest() {
	suite.mocker.Finish()
}

func (suite *TestControllerImplSuite) TestAddUserError() {
	mockDao := mocks.NewMockDao(suite.mocker)
	user := model.User{RegNo: "CA1", Name: "name1", Phone: "123456789"}
	mockDao.EXPECT().Add(user).Return(errors.New(""))

	if err := NewController(mockDao).AddUser(user); err == nil {
		suite.T().Error("expected err but got nil")
	}
}

func (suite *TestControllerImplSuite) TestAddUserSuccess() {
	mockDao := mocks.NewMockDao(suite.mocker)
	user := model.User{RegNo: "CA1", Name: "name1", Phone: "123456789"}
	mockDao.EXPECT().Add(user).Return(nil)

	if err := NewController(mockDao).AddUser(user); err != nil {
		suite.T().Error("expected no err but got: ", err)
	}
}

func (suite *TestControllerImplSuite) TestDeleteUserError() {
	mockDao := mocks.NewMockDao(suite.mocker)
	id := 1
	mockDao.EXPECT().DeleteById(id).Return(errors.New(""))

	if err := NewController(mockDao).DeleteUser(id); err == nil {
		suite.T().Error("expected err but got nil")
	}
}

func (suite *TestControllerImplSuite) TestDeleteUserSuccess() {
	mockDao := mocks.NewMockDao(suite.mocker)
	id := 1
	mockDao.EXPECT().DeleteById(id).Return(nil)

	if err := NewController(mockDao).DeleteUser(id); err != nil {
		suite.T().Error("expected no err but got: ", err)
	}
}

func (suite *TestControllerImplSuite) TestGetUserError() {
	reg := "CA1"
	mockDao := mocks.NewMockDao(suite.mocker)
	mockDao.EXPECT().GetById(reg).Return(model.User{}, errors.New(""))

	if _, err := NewController(mockDao).GetUser(reg); err == nil {
		suite.T().Error("expected err but got nil")
	}
}

func (suite *TestControllerImplSuite) TestGetUserSuccess() {
	reg := "CA1"
	mockDao := mocks.NewMockDao(suite.mocker)
	want := model.User{Id: 1, RegNo: reg, Name: "name1", Phone: "123456789"}
	mockDao.EXPECT().GetById(reg).Return(want, nil)

	if got, err := NewController(mockDao).GetUser(reg); err != nil {
		suite.T().Error("expected no error but got: ", err)
	} else if !reflect.DeepEqual(want, got) {
		suite.T().Errorf("got: %v, want: %v", got, want)
	}

}