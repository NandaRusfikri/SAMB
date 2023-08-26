package schemas

import "github.com/gin-contrib/cors"

type AppConfig struct {
	Debug        bool
	Host         string
	Port         string
	UseTLS       bool
	CertFilePath string
	KeyFilePath  string
	Cors         *cors.Config
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Dialect  string
	TimeZone string
	SSLMode  string
}
