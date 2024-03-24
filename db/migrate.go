package db

import (
	"log"
	"os"

	"github.com/nnnnn81/stampy-be/model"
	"github.com/nnnnn81/stampy-be/util"
)

func Migrate() {

	DB.Exec("DROP TABLE IF EXISTS notices")
	DB.Exec("DROP TABLE IF EXISTS letters")
	DB.Exec("DROP TABLE IF EXISTS stamps")
	DB.Exec("DROP TABLE IF EXISTS users")
	DB.Exec("DROP TABLE IF EXISTS stampcards")

	DB.AutoMigrate(&model.Notice{})
	DB.AutoMigrate(&model.Letter{})
	DB.AutoMigrate(&model.Stamp{})
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Stampcard{})

	// test user 作成
	hashedPass, _ := util.HashPassword(os.Getenv("STAMPY_PASSWORD"))
	user := model.User{
		Username:       "stampy",
		Email:          "stampy@gmail.com",
		HashedPassword: hashedPass,
		AvatarUrl:      "https://stampy.com",
	}
	DB.Create(&user)
	log.Print("[INFO] DB Migrated!")
}
