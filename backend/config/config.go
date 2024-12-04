package config

import (
	"fmt"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func LoadConfig() Config {
	return Config{
		DBUsername: "pump",                    // Your database username
		DBPassword: "pump",                    // Your database password
		DBName:     "solidithai_assignment_2", // Your database name
		DBHost:     "127.0.0.1",
		DBPort:     "3306", // Default MySQL port in XAMPP
	}
}

func GetDSN(config Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
}
