package controllers

type Response struct {
	Data   interface{}
	Status int
}

type Controller interface {
	Handle(data any) Response
}
