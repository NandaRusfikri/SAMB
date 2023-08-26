package pkg

import (
	"SAMB-BE/schemas"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func GodotEnv(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Errorln("Error loading .env file")
	}

	env := make(chan string, 1)
	//fmt.Println(os.Getenv("GO_ENV"))

	if os.Getenv("GO_ENV") != "production" {
		godotenv.Load(filepath.Join(".env"))
		env <- os.Getenv(key)
	} else {
		env <- os.Getenv(key)
	}

	return <-env
}

func Environment(path string) (config schemas.SchemaEnvironment) {
	err := godotenv.Load(path)
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Read environment variables from docker-compose.yml
	config.DB_HOST = os.Getenv("DB_HOST")
	config.DB_PORT = os.Getenv("DB_PORT")
	config.DB_USER = os.Getenv("DB_USER")
	config.DB_NAME = os.Getenv("DB_NAME")
	config.DB_PASS = os.Getenv("DB_PASS")
	config.DB_SSLMODE = os.Getenv("DB_SSLMODE")

	config.SMTP_HOST = os.Getenv("SMTP_HOST")
	config.SMTP_PORT = os.Getenv("SMTP_PORT")
	config.SMTP_NAME = os.Getenv("SMTP_NAME")
	config.SMTP_EMAIL = os.Getenv("SMTP_EMAIL")
	config.SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")

	config.TIMEZONE = os.Getenv("TIMEZONE")
	config.VERSION = os.Getenv("VERSION")
	config.REST_PORT = os.Getenv("REST_PORT")
	config.GO_ENV = os.Getenv("GO_ENV")
	config.SWAGGER_HOST = os.Getenv("SWAGGER_HOST")
	config.JWT_SECRET = os.Getenv("JWT_SECRET")
	config.Minio_Host = os.Getenv("MINIO_HOST")
	config.Minio_SSL = os.Getenv("MINIO_SSL")
	config.Minio_SecretKey = os.Getenv("MINIO_SECRET_KEY")
	config.Minio_AccessKey = os.Getenv("MINIO_ACCESS_KEY")
	config.Minio_Domain = os.Getenv("MINIO_DOMAIN")

	//fmt.Printf("Environment sini: %+v\n", config)

	return config
}
