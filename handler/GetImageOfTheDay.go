package handler

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/NFortun/Astrobot/astrobin"
	"github.com/NFortun/Astrobot/helper"
	"github.com/NFortun/Astrobot/models"
	api "github.com/NFortun/Astrobot/restapi/operations"
)

type Url struct {
	Url string `json:"url_hd"`
}

func GetImageOfTheDay(params api.GetImageOfTheDayParams) middleware.Responder {
	imageInformations, err := helper.GetImageOfTheDay(astrobin.NewClient(http.DefaultClient))
	if err != nil {
		errMessage := err.Error()
		return api.NewGetImagesDefault(http.StatusInternalServerError).WithPayload(&models.Error{
			Message: &errMessage,
		})
	}

	return api.NewGetImageOfTheDayOK().WithPayload(
		&models.ImageResponse{
			URL:         &imageInformations.Url,
			User:        &imageInformations.User,
			Description: &imageInformations.Description,
			Title:       &imageInformations.Title,
		},
	)
}
