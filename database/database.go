package database

import (
	"fmt"
	"log"
	"time"

	"api-books-go/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	// TODO: Ver variaveis de ambiente
	strConexao := "host=localhost port=25432 user=admin dbname=books sslmode=disable password=123456"

	database, err := gorm.Open(postgres.Open(strConexao), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	}

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
	migrations.RunMigrations(db)
}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}
func GetDatabase() *gorm.DB {
	return db
}
