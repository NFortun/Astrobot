package astrobin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NFortun/Astrobot/config"

	"github.com/sirupsen/logrus"
)

func (a *astrobin) GetImageOfTheDay() (*ImageOfTheDay, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/imageoftheday/?limit=1&api_key=%s&api_secret=%s&format=json", config.Data.BasePath, config.Data.ApiKey, config.Data.ApiSecret), nil)
	if err != nil {
		return nil, err
	}

	logrus.Info("sending request")
	response, err := a.Client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error while requesting astrobin: %s", response.Status)

	}

	defer response.Body.Close()

	responseContent, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	mapResult := ImagesOfTheDay{}
	err = json.Unmarshal(responseContent, &mapResult)
	if err != nil {
		return nil, err
	}

	if len(mapResult.Images) == 0 {
		return nil, fmt.Errorf("missing image of the day")
	}

	return &mapResult.Images[0], nil
}
