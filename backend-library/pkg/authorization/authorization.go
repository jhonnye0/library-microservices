package authorization

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type authorizationService struct {
	tokens map[string]time.Time
}

func NewService() Service { return &authorizationService{
	tokens: make(map[string]time.Time),
} }

func (s *authorizationService) Login(ctx context.Context, username, password string) (string, error) {
	if username == "" || password == "" {
		return "", errors.New("invalid username or password")
	}

	if username == "admin" && password == "admin" {
		token := fmt.Sprintf("%s:%s", username, password)
		s.tokens[token] = time.Now().Add(time.Minute * 5)
		return token, nil
	}

	return "", errors.New("invalid username or password")
}

func (s *authorizationService) Logout(ctx context.Context, token string) error {
	if _, ok := s.tokens[token]; ok {
		delete(s.tokens, token)
		return nil
	}

	return errors.New("invalid token")
}
