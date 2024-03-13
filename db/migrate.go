package db

import (
	"log"
	"os"

	"github.com/nnnnn81/stampy-be/model"
	"github.com/nnnnn81/stampy-be/util"
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

	// test user 作成
	hashedPass, _ := util.HashPassword(os.Getenv("TESTUSER_PASSWORD"))
	user := model.User{
		Username:       "user",
		Email:          "user@teamb.com",
		HashedPassword: hashedPass,
		AvaterUrl:      "https://hogheoge.com",
	}
	DB.Create(&user)
	log.Print("[INFO] DB Migrated!")
}
