package testutils

import (
	"time"

	"github.com/NFortun/Astrobot/astrobin"
)

var (
	URL           = "url"
	Title         = "title"
	Description   = "description"
	User          = "user"
	ImageOfTheDay = &astrobin.ImageOfTheDay{
		Date: time.Now().String(),
		Path: URL,
	}

	ImageInformations = astrobin.ImageInformations{
		Url:         URL,
		Title:       Title,
		Description: Description,
		User:        User,
	}

	ImagesInformations = &astrobin.ImagesInformations{
		Images: []astrobin.ImageInformations{ImageInformations},
	}
)
