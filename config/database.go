package config

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "dev", //GET THIS OUT OF THE CODE AND SOMEWHERE ELSE ASAP
		Password: "",
		DBName:   "stuco2020",
	}
	return &dbConfig
}

//DbURL generates a database URL from a database config
func (dbConfig *DBConfig) DbURL() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
