package api_utils

import (
	"net/http"
)

func Authenticate(r *http.Request) error {
	// todo: actually authenticate
	return nil
	// return fmt.Errorf("unauthorized")
}
