package mock

import (
	astrobin "github.com/NFortun/Astrobot/astrobin"

	"github.com/stretchr/testify/mock"
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
