package main

import (
	"fmt"
	"net/http"

	"github.com/jokilagila/jokilagila-be/config"
	"github.com/jokilagila/jokilagila-be/seed"
)

func main() {
	database, err := config.PostgresConfig()
	if err != nil {
		fmt.Println("Gagal untuk terhubung ke database:", err)
		return
	}

	if err := seed.UserSeed(); err != nil {
		fmt.Println("Gagal melakukan seeding user:", err)
		return
	}

	database.Logger.LogMode(1)

	fmt.Println("Server berjalan di :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}
