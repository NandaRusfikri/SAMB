package driver

import (
	"SAMB-BE/schemas"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDBSQL(config schemas.SchemaEnvironment) (*gorm.DB, error) {
	logrus.Debug("🔌 Connecting into Database")
	dbHost := config.DB_HOST
	dbUsername := config.DB_USER
	dbPassword := config.DB_PASS
	dbName := config.DB_NAME
	dbPort := config.DB_PORT
	dbSSLMode := config.DB_SSLMODE
	timezone := config.TIMEZONE

	path := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		dbHost, dbUsername, dbPassword, dbName, dbPort, dbSSLMode, timezone)

	db, err := gorm.Open(postgres.Open(path), &gorm.Config{})

	if err != nil {
		defer logrus.Errorln("❌ Error Connect into Database ", err.Error())
		return nil, err
	}

	err = MigrateDB(db)
	if err != nil {
		logrus.Errorln("❌ Error Migrate ", err.Error())
		return nil, err
	}

	logrus.Debug("❌ Connect into Database Success")

	return db, nil
}
