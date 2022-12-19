package helper_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	astrobin "github.com/NFortun/Astrobot/astrobin"
	config "github.com/NFortun/Astrobot/config"
	helper "github.com/NFortun/Astrobot/helper"
	models "github.com/NFortun/Astrobot/models"
	testutils "github.com/NFortun/Astrobot/testutils"

	astromock "github.com/NFortun/Astrobot/testutils/mock"
)

func TestGetImages(t *testing.T) {

	config.Data = config.Config{
		ApiKey:    "KEY",
		ApiSecret: "SECRET",
		BasePath:  "path",
	}

	mockAstrobinClient := new(astromock.MockAstrobinClient)

	t.Run("get_images_error", func(t *testing.T) {
		expectedErr := fmt.Errorf("GetImages error")
		getImagesMock := mockAstrobinClient.On("GetImages", mock.Anything).Return(&astrobin.ImagesInformations{}, expectedErr)
		defer getImagesMock.Unset()
		images, err := helper.GetImages(nil, mockAstrobinClient)
		assert.Nil(t, images)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("success", func(t *testing.T) {
		getImagesMock := mockAstrobinClient.On("GetImages", mock.Anything).Return(testutils.ImagesInformations, nil)
		defer getImagesMock.Unset()

		images, err := helper.GetImages(nil, mockAstrobinClient)
		assert.Nil(t, err)
		assert.Equal(t, models.ImagesResponse{
			{
				Description: &testutils.Description,
				Title:       &testutils.Title,
				URL:         &testutils.URL,
				User:        &testutils.User,
			},
		}, images)
	})
}
