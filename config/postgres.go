package config

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/jokilagila/jokilagila-be/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB
	Err      error
	Once     sync.Once
)

func PostgresConfig() (*gorm.DB, error) {

	Once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Gagal memuat file .env:", err)
		}

		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, user, password, dbname)
		Database, Err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if Err != nil {
			log.Fatal("Gagal terhubung ke database:", Err)
		}

		sqlDB, err := Database.DB()
		if err != nil {
			log.Fatal("Gagal mendapatkan database instance:", err)
		}

		sqlDB.SetMaxIdleConns(20)
		sqlDB.SetMaxOpenConns(50)
		sqlDB.SetConnMaxLifetime(60 * time.Minute)

		if err := Database.AutoMigrate(&model.User{}); err != nil {
			log.Fatal("Gagal melakukan migrasi database:", err)
		}

	})
	return Database, Err

}
