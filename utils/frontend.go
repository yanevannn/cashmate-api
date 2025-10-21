package utils

import "os"

func GetFrontendURL() string {
	frontendURL := os.Getenv("FRONTEND_URL")
	return frontendURL
}