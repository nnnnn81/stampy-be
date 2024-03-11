package db

import (
	"log"

	"github.com/nnnnn81/stampy-be/model"
)

func Migrate() {

	DB.Exec("DROP TABLE IF EXISTS notices")
	DB.Exec("DROP TABLE IF EXISTS stamps")
	DB.Exec("DROP TABLE IF EXISTS users")
	DB.Exec("DROP TABLE IF EXISTS stampcards")

	DB.AutoMigrate(&model.Notice{})
	DB.AutoMigrate(&model.Stamp{})
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Stampcard{})

	log.Print("[INFO] DB Migrated!")
}
