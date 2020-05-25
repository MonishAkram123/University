package user

import "net/http"

type Handler interface {
	Add(request *http.Request) (response http.Response)
	Delete(request *http.Request) (response http.Response)
	Get(request *http.Request) (response http.Response)
}
