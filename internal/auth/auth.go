package auth

import (
	"errors"
	"net/http"
	"strings"
)

/*
request header
Authorization : API_KEY {64hexadecimal}
*/
func GetApiKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no Header found")
	}
	values := strings.Split(val, " ")

	if len(values) != 2 {
		return "", errors.New("malformed auth header")
	}
	if values[0] != "API_KEY" {
		return "", errors.New("malformed 1st part of auth header")
	}

	return values[1], nil
}
