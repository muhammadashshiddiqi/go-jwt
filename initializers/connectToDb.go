package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=" + os.Getenv("HOSTDB") + " user=" + os.Getenv("USERDB") + " password=" + os.Getenv("PASSDB") + " dbname=" + os.Getenv("DBNAME") + " port=" + os.Getenv("PORTDB") + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to Connect DB")
	}
}
