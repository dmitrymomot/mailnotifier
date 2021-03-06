package mailnotifier

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var config *Config

// SetupConfig setup configuration method
func SetupConfig(c *Config) {
	config = c
}

// InitFromDotenvFile config setup from dotenv file
func InitFromDotenvFile() (err error) {
	err = godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return
	}

	port, _ := strconv.Atoi(os.Getenv("MAIL_SMTP_PORT"))

	SetupConfig(&Config{
		Host:     os.Getenv("MAIL_SMTP_HOST"),
		Port:     port,
		Username: os.Getenv("MAIL_SMTP_USERNAME"),
		Password: os.Getenv("MAIL_SMTP_PASSWORD"),

		SenderEmail: os.Getenv("MAIL_FROM_ADDR"),
		SenderName:  os.Getenv("MAIL_FROM_NAME"),

		ProductName: os.Getenv("APP_NAME"),
		ProductLink: os.Getenv("WEBSITE_URL"),

		UnsubscribeLink:   os.Getenv("MAIL_UNSUBSCRIBE_LINK"),
		RemoveEmailLink:   os.Getenv("MAIL_REMOVE_LINK"),
		ConfirmEmailLink:  os.Getenv("MAIL_CONFIRM_LINK"),
		ResetPasswordLink: os.Getenv("MAIL_RESET_LINK"),
	})

	return
}

// Config email client configuration
type Config struct {
	SenderEmail string
	SenderName  string

	ProductName string
	ProductLink string

	Host     string
	Port     int
	Username string
	Password string

	UnsubscribeLink   string
	RemoveEmailLink   string
	ConfirmEmailLink  string
	ResetPasswordLink string
}

func (c *Config) getSenderName() (name string) {
	if c.SenderName != "" {
		name = c.SenderName
	}
	if c.ProductName != "" && name != "" {
		name = name + " from " + c.ProductName
	}
	if c.ProductName != "" && name == "" {
		name = c.ProductName
	}
	return
}
