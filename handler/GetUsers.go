package handler

import (
	"net/http"

	"github.com/NFortun/Astrobot/db"
	"github.com/NFortun/Astrobot/models"
	"github.com/go-openapi/runtime/middleware"

	api "github.com/NFortun/Astrobot/restapi/operations"
)

func GetUsers(params api.GetUsersParams) middleware.Responder {
	users, err := db.GetUsers()
	if err != nil {
		errMsg := err.Error()
		return api.NewGetImageOfTheDayDefault(http.StatusInternalServerError).WithPayload(&models.Error{
			Message: &errMsg,
		})
	}

	return api.NewGetUsersOK().WithPayload(users)
}
