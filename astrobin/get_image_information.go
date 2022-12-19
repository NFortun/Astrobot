package astrobin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/NFortun/Astrobot/config"
)

func (a *astrobin) GetImageInformations(path string) (ImageInformations, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?api_key=%s&api_secret=%s&format=json", config.Data.BasePath, path, config.Data.ApiKey, config.Data.ApiSecret), nil)
	if err != nil {
		return ImageInformations{}, err
	}

	response, err := a.Client.Do(req)
	if err != nil {
		return ImageInformations{}, err
	}

	if response.StatusCode != 200 {
		return ImageInformations{}, fmt.Errorf("error while requesting astrobin: %s", response.Status)
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return ImageInformations{}, err
	}

	imageInformations := ImageInformations{}
	err = json.Unmarshal(content, &imageInformations)
	if err != nil {
		return ImageInformations{}, err
	}

	if urlSplit := strings.Split(imageInformations.Url, "/0/"); len(urlSplit) > 0 {
		imageInformations.Url = urlSplit[0]
	}
	return imageInformations, nil
}
