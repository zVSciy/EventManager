package util

import (
	"errors"
	"log"
)

var AUTH_URL = Getenv("AUTH_URL", "https://user-service/auth")

func ValidateAuthHeader(userID, authHeader string) error {
	if authHeader == "" {
		return errors.New("missing auth header")
	}

	err := requestAuth(AUTH_URL, userID, authHeader)
	if err != nil {
		return errors.New("invalid auth")
	}

	return nil
}

func requestAuth(authURL, userID, authHeader string) error {
	log.Println(authURL, userID, authHeader)
	// AUTH LOGIC HERE

	return nil
}
