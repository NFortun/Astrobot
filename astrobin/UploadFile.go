package astrobin

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/NFortun/Astrobot/internal/uploader"
	"github.com/NFortun/Astrobot/models"

	api "github.com/NFortun/Astrobot/restapi/operations"
)

func UploadFile(params api.UploadImageParams) middleware.Responder {
	if params.Name == "" {
		msg := "missing name parameter"
		return api.NewUploadImageDefault(http.StatusBadRequest).WithPayload(&models.Error{Message: &msg})
	}

	ctx := context.Background()
	fmt.Printf("content type: %s\n", params.HTTPRequest.Header.Get("Content-Type"))

	uploaders := []uploader.Uploader{&uploader.Drive{}, &uploader.Instagram{}}
	file := params.UpFile
	data, err := io.ReadAll(file)
	if err != nil {
		msg := fmt.Sprintf("fail to read data: %s", err.Error())
		return api.NewUploadImageDefault(http.StatusInternalServerError).WithPayload(&models.Error{Message: &msg})
	}

	var caption string
	if params.Caption != nil {
		caption = *params.Caption
	}

	for _, uploader := range uploaders {
		reader := bytes.NewReader(data)
		if err := uploader.Connect(ctx); err != nil {
			msg := fmt.Sprintf("fail to connect: %s", err.Error())
			return api.NewUploadImageDefault(http.StatusInternalServerError).WithPayload(&models.Error{Message: &msg})
		}

		if err := uploader.Upload(ctx, reader, params.Name, caption); err != nil {
			msg := fmt.Sprintf("fail to upload: %s", err.Error())
			return api.NewUploadImageDefault(http.StatusInternalServerError).WithPayload(&models.Error{Message: &msg})
		}
	}

	return api.NewUploadImageNoContent()
}
