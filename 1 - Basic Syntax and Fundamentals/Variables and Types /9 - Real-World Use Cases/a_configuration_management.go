package main

import "fmt"

// Using structs and custom types to manage application configurations.

// Best Practice: Organize configuration data using nested structs and use custom types for clarity.

type Config struct {
	Server   ServerConfig
	Database DBConfig
}

type ServerConfig struct {
	Port int
	Host string
}

type DBConfig struct {
	User     string
	Password string
	Name     string
}

func loadConfig() Config {
	// Load configuration from file/environment
	return Config{
		Server: ServerConfig{
			Port: 8080,
			Host: "localhost",
		},
		Database: DBConfig{
			User:     "admin",
			Password: "secret",
			Name:     "app_db",
		},
	}
}

func main() {
	config := loadConfig()
	fmt.Printf("Server config: %+v\n", config.Server)
	fmt.Printf("Database config: %+v\n", config.Database)
}
