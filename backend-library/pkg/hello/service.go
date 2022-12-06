package hello

import (
	"context"
)

type Service interface {
	// Get the list of all documents
	Hello(ctx context.Context, name string) string
}
