package user

import (
	"University/internal/user/mocks"
	"University/model"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type HandlerTestSuite struct {
	suite.Suite
	mocker *gomock.Controller
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) SetupTest() {
	suite.mocker = gomock.NewController(suite.T())
}

func (suite *HandlerTestSuite) TearDownTest() {
	suite.mocker.Finish()
}

func (suite *HandlerTestSuite) TestHandlerImplAddValidationError() {
	tests := []struct {
		body string
		want int
	}{
		{``, http.StatusBadRequest},
		{`{`, http.StatusBadRequest},
		{`}`, http.StatusBadRequest},
		{`{"name": "name1", "phone": "123456789"}`, http.StatusBadRequest},
		{`{"reg_no": "", "name": "name1", "phone": "123456789"}`, http.StatusBadRequest},
		{`{"reg_no": 1, "name": "name1", "phone": "123456789"}`, http.StatusBadRequest},
		{`{"reg_no": "CA1", "phone": "123456789"}`, http.StatusBadRequest},
		{`{"reg_no": "CA1", "name": "", "phone": "123456789"}`, http.StatusBadRequest},
		{`{"reg_no": "CA1", "name": 1, "phone": "123456789"}`, http.StatusBadRequest},
		{`{"reg_no": "CA1", "name": "name1"}`, http.StatusBadRequest},
		{`{"reg_no": "CA1", "name": "name1", "phone": ""}`, http.StatusBadRequest},
		{`{"reg_no": "CA1", "name": "name1", "phone": 123456789}`, http.StatusBadRequest},
	}

	for i, test := range tests {
		testCase := "t" + strconv.Itoa(i)
		suite.T().Run(testCase, func(tt *testing.T) {
			req, _ := http.NewRequest(http.MethodPost, "", strings.NewReader(test.body))
			handler := NewHandler(nil)
			got := handler.Add(req).StatusCode
			if got != test.want {
				tt.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}

func (suite *HandlerTestSuite) TestHandlerImplAddControllerError() {
	body := `{"reg_no": "CA1", "name": "name1", "phone": "123456789"}`
	want := http.StatusServiceUnavailable

	mockCtrl := mocks.NewMockController(suite.mocker)
	mockCtrl.EXPECT().
		AddUser(model.User{RegNo: "CA1", Name: "name1", Phone: "123456789"}).Times(1).Return(errors.New(""))
	handler := NewHandler(mockCtrl)
	req, _ := http.NewRequest(http.MethodPost, "", strings.NewReader(body))

	got := handler.Add(req).StatusCode
	if got != want {
		suite.T().Errorf("got: %v, want: %v", got, want)
	}
}

func (suite *HandlerTestSuite) TestHandlerImplAddControllerSuccess() {
	body := `{"reg_no": "CA1", "name": "name1", "phone": "123456789"}`
	want := http.StatusCreated

	mockCtrl := mocks.NewMockController(suite.mocker)
	mockCtrl.EXPECT().AddUser(model.User{RegNo: "CA1", Name: "name1", Phone: "123456789"}).Times(1).Return(nil)
	handler := NewHandler(mockCtrl)
	req, _ := http.NewRequest(http.MethodPost, "", strings.NewReader(body))

	got := handler.Add(req).StatusCode
	if got != want {
		suite.T().Errorf("got: %v, want: %v", got, want)
	}
}

func (suite *HandlerTestSuite) TestHandlerImplDeleteIdError() {
	tests := []struct {
		id   string
		want int
	}{
		{"", http.StatusBadRequest},
		{"a", http.StatusBadRequest},
		{"1a", http.StatusBadRequest},
		{"null", http.StatusBadRequest},
	}

	for i, test := range tests {
		testCase := "t" + strconv.Itoa(i)
		suite.T().Run(testCase, func(tt *testing.T) {
			req, _ := http.NewRequest(http.MethodDelete, "user/"+test.id, strings.NewReader(""))

			req = mux.SetURLVars(req, map[string]string{"id": test.id})
			handler := NewHandler(nil)
			got := handler.Delete(req).StatusCode
			if got != test.want {
				tt.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}

func (suite *HandlerTestSuite) TestHandlerImplDeleteControllerError() {
	id := 1
	want := http.StatusServiceUnavailable
	mockCtrl := mocks.NewMockController(suite.mocker)
	mockCtrl.EXPECT().DeleteUser(id).Times(1).Return(errors.New(""))
	handler := NewHandler(mockCtrl)
	req, _ := http.NewRequest(http.MethodDelete, "user/"+strconv.Itoa(id), strings.NewReader(""))

	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(id)})
	got := handler.Delete(req).StatusCode
	if got != want {
		suite.T().Errorf("got: %v, want: %v", got, want)
	}
}

func (suite *HandlerTestSuite) TestHandlerImplDeleteControllerSuccess() {
	id := 1
	want := http.StatusNoContent
	mockCtrl := mocks.NewMockController(suite.mocker)
	mockCtrl.EXPECT().DeleteUser(id).Times(1).Return(nil)
	handler := NewHandler(mockCtrl)

	req, _ := http.NewRequest(http.MethodDelete, "user/"+strconv.Itoa(id), strings.NewReader(""))

	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(id)})
	got := handler.Delete(req).StatusCode
	if got != want {
		suite.T().Errorf("got: %v, want: %v", got, want)
	}
}

func (suite *HandlerTestSuite) TestHandlerImplGetIdError() {
	tests := []struct {
		id   string
		want int
	}{
		{"", http.StatusBadRequest},
		{"a", http.StatusBadRequest},
		{"1a", http.StatusBadRequest},
		{"null", http.StatusBadRequest},
	}

	for i, test := range tests {
		testCase := "t" + strconv.Itoa(i)
		suite.T().Run(testCase, func(tt *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "user/"+test.id, strings.NewReader(""))

			req = mux.SetURLVars(req, map[string]string{"id": test.id})
			handler := NewHandler(nil)
			got := handler.Get(req).StatusCode
			if got != test.want {
				tt.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}

func (suite *HandlerTestSuite) TestHandlerImplGetControllerError() {
	reg := "CA1"
	wantCode := http.StatusServiceUnavailable

	mockCtrl := mocks.NewMockController(suite.mocker)
	mockCtrl.EXPECT().GetUser(reg).Times(1).Return(model.User{}, errors.New(""))
	handler := NewHandler(mockCtrl)
	req, _ := http.NewRequest(http.MethodGet, "user/"+reg, strings.NewReader(""))

	req = mux.SetURLVars(req, map[string]string{"reg": reg})
	got := handler.Get(req)
	if got.StatusCode != wantCode {
		suite.T().Errorf("got: %v, want: %v", got, wantCode)
	}
}

func (suite *HandlerTestSuite) TestHandlerImplGetControllerSuccess() {
	reg := "CA1"
	wantCode := http.StatusOK
	wantUser := model.User{Id: 1, RegNo: reg, Name: "name1", Phone: "123456789"}

	mockCtrl := mocks.NewMockController(suite.mocker)
	mockCtrl.EXPECT().GetUser(reg).Times(1).Return(wantUser, nil)
	handler := NewHandler(mockCtrl)
	req, _ := http.NewRequest(http.MethodGet, "user/"+reg, strings.NewReader(""))

	req = mux.SetURLVars(req, map[string]string{"reg": reg})
	got := handler.Get(req)
	if got.StatusCode != wantCode {
		suite.T().Errorf("got: %v, want: %v", got, wantCode)
		return
	}

	var gotBytes []byte
	gotBytes, _ = ioutil.ReadAll(got.Body)
	wantBytes, _ := json.Marshal(wantUser)

	if !reflect.DeepEqual(gotBytes, wantBytes) {
		suite.T().Errorf("got: %v, want: %v", string(gotBytes), string(wantBytes))
	}
}
