package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host string
	User string
	Password string
	DBname string
	Port int
}

func NewGORMDB(cfg DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d "+
		"sslmode=disable TimeZone=Europe/Moscow",
		cfg.Host, cfg.User, cfg.Password, cfg.DBname, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
