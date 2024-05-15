package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_User    string
	DB_Pass    string
	DB_Name    string
	DB_Host    string
	DB_Port    string
	SourceData string
}

func NewEmptyConfig() *Config {
	return &Config{}
}

func (c *Config) LoadConfigFromEnv() error {
	godotenv.Load()

	c.DB_User = os.Getenv("DB_USER")
	c.DB_Pass = os.Getenv("DB_PASS")
	c.DB_Name = os.Getenv("DB_NAME")
	c.DB_Host = os.Getenv("DB_HOST")
	c.DB_Port = os.Getenv("DB_PORT")
	c.SourceData = os.Getenv("SOURCE_DATA")

	return nil
}

func (c *Config) LoadConfigWithValues(user, pass, db, host, port, data string) {
	c.DB_User = user
	c.DB_Pass = pass
	c.DB_Name = db
	c.DB_Host = host
	c.DB_Port = port
	c.SourceData = data
}
