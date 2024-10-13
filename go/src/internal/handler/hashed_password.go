package handler

import "github.com/kynmh69/mormorare/pkg/env"

func HashPassword(password string) string {
	// Hash password
	appKey := env.FindEnv("APP_KEY")

	return password
}
