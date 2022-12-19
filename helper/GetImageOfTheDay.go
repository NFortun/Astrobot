package helper

import (
	astrobin "github.com/NFortun/Astrobot/astrobin"

	"github.com/sirupsen/logrus"
)

// GetImageOfTheDay
func GetImageOfTheDay(client astrobin.AstrobinClient) (*astrobin.ImageInformations, error) {
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
