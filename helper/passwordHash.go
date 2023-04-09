package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func GenerateHashedPassword(password string) (string, error) {

	passwordByte := []byte(password)

	hash := sha256.Sum256(passwordByte)

	hashString := hex.EncodeToString(hash[:])

	return hashString, nil
}

func CompareHashedPassword(hashed string, password string) bool {

	passwordByte := []byte(password)

	hashByte, err := hex.DecodeString(hashed)
	if err != nil {
		log.Println(err)
		return false
	}

	newHash := sha256.Sum256(passwordByte)

	if hex.EncodeToString(newHash[:]) == hex.EncodeToString(hashByte) {
		return true
	} else {
		return false
	}
}

// Dapatkan representasi byte dari password
// Buat hash baru dengan menggunakan SHA256
// Konversi hash menjadi string heksadesimal

// Dapatkan representasi byte dari password
// Konversi hash dari string heksadesimal menjadi byte
// Buat hash baru dengan menggunakan SHA256
// Bandingkan hash baru dengan hash yang sudah tersimpan
