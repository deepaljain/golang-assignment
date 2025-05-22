package utils

import (
	"github.com/google/uuid"
	"time"
)

func GenerateTransactionID() string {
	return uuid.NewString()
}

func GetCurrentTimestamp() string {
	date := time.Now()
	return date.Format("02-01-2025 15:04:05")
}