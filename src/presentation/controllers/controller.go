package controllers

type Response struct {
	Data   map[string]interface{}
	Status int
}

type Controller interface {
	Handle(data any) (res Response, err error)
}
