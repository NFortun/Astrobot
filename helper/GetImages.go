package helper

import (
	astrobin "github.com/NFortun/Astrobot/astrobin"
	models "github.com/NFortun/Astrobot/models"
)

// GetImages
func GetImages(opts []*astrobin.QueryOpts, client astrobin.AstrobinClient) (models.ImagesResponse, error) {
	images, err := client.GetImages(opts)
	if err != nil {
		return nil, err
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

	return imagesResponse, nil
}
