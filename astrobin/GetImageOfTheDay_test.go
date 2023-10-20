package astrobin

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/NFortun/Astrobot/internal/astrobin"
)

type MockAstrobinClient struct {
	mock.Mock
}

func (m *MockAstrobinClient) GetImageOfTheDay() (*astrobin.ImageOfTheDay, error) {
	args := m.Called()
	return args.Get(0).(*astrobin.ImageOfTheDay), args.Error(1)
}

func (m *MockAstrobinClient) GetImageInformations(path string) (astrobin.ImageInformations, error) {
	args := m.Called(path)
	return args.Get(0).(astrobin.ImageInformations), args.Error(1)
}

func (m *MockAstrobinClient) GetImages(opts []*astrobin.QueryOpts) (*astrobin.ImagesInformations, error) {
	args := m.Called(opts)
	return args.Get(0).(*astrobin.ImagesInformations), args.Error(1)
}

func TestGetImageOfTheDay(t *testing.T) {

	ImageOfTheDay := &astrobin.ImageOfTheDay{
		Date: "Date",
		Path: "Path",
	}

	mockAstrobinClient := new(MockAstrobinClient)
	t.Run("GetImageOfTheDay_error", func(t *testing.T) {
		GetImageOfTheDayError := errors.New("GetImageOfTheDay error")
		mockCall := mockAstrobinClient.On("GetImageOfTheDay").Return(ImageOfTheDay, GetImageOfTheDayError)
		defer mockCall.Unset()
		response, err := getImageOfTheDay(mockAstrobinClient)
		assert.Equal(t, GetImageOfTheDayError, err)
		assert.Nil(t, response)
	})

	t.Run("GetImageInformations_error", func(t *testing.T) {
		imageInformations := astrobin.ImageInformations{
			Url:         "URL",
			Title:       "TITLE",
			Description: "DESC",
			User:        "USER",
		}
		GetImageInformationsError := errors.New("GetImageInformations error")
		GetImageOfTheDayMock := mockAstrobinClient.On("GetImageOfTheDay").Return(ImageOfTheDay, nil)
		defer GetImageOfTheDayMock.Unset()
		GetImageInformationsMock := mockAstrobinClient.On("GetImageInformations", ImageOfTheDay.Path).Return(imageInformations, GetImageInformationsError)
		defer GetImageInformationsMock.Unset()
		response, err := getImageOfTheDay(mockAstrobinClient)
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
		response, err := getImageOfTheDay(mockAstrobinClient)
		assert.Nil(t, err)
		assert.Equal(t, imageInformations, *response)
	})

}
