package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	// gorm postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/nguyenbt456/todolist-go/util"
)

const (
	pgHost     = "127.0.0.1"
	pgPort     = "5432"
	pgUser     = "nhn"
	pgPassword = "123456"
	pgName     = "todolist"
	pgSSLMode  = "disable"
)

// PostgresConfig constains postgres info to connect postgres database
type postgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

var db *gorm.DB

// GetDB return postgres database
func GetDB() *gorm.DB {
	return db
}

// ConnectPostgresDB connect to postgres database
func ConnectPostgresDB() (*gorm.DB, error) {
	cfg := postgresConfig{
		Host:     util.EVString("POSTGRES_HOST", pgHost),
		Port:     util.EVString("POSTGRES_PORT", pgPort),
		User:     util.EVString("POSTGRES_USER", pgUser),
		Password: util.EVString("POSTGRES_PASSWORD", pgPassword),
		Name:     util.EVString("POSTGRES_NAME", pgName),
		SSLMode:  util.EVString("POSTGRES_SSLMODE", pgSSLMode),
	}

	cfgString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode)
	log.Println(cfgString)

	var err error
	db, err = gorm.Open("postgres", cfgString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// DisconnectPostgresDB disconnect to postgres database
func DisconnectPostgresDB(db *gorm.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	db = nil

	return nil
}
