package database

import (
	"Mou1ght-Server/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initPostgres() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Conf.DBAddr, config.Conf.DBName, config.Conf.DBPassword, config.Conf.DatabaseName, config.Conf.DBPort)
	postgresDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return postgresDB
}
