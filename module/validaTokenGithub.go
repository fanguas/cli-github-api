package module

import (
	"os"
)

func ValidateTokenGithub() string {
	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		return "No se encontró el token de GitHub. Por favor, establece la variable de entorno GITHUB_TOKEN."
	}

	return token
}