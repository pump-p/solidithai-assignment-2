package config

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

var ESClient *elasticsearch.Client

func LoadConfig() Config {
	// Setup Elasticsearch client
	var err error
	ESClient, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	})

	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	log.Println("Elasticsearch client initialized successfully")

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
