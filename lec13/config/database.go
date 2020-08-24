package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//DB ...
var DB *gorm.DB

//DBConfig ...
type DBConfig struct {
	User     string
	Password string
	DBname   string
	SSLmode  string
}

//BuildDBConfig ...
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		User:     "postgres",
		Password: "1",
		DBname:   "userproject",
		SSLmode:  "disable",
	}

	return &dbConfig
}

//DbURI ...
func DbURI(dbConfig *DBConfig) string {
	return fmt.Sprintf("user=%v password=%v dbname=%v sslmode=%v",
		dbConfig.User, dbConfig.Password, dbConfig.DBname, dbConfig.SSLmode)
}
