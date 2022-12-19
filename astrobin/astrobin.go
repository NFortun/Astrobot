package astrobin

import "net/http"

type AstrobinClient interface {
	GetImageInformations(path string) (ImageInformations, error)
	GetImages(opts []*QueryOpts) (*ImagesInformations, error)
	GetImageOfTheDay() (*ImageOfTheDay, error)
}

type astrobin struct {
	Client *http.Client
}

func NewClient(client *http.Client) AstrobinClient {
	return &astrobin{
		Client: client,
	}
}
