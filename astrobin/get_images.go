package astrobin

import (
	"astrobot/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func GetImages(opts []*QueryOpts) (*ImagesInformations, error) {
	var stringOpts []string
	for _, opt := range opts {
		stringOpts = append(stringOpts, opt.String())
	}

	queryParams := ""
	if len(stringOpts) > 0 {
		queryParams = fmt.Sprintf("&%s", strings.Join(stringOpts, "&"))
	}

	completePath := fmt.Sprintf("%s/api/v1/image/?api_key=%s&api_secret=%s&format=json%s",
		config.Data.BasePath,
		config.Data.ApiKey,
		config.Data.ApiSecret,
		queryParams,
	)
	req, err := http.NewRequest(
		http.MethodGet,
		completePath,
		nil,
	)
	if err != nil {
		return nil, err
	}

	logrus.Info(completePath)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error while retrieving images: %s %s", string(data), response.Status)
	}

	ImagesResponse := &ImagesInformations{}
	err = json.Unmarshal(data, ImagesResponse)
	if err != nil {
		return nil, err
	}

	return ImagesResponse, nil

}
