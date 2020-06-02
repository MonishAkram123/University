package main

import (
	"University/internal/user"
	"University/pkg/server"
	"github.com/sirupsen/logrus"
	"net/http"
)

func listenAndServe() {
	s := server.NewHttpServer()
	logrus.Fatal(s.ListenAndServe(getEndpoints()))
}

func getEndpoints() []server.Endpoint {
	handler := user.NewHandler(user.NewController(user.GetUsersDao()))
	return []server.Endpoint {
		{Method:http.MethodGet, Path: "/users/{id}", Handler: handler.Get},
		{Method:http.MethodPost, Path: "/users", Handler: handler.Add},
		{Method:http.MethodDelete, Path: "/users/{id}", Handler: handler.Delete},
	}
}