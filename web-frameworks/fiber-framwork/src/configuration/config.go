package configuration

import "os"

func GetSecret() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}

	return secret
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}
