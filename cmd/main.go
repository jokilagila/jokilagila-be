package main

import (
	"log"

	"github.com/jokilagila/jokilagila-be/config"
	"github.com/jokilagila/jokilagila-be/internal/router"
	"github.com/jokilagila/jokilagila-be/seed"
)

func main() {
	db, err := config.PostgresConfig()
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	if err := seed.UserSeed(); err != nil {
		log.Fatalf("Gagal melakukan seeding user: %v", err)
	}

	db.Logger.LogMode(1)

	r := router.SetupRoutes()
	log.Println("Server berjalan di http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
