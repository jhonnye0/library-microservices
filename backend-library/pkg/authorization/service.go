package authorization

import (
	"context"
)

type Service interface {
	// Get the list of all documents
	Login(ctx context.Context, username, password string) (string, error)
	Logout(ctx context.Context, token string) error
}
