package uploader

import (
	"context"
	"errors"
	"io"
	"os"

	"github.com/Davincible/goinsta/v3"
)

type Instagram struct {
	insta *goinsta.Instagram
}

func (i *Instagram) Connect(ctx context.Context) error {
	insta := goinsta.New(os.Getenv("INSTA_USER"), os.Getenv("INSTA_PASSWORD"))
	if err := insta.Login(); err != nil {
		return err
	}

	i.insta = insta
	return nil
}

func (i *Instagram) Upload(ctx context.Context, file io.Reader, name, caption string) error {
	_, err := i.insta.Upload(
		&goinsta.UploadOptions{
			File:                 file,
			Caption:              caption,
			IsStory:              false,
			MuteAudio:            false,
			DisableComments:      false,
			DisableLikeViewCount: false,
			DisableSubtitles:     false,
		},
	)
	if err != nil {
		return errors.Join(ErrUpload, err)
	}

	return nil

}
