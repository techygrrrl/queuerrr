package api_utils

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func Authenticate(r *http.Request) error {
	token := os.Getenv("AUTH_TOKEN")
	if token == "" {
		return fmt.Errorf("missing AUTH_TOKEN")
	}

	userToken, err := parseAuthTokenFromRequest(r)

	if err != nil {
		return err
	}

	tokenBytes := []byte(token)
	userTokenBytes := []byte(*userToken)

	result := subtle.ConstantTimeCompare(tokenBytes, userTokenBytes)
	if result == 0 {
		return fmt.Errorf("unauthorized")
	}

	return nil
}

func parseAuthTokenFromRequest(r *http.Request) (*string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("no auth header")
	}

	splits := strings.Split(authHeader, " ")
	if len(splits) != 2 {
		return nil, fmt.Errorf("invalid authorization header")
	}

	prefix := splits[0]
	if prefix != "Bearer" {
		return nil, fmt.Errorf("invalid authorization type")
	}

	return &splits[1], nil
}
