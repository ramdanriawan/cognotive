package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	// host := os.Getenv("MYSQL_HOST")
	// port := os.Getenv("MYSQL_PORT")
	// dbName := os.Getenv("MYSQL_DBNAME")
	// username := os.Getenv("MYSQL_USER")
	// password := os.Getenv("MYSQL_PASSWORD")

	host := "63.250.44.168"
	port := "3306"
	dbName := "cognotive"
	username := "root"
	password := "M@utauaja982"

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8&parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Tidak dapat terkoneksi dengan database, periksa config!")
	}

	return db
}
