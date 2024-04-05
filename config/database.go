package config

import (
	"boilerplate/helper"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type DBConfig struct {
	Type     string
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func InitDB(dbConfig DBConfig) {
	var (
		dialect gorm.Dialector
		dsn     string
		err     error
	)

	switch dbConfig.Type {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.Username,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Database,
		)
		dialect = mysql.Open(dsn)
	case "pgsql":
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Username,
			dbConfig.Password,
			dbConfig.Database,
		)
		dialect = postgres.Open(dsn)
	default:
		helper.LogError(fmt.Errorf("Unknown database type: %s", dbConfig.Type))
	}

	DB, err = gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		helper.LogError(err)
	}
}
