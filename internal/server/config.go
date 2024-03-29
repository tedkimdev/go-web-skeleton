package server

import (
	"fmt"
	"os"
	"strings"

	"github.com/tedkimdev/go-web-skeleton/internal/server/database"
)

const (
	DefaultPort = 8080
	DefaultAddr = "0.0.0.0:8080"

	DefaultDBHost     = "postgres"
	DefaultDBPort     = "5432"
	DefaultDBUser     = "postgres"
	DefaultDBPassword = "postgres"
	DefaultDBName     = "postgres"
	DefaultDBSSLMode  = "disable"
)

type Config struct {
	Addr string
	Port string
	DB   *database.Config
}

func NewConfig() (*Config, error) {
	Port := os.Getenv("PORT")
	Addr := os.Getenv("ADDR")

	AddrWithPort := fmt.Sprintf("%s:%s", Addr, Port)

	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")
	SSLMode := os.Getenv("DB_SSL_MODE")

	config := Config{
		Addr: AddrWithPort,
		Port: Port,
		DB: &database.Config{
			DBHost:     DBHost,
			DBPort:     DBPort,
			DBUser:     DBUser,
			DBPassword: DBPassword,
			DBName:     DBName,
			SSLMode:    SSLMode,
		},
	}

	config.ensureDefaultValue()
	return &config, nil
}

func (c *Config) ensureDefaultValue() {
	if c.Port == "" {
		c.Port = fmt.Sprintf("%d", DefaultPort)
	}
	if c.Addr == ":" || len(strings.Split(c.Addr, ":")) < 2 {
		c.Addr = DefaultAddr
	}

	if c.DB.DBHost == "" {
		c.DB.DBHost = DefaultDBHost
	}
	if c.DB.DBPort == "" {
		c.DB.DBPort = DefaultDBPort
	}
	if c.DB.DBUser == "" {
		c.DB.DBUser = DefaultDBUser
	}
	if c.DB.DBPassword == "" {
		c.DB.DBPassword = DefaultDBPassword
	}
	if c.DB.DBName == "" {
		c.DB.DBName = DefaultDBName
	}
	if c.DB.SSLMode == "" {
		c.DB.SSLMode = DefaultDBSSLMode
	}
}
