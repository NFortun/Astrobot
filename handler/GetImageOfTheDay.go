package handler

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"

	"astrobot/astrobin"
	"astrobot/models"
	api "astrobot/restapi/operations"
)

type Url struct {
	Url string `json:"url_hd"`
}

func GetImageOfTheDay(params api.GetImageOfTheDayParams) middleware.Responder {
	logrus.Info("getting image of the day")
	image, err := astrobin.GetImageOfTheDay()
	if err != nil {
		errMessage := err.Error()
		logrus.Warn("failed to retrive IOTD informations: %s", errMessage)
		return api.NewGetImageOfTheDayDefault(http.StatusInternalServerError).WithPayload(&models.Error{Message: &errMessage})
	}

	logrus.Infof("image url %s", image.Path)
	imageInformations, err := astrobin.GetImageInformation(image.Path)
	if err != nil {
		errMessage := err.Error()
		logrus.Warn("failed to retrieve image informations")
		return api.NewGetImageOfTheDayDefault(http.StatusInternalServerError).WithPayload(&models.Error{Message: &errMessage})
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
