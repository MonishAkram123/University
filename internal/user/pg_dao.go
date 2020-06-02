package user

import (
	"University/model"
	"database/sql"
	"github.com/eapache/go-resiliency/retrier"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var (
	dao  Dao
	once sync.Once
)

const (
	retryCount = 3
	retryDelay = 50* time.Millisecond
)

type PGDao struct {
	jobRetrier *retrier.Retrier
	db         *sql.DB
}

func GetUsersDao() Dao {
	once.Do(func() {
		pgDao, err := newPgDao()
		if err != nil {
			logrus.WithError(err).Fatal("newPgDao.error")
		}
		dao = &pgDao
	})

	return dao
}

func newPgDao() (dao PGDao, err error) {
	client, err := getClient()
	if err != nil {
		return dao, errors.Wrap(err, "getClient.error")
	}

	logrus.Info("connection.establish.success")

	dao = PGDao{
		jobRetrier: retrier.New(retrier.ConstantBackoff(retryCount, retryDelay), nil),
		db:         client,
	}

	return
}

func (dao *PGDao) Add(user model.User) (err error) {
	query := `INSERT INTO users(reg_no, name, phone) VALUES($1, $2, $3);`
	return dao.jobRetrier.Run(func() error {
		_, err := dao.db.Exec(query, user.RegNo, user.Name, user.Phone)
		return err
	})
}

func (dao *PGDao) DeleteById(id int) (err error) {
	query := `DELETE FROM users WHERE id=$1`
	return dao.jobRetrier.Run(func() error {
		_, err := dao.db.Exec(query, id)
		return err
	})
}

func (dao *PGDao) GetByReg(regNo string) (user model.User, err error) {
	query := `SELECT id, name, reg_no, phone FROM users WHERE reg_no=$1`
	err = dao.jobRetrier.Run(func() error {
		row := dao.db.QueryRow(query, regNo)
		return row.Scan(&user.Id, &user.Name, &user.RegNo, &user.Phone)
	})

	logrus.Info(user)

	return
}
