package config

import (
	"os"
)

var (
	SERVER_PORT     = os.Getenv("SERVER_PORT")
	DATABASE_DRIVER = os.Getenv("DATABASE_DRIVER")
	DATABASE_HOST   = os.Getenv("DATABASE_HOST")
	DATABASE_NAME   = os.Getenv("DATABASE_NAME")
	DATABASE_PASS   = os.Getenv("DATABASE_PASS")
	DATABASE_PORT   = os.Getenv("DATABASE_PORT")
	DATABASE_USER   = os.Getenv("DATABASE_USER")
	ENVIRONMENT     = os.Getenv("ENVIRONMENT")
	ERROR_FILE      = os.Getenv("ERROR_FILE")
)
