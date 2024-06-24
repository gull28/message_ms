package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Dsn          string
	Smtp         Smtp
	Sms          Sms
	CodeSettings CodeSettings
}

type Smtp struct {
	Host     string
	Port     string
	Username string
	Password string
}

type Sms struct {
	APIKey   string
	SenderID string
}

type CodeSettings struct {
	Length      int
	Expiry      int
	Attempts    int
	Resends     int
	ResendTimer int
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Warning: Invalid value for %s: %s. Using default: %d\n", key, valueStr, defaultValue)
		return defaultValue
	}
	return value
}

func getEnvAsString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Error loading .env file")
	}
}

func getDsn() string {
	loadEnv()

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbUrl := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName)

	return dbUrl
}

func getSmtp() Smtp {
	loadEnv()

	return Smtp{
		Host:     getEnvAsString("SMTP_HOST", "localhost"),
		Port:     getEnvAsString("SMTP_PORT", "1025"),
		Username: getEnvAsString("SMTP_USERNAME", "user"),
		Password: getEnvAsString("SMTP_PASSWORD", "123"),
	}
}

func getSms() Sms {
	loadEnv()

	return Sms{
		APIKey:   getEnvAsString("SMS_API_KEY", ""),
		SenderID: getEnvAsString("SMS_SENDER_ID", ""),
	}
}

func getCodeSettings() CodeSettings {
	loadEnv()

	return CodeSettings{
		Length:      getEnvAsInt("CODE_LENGTH", 6),
		Expiry:      getEnvAsInt("CODE_EXPIRY", 300),
		Attempts:    getEnvAsInt("CODE_ATTEMPTS", 3),
		Resends:     getEnvAsInt("CODE_RESENDS", 2),
		ResendTimer: getEnvAsInt("CODE_RESEND_TIMER", 60),
	}
}

func LoadConfig() Config {
	Dsn := getDsn()

	Smtp := getSmtp()

	Sms := getSms()

	CodeSettings := getCodeSettings()

	return Config{
		Dsn:          Dsn,
		Smtp:         Smtp,
		CodeSettings: CodeSettings,
		Sms:          Sms,
	}
}
