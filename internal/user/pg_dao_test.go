package user

import (
	"University/model"
	"University/pkg/config"
	"github.com/eapache/go-resiliency/retrier"
	"github.com/stretchr/testify/suite"
	"reflect"
	"strconv"
	"testing"
)

type UsersDaoTestSuite struct {
	suite.Suite
	testDao PGDao
}

func TestPGDaoSuite(t *testing.T) {
	suite.Run(t, new(UsersDaoTestSuite))
}

func (suite *UsersDaoTestSuite) SetupSuite() {
	err := config.Init()
	if err != nil {
		panic(err)
	}

	client, err := getClient()
	if err != nil {
		panic(err)
	}

	suite.testDao = PGDao{db: client, jobRetrier: retrier.New(nil, nil)}
}

func (suite *UsersDaoTestSuite) TearDownSuite() {
	query := `TRUNCATE users;`
	_, err := suite.testDao.db.Exec(query)
	if err != nil {
		suite.T().Errorf("unable to tear down suite: %v", err)
	}
}

func (suite *UsersDaoTestSuite) SetupTest() {
	queries := []string{
		`TRUNCATE users;`,
		`INSERT INTO users(id, reg_no, name, phone) VALUES(1, 'CA1', 'John', '123456789')`,
	}

	for _, query := range queries {
		_, err := suite.testDao.db.Exec(query)
		if err != nil {
			suite.T().Errorf("setup test failed: %v", err)
		}
	}

}

func (suite *UsersDaoTestSuite) Test_PgDao_Add() {
	tests := []struct {
		input   model.User
		wantErr bool
	}{
		{input: model.User{RegNo: "CA1", Name: "John", Phone: "123456789"}, wantErr: true},
		{input: model.User{RegNo: "CA2", Name: "", Phone: ""}, wantErr: false},
		{input: model.User{RegNo: "CA2", Name: "Doe", Phone: "1234"}, wantErr: true},
		{input: model.User{RegNo: "CA3", Name: "", Phone: ""}, wantErr: true},
		{input: model.User{RegNo: "CA3", Name: "", Phone: "1234"}, wantErr: false},
	}

	for i, test := range tests {
		suite.T().Run("t"+strconv.Itoa(i), func(t *testing.T) {
			if err := suite.testDao.Add(test.input); (err != nil) != test.wantErr {
				t.Errorf("Add wantErr: %v, got: %v", test.wantErr, err != nil)
			}
		})
	}
}

func (suite *UsersDaoTestSuite) Test_PgDao_GetById() {
	tests := []struct {
		input    int
		wantUser model.User
		wantErr  bool
	}{
		{input: 1, wantUser: model.User{Id: 1, RegNo: "CA1", Name: "John", Phone: "123456789"}, wantErr: false},
		{input: 2, wantErr: true},
	}

	for i, test := range tests {
		suite.T().Run("t"+strconv.Itoa(i), func(t *testing.T) {
			if user, err := suite.testDao.GetById(test.input); (err != nil) != test.wantErr {
				t.Errorf("GetById wantErr: %v, got: %v", test.wantErr, err != nil)
			} else if !reflect.DeepEqual(user, test.wantUser) {
				t.Errorf("GetById wantUser: %v, got: %v", test.wantUser, user)
			}
		})
	}
}

func (suite *UsersDaoTestSuite) Test_PgDao_DeleteById() {
	tests := []struct {
		input   int
		wantErr bool
	}{
		{input: 1, wantErr: false},
		{input: 2, wantErr: false},
	}

	for i, test := range tests {
		suite.T().Run("t"+strconv.Itoa(i), func(t *testing.T) {
			if err := suite.testDao.DeleteById(test.input); (err != nil) != test.wantErr {
				t.Errorf("DeleteById wantErr: %v, got: %v", test.wantErr, err != nil)
			} else if _, err := suite.testDao.GetById(test.input); err == nil {
				t.Errorf("DeleteById wantErr but got nil")
			}
		})
	}
}
