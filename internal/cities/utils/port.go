package utils

import "os"

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3010"
	}
	return ":" + port
}
