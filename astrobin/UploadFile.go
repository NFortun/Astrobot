package astrobin

import (
	"context"
	"fmt"
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
	for _, uploader := range uploaders {
		file := params.UpFile
		if err := uploader.Connect(ctx); err != nil {
			msg := fmt.Sprintf("fail to connect: %s", err.Error())
			return api.NewUploadImageDefault(http.StatusInternalServerError).WithPayload(&models.Error{Message: &msg})
		}

		if err := uploader.Upload(ctx, file, params.Name, *params.Caption); err != nil {
			msg := fmt.Sprintf("fail to upload: %s", err.Error())
			return api.NewUploadImageDefault(http.StatusInternalServerError).WithPayload(&models.Error{Message: &msg})
		}
	}

	return api.NewUploadImageNoContent()
}
