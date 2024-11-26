package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Upgrade(db *gorm.DB) error {
	// Tambahkan kolom ke tabel
	if err := db.AutoMigrate(
		&User{},
		&RoleUser{},
		&Instansi{},
		&Perusahaan{},
	); err != nil {
		return err
	}
	//*** TRIGGER ***//

	return nil
}
func ConnectDatabase() {
	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")
	var dbUser = os.Getenv("DB_USER")
	var dbPass = os.Getenv("DB_PASS")
	var dbName = os.Getenv("DB_NAME")
	//connect database with user, host, port
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//migrasikan seluruh model Tabel
	if err := db.AutoMigrate(
		&User{},
		&RoleUser{},
		&Instansi{},
		&Perusahaan{},
	); err != nil {
		panic(err)
	}

	// Panggil fungsi Upgrade untuk menambah trigger
	if err := Upgrade(db); err != nil {
		panic(err)
	}
	DB = db
}
