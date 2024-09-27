package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"toDoListRestApi/src/internal/configs"
	"toDoListRestApi/src/internal/domain"
)

func NewClient(dbModel configs.DbInitModel) *gorm.DB {
	dsn := "host=" + dbModel.DbHost +
		" user=" + dbModel.DbUser +
		" password=" + dbModel.DbPassword +
		" dbname=" + dbModel.DbName +
		" port=" + dbModel.DbPort +
		" sslmode=disable"

	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	dbMigrateErr := db.AutoMigrate(&domain.Todo{})
	if dbMigrateErr != nil {
		log.Fatal(dbMigrateErr)
	}

	return db
}
