package astrobin

import (
	"net/http"
)

type astrobin struct {
	Client *http.Client
}

type AstrobinClient interface {
	GetImageInformations(path string) (ImageInformations, error)
	GetImages(opts []*QueryOpts) (*ImagesInformations, error)
	GetImageOfTheDay() (*ImageOfTheDay, error)
}

type ImageOfTheDay struct {
	Date string `json:"date"`
	Path string `json:"image"`
}

type ImagesOfTheDay struct {
	Images []ImageOfTheDay `json:"objects"`
}

type ImageInformations struct {
	Url         string `json:"url_hd"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User        string `json:"user"`
}

type ImagesInformations struct {
	Images []ImageInformations `json:"objects"`
}
