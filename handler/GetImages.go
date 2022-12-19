package handler

import (
	"net/http"

	"github.com/NFortun/Astrobot/astrobin"
	"github.com/NFortun/Astrobot/models"
	api "github.com/NFortun/Astrobot/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

func GetImages(params api.GetImagesParams) middleware.Responder {
	var opts []*astrobin.QueryOpts

	logrus.Info("retrieving images")
	if params.EndDate != nil {
		opts = append(opts, &astrobin.QueryOpts{
			Name:     "uploaded",
			Operator: astrobin.LESS_EQUAL,
			Value:    params.EndDate.String(),
		})
	}

	if params.Limit != nil {
		opts = append(opts, &astrobin.QueryOpts{
			Name:     "limit",
			Operator: astrobin.EQUAL,
			Value:    *params.Limit,
		})
	}

	if params.StartDate != nil {
		opts = append(opts, &astrobin.QueryOpts{
			Name:     "uploaded",
			Operator: astrobin.GREATER_EQUAL,
			Value:    params.StartDate.String(),
		})
	}

	if params.Offset != nil {
		opts = append(opts, &astrobin.QueryOpts{
			Name:     "offset",
			Operator: astrobin.EQUAL,
			Value:    *params.Offset,
		})
	}

	if params.User != nil {
		opts = append(opts, &astrobin.QueryOpts{
			Name:     "user",
			Operator: astrobin.EQUAL,
			Value:    *params.User,
		})
	}

	logrus.Infof("getting images with %d parameters", len(opts))
	client := astrobin.NewClient(http.DefaultClient)
	images, err := client.GetImages(opts)
	if err != nil {
		errMessage := err.Error()
		return api.NewGetImageOfTheDayDefault(http.StatusInternalServerError).WithPayload(&models.Error{Message: &errMessage})
	}

	imagesResponse := models.ImagesResponse{}
	for i := range images.Images {
		imagesResponse = append(imagesResponse, &models.ImageResponse{
			Description: &images.Images[i].Description,
			Title:       &images.Images[i].Title,
			URL:         &images.Images[i].Url,
			User:        &images.Images[i].User,
		})

	}

	return api.NewGetImagesOK().WithPayload(imagesResponse)

}
