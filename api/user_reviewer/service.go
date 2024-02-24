package api

import (
	"errors"
	"os"
)

// check service admin
func CheckServiceUserReviewer() error {
	serviceAdmin := os.Getenv("SERVICE_REVIEWER")
	if serviceAdmin == "" {
		return errors.New("service user reviewer is empty")
	}
	return nil
}
