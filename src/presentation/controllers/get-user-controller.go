package controllers

import (
	"net/http"

	"github.com/thunderjr/go-clean-api/src/application/features"
)

type GetUserController struct {
	GetUserFeature *features.GetUserFeature
}

func NewGetUserController(f *features.GetUserFeature) Controller {
	return &GetUserController{
		GetUserFeature: f,
	}
}

func (c *GetUserController) Handle(data map[string]interface{}) *Response {
	var userId string

	if _, ok := data["id"].(string); !ok {
		return &Response{
			Data: map[string]interface{}{
				"message": "Error parsing the query data",
			},
			Status: http.StatusBadRequest,
		}
	}

	userId = data["id"].(string)
	if userId == "" {
		return &Response{
			Data: map[string]interface{}{
				"message": "You must past a user ID",
			},
			Status: http.StatusBadRequest,
		}
	}

	user, err := c.GetUserFeature.GetUser(userId)
	if err != nil {
		return &Response{
			Data: map[string]interface{}{
				"message": err.Error(),
			},
			Status: http.StatusInternalServerError,
		}
	}

	return &Response{
		Data:   user,
		Status: http.StatusOK,
	}
}
