package schemas

type RedisConfig struct {
	Host     string
	Password string
	DB       int
}

type SMTPConfig struct {
	Host     string
	Port     int
	Email    string
	Password string
	Name     string
}

type MinioConfig struct {
	Host          string
	Location      string
	AccessKey     string
	SecretKey     string
	SSL           bool
	ReplaceDomain string
}
type SchemaEnvironment struct {
	DB_USER    string
	DB_PASS    string
	DB_HOST    string
	DB_PORT    string
	DB_NAME    string
	DB_SSLMODE string

	SMTP_HOST     string
	SMTP_PORT     string
	SMTP_EMAIL    string
	SMTP_PASSWORD string
	SMTP_NAME     string

	TIMEZONE        string
	VERSION         string
	REST_PORT       string
	GO_ENV          string
	SWAGGER_HOST    string
	JWT_SECRET      string
	Minio_Host      string
	Minio_Location  string
	Minio_AccessKey string
	Minio_SecretKey string
	Minio_SSL       string
	Minio_Domain    string
}
