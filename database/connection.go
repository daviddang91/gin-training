package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB

type DBConnection struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
	Timezone string
}

func NewDBConnection() DBConnection {
	return DBConnection{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Timezone: os.Getenv("DB_TIMEZONE"),
	}
}

func (c DBConnection) ToConnectionString() string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		c.Host,
		c.Username,
		c.Password,
		c.Database,
		c.Port,
		c.SSLMode,
		c.Timezone,
	)
}

func SetupDatabaseConnection() {
	dsn := NewDBConnection().ToConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Instance = db
}

func CloseDatabaseConnection() {
	db, err := Instance.DB()
	if err != nil {
		panic(err)
	}
	db.Close()
}
