package main

import (
	"fmt"
	"github.com/jokilagila/jokilagila-be/config"
	"net/http"
)

func main() {
	database, err := config.PostgresConfig()
	if err != nil {
		fmt.Println("Gagal untuk terhubung ke database:", err)
		return
	}
	database.Logger.LogMode(1)
	http.ListenAndServe(":8080", nil)
}
