package uploader

import (
	"context"
	"io"
)

type Uploader interface {
	Connect(ctx context.Context) error
	Upload(ctx context.Context, content io.Reader, name, caption string) error
}
