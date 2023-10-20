package astrobin

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"

	"github.com/NFortun/Astrobot/models"

	"github.com/NFortun/Astrobot/internal/astrobin"
	api "github.com/NFortun/Astrobot/restapi/operations"
)

type Url struct {
	Url string `json:"url_hd"`
}

func GetImageOfTheDay(params api.GetImageOfTheDayParams) middleware.Responder {
	imageInformations, err := getImageOfTheDay(astrobin.NewClient(http.DefaultClient))
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

// getImageOfTheDay
func getImageOfTheDay(client astrobin.AstrobinClient) (*astrobin.ImageInformations, error) {
	logrus.Info("getting image of the day")
	image, err := client.GetImageOfTheDay()
	if err != nil {
		logrus.Warnf("failed to retrieve IOTD informations: %s", err.Error())
		return nil, err
	}

	logrus.Infof("image url %s", image.Path)
	imageInformations, err := client.GetImageInformations(image.Path)
	if err != nil {
		logrus.Warnf("failed to retrieve image informations: %s", err.Error())
		return nil, err
	}

	return &imageInformations, nil
}
