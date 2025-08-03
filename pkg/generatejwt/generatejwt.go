package generatejwt

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateJWT(email string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memuat file .env:", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		log.Fatal("JWT_SECRET tidak ditemukan di file .env")
	}

	expirationTime := time.Now().Add(30 * time.Hour)

	claims := &jwt.RegisteredClaims{
		Subject:   email,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwtSecret))
	if err != nil {
		log.Fatal("Gagal membuat token JWT:", err)
	}

	log.Println("Token JWT berhasil dibuat untuk email:", email)
	return tokenString
}
