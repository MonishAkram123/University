package server

type Server interface {
	ListenAndServe(endpoints []Endpoint) error
}



