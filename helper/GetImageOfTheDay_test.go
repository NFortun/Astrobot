package helper_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	astrobin "github.com/NFortun/Astrobot/astrobin"
	config "github.com/NFortun/Astrobot/config"
	helper "github.com/NFortun/Astrobot/helper"
	testutils "github.com/NFortun/Astrobot/testutils"
	"github.com/NFortun/Astrobot/testutils/mock"
)

func TestGetImageOfTheDay(t *testing.T) {

	config.Data = config.Config{
		ApiKey:    "KEY",
		ApiSecret: "SECRET",
		BasePath:  "path",
	}

	mockAstrobinClient := new(mock.MockAstrobinClient)
	t.Run("GetImageOfTheDay_error", func(t *testing.T) {
		GetImageOfTheDayError := errors.New("GetImageOfTheDay error")
		mockCall := mockAstrobinClient.On("GetImageOfTheDay").Return(testutils.ImageOfTheDay, GetImageOfTheDayError)
		defer mockCall.Unset()
		response, err := helper.GetImageOfTheDay(mockAstrobinClient)
		assert.Equal(t, GetImageOfTheDayError, err)
		assert.Nil(t, response)
	})

	t.Run("GetImageInformations_error", func(t *testing.T) {
		imageOfTheDay := &astrobin.ImageOfTheDay{
			Date: time.Now().String(),
			Path: "path",
		}
		GetImageInformationsError := errors.New("GetImageInformations error")
		GetImageOfTheDayMock := mockAstrobinClient.On("GetImageOfTheDay").Return(imageOfTheDay, nil)
		defer GetImageOfTheDayMock.Unset()
		GetImageInformationsMock := mockAstrobinClient.On("GetImageInformations", imageOfTheDay.Path).Return(astrobin.ImageInformations{}, GetImageInformationsError)
		defer GetImageInformationsMock.Unset()
		response, err := helper.GetImageOfTheDay(mockAstrobinClient)
		assert.Equal(t, GetImageInformationsError, err)
		assert.Nil(t, response)
	})

	t.Run("success", func(t *testing.T) {

		imageOfTheDay := &astrobin.ImageOfTheDay{
			Date: time.Now().String(),
			Path: "path",
		}

		imageInformations := astrobin.ImageInformations{
			Url:         "url",
			Title:       "title",
			Description: "description",
			User:        "user",
		}

		GetImageOfTheDayMock := mockAstrobinClient.On("GetImageOfTheDay").Return(imageOfTheDay, nil)
		defer GetImageOfTheDayMock.Unset()
		GetImageInformationsMock := mockAstrobinClient.On("GetImageInformations", imageOfTheDay.Path).Return(imageInformations, nil)
		defer GetImageInformationsMock.Unset()
		response, err := helper.GetImageOfTheDay(mockAstrobinClient)
		assert.Nil(t, err)
		assert.Equal(t, imageInformations, *response)
	})

}
