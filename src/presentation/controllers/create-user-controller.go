package controllers

import (
	"net/http"

	"github.com/thunderjr/go-clean-api/src/application/features"
	"github.com/thunderjr/go-clean-api/src/domain/entities"
)

type CreateUserController struct {
	createUserFeature *features.CreateUserFeature
}

func NewCreateUserController(f *features.CreateUserFeature) Controller {
	return &CreateUserController{
		createUserFeature: f,
	}
}

func (c *CreateUserController) Handle(data any) Response {
	user, err := c.createUserFeature.CreateUser((data.(*entities.User)).Name)

	if err != nil {
		return Response{
			Data: map[string]interface{}{
				"message": err.Error(),
			},
			Status: http.StatusInternalServerError,
		}
	}

	return Response{
		Data:   user,
		Status: http.StatusOK,
	}
}
