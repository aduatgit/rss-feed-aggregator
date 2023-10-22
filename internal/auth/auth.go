package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(r *http.Request) (string, error) {
	s := r.Header.Get("Authorization")
	apiKey := strings.Split(s, " ")
	if apiKey[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}
	return apiKey[1], nil
}
