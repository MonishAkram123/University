package main

import (
	"University/pkg/config"
	"github.com/sirupsen/logrus"
)

func main() {
	err := config.Init()
	if err != nil {
		logrus.WithError(err).Fatal("user.main.Init.error")
	}
	logrus.Info("config.Init.success")
	listenAndServe()
}
