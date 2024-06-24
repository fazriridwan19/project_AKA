package config

import (
	"fmt"
	"os"
)

type Config struct {
	dbHost     string
	dbPort     string
	dbName     string
	dbUser     string
	dbPassword string
}

func New() *Config {
	return &Config{
		dbHost:     GetEnv("DB_HOST"),
		dbPort:     GetEnv("DB_PORT"),
		dbName:     GetEnv("DB_NAME"),
		dbUser:     GetEnv("DB_USERNAME"),
		dbPassword: GetEnv("DB_PASSWORD"),
	}
}

func NewTest() *Config {
	os.Setenv("DB_HOST", "192.168.0.103")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "project_aka")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "root")

	return New()
}

func GetEnv(key string) string {
	res := os.Getenv(key)
	if res == "" {
		panic("key " + key + " is empty")
	}
	return res
}

func (c *Config) DatabaseConnection() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", c.dbUser, c.dbPassword, c.dbHost, c.dbPort, c.dbName)
}
