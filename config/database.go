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

func BuildDBConfig(host string, port int, user string, pass string, name string) *DBConfig {
	dbConfig := DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: pass,
		DBName:   name,
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
