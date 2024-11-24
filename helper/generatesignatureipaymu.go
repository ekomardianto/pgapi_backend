package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"time"
)

// Fungsi untuk hash SHA-256
func hashSHA256(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// Fungsi untuk generate HMAC-SHA256
func generateHMACSHA256(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
func GenerateSignature(requestBody []byte, method string) (string, int64, string) {
	// Generate Signature
	timestamp := time.Now().UnixMilli()
	apiKey := os.Getenv("IPAY_API_KEY")   // Ambil API Key dari environment variable
	va := os.Getenv("IPAY_VA")            // Virtual Account dari Environtmen
	hashedBody := hashSHA256(requestBody) // Hashing body menggunakan SHA-256
	stringToSign := method + ":" + va + ":" + hashedBody + ":" + apiKey
	signature := generateHMACSHA256(stringToSign, apiKey) // Generate HMAC-SHA256 signature

	return signature, timestamp, va
}
