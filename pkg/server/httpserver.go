package server

import (
	"University/pkg/config"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	hostAddrKey = "HOST"
	portKey     = "PORT"
)

type HttpServer struct {
	router *mux.Router
}

func NewHttpServer() Server {
	return &HttpServer{router: mux.NewRouter()}
}

func (server *HttpServer) ListenAndServe(endpoints []Endpoint) error {
	for _, endpoint := range endpoints {
		server.router.HandleFunc(endpoint.Path, wrapperHandler(endpoint.Handler)).Methods(endpoint.Method)
	}

	addr := config.V.GetString(hostAddrKey) + ":" + strconv.Itoa(config.V.GetInt(portKey))
	logrus.WithField("addr", addr).Info("server started")
	return http.ListenAndServe(addr, server.router)
}

func wrapperHandler(handler func(req *http.Request) http.Response) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := handler(req)
		w.WriteHeader(resp.StatusCode)

		if resp.Body == nil {
			return
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.WithError(err).Fatal("resp.Body.Read.error")
		}

		_, err = w.Write(bytes)
		if err != nil {
			logrus.WithError(err).Fatal("w.Write.bytes.error")
		}

		err = resp.Body.Close()
		if err != nil {
			logrus.WithError(err).Fatal("resp.Body.Close.error")
		}
	}
}
