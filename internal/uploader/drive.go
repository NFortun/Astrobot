package uploader

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type Drive struct {
	srv       *drive.Service
	directory string
}

func (d *Drive) Connect(ctx context.Context) error {
	configFile := os.Getenv("DRIVE_CONFIG")
	if configFile == "" {
		return fmt.Errorf("fail to get DRIVE_CONFIG")
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	config := struct {
		Directory string `json:"directory"`
	}{}

	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}

	srv, err := drive.NewService(ctx, option.WithCredentialsFile(configFile), option.WithScopes("https://www.googleapis.com/auth/drive.file"))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	d.srv = srv
	d.directory = config.Directory
	return nil
}

func (d *Drive) Upload(ctx context.Context, file io.Reader, name, caption string) error {
	_, err := d.srv.Files.Create(&drive.File{
		Description:     caption,
		Name:            name,
		Parents:         []string{d.directory},
		WritersCanShare: true,
	}).Media(file).Do()
	if err != nil {
		return errors.Join(ErrUpload, err)
	}

	return nil
}
