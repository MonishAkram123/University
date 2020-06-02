package user

import (
	"University/pkg/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func getClient() (client *sql.DB, err error) {
	connStr := getConnStr()
	client, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, errors.Wrap(err, "sql.Open.error")
	}

	err = client.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "client.Ping.error")
	}

	return
}

func getConnStr() string {
	host := config.V.GetString("USERS_DB_HOST")
	port := config.V.GetString("USERS_DB_PORT")
	username := config.V.GetString("USERS_DB_USER")
	password := config.V.GetString("USERS_DB_PWD")
	dbName := config.V.GetString("USERS_DB_NAME")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbName)
}
