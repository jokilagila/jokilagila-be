package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/jokilagila/jokilagila-be/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var Database *gorm.DB

func PostgresConfig() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Gagal untuk memuat file environment variables:", err)
		return nil, err
	}

	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	if err != nil {
		log.Println("Gagal untuk terhubung ke database:", err)
		return nil, err
	}

	Database = db

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Println("Gagal untuk melakukan migrasi:", err)
		return nil, err
	}

	log.Println("Berhasil terhubung ke PostgreSQL")

	return db, nil
}
