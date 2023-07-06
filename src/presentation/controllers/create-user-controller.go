package controllers

import (
	"net/http"

	"github.com/thunderjr/go-clean-api/src/application/features"
)

type CreateUserController struct {
	createUserFeature *features.CreateUserFeature
}

func NewCreateUserController(f *features.CreateUserFeature) Controller {
	return &CreateUserController{
		createUserFeature: f,
	}
}

func (c *CreateUserController) Handle(data map[string]interface{}) Response {
	user, err := c.createUserFeature.CreateUser(data["name"].(string))

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
		Status: http.StatusCreated,
	}
}
