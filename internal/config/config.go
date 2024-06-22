package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN  string
	Smtp Smtp
}

type Smtp struct {
	host     string
	port     string
	username string
	password string
}

type Sms struct {
}

type CodeSettings struct {
	Length      int
	Expiry      int
	Attempts    int
	Resends     int
	ResendTimer int
}

func getDsn() string {
	dotEnv := godotenv.Load("")

	if dotEnv != nil {
		fmt.Println("Error loading .env file")
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Open a connection to the database
	dbUrl := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName)

	return dbUrl
}

func getSmtp() Smtp {
	dotEnv := godotenv.Load("")

	if dotEnv != nil {
		fmt.Println("Error loading .env file")
	}

	smtpHost := os.Getenv("SMTP_SERVER_HOST")
	smtpPort := os.Getenv("SMTP_SERVER_PORT")
	smtpUsername := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASS")

	return Smtp{
		host:     smtpHost,
		port:     smtpPort,
		username: smtpUsername,
		password: smtpPassword,
	}
}

func LoadConfig() Config {
	DSN := getDsn()

	Smtp := getSmtp()

	// SMS := getSMS()

	// CodeSettings := getCodeSettings()

	return Config{
		DSN:  DSN,
		Smtp: Smtp,
		// CodeSettings: CodeSettings,
		// SMS: SMS,
	}
}
