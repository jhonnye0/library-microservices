package hello

import (
	"context"
	"os"

	"github.com/go-kit/log"
)

type helloService struct{}

func NewService() Service { return &helloService{} }

func (w *helloService) Hello(_ context.Context, name string) string {

	if name != "" {
		return "Hello " + name
	}
	return "Hello World!"
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
