package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config holds the entire configuration for the application.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logging  LoggingConfig
}

// ServerConfig holds server-related configuration.
type ServerConfig struct {
	Port             int
	Host             string
	TimeOutInSeconds int
}

// DatabaseConfig holds database-related configuration.
type DatabaseConfig struct {
	Name     string
	Host     string
	Port     int
	Username string
	Password string
}

// LogLevel defines the level of logging.
type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

type LoggingConfig struct {
	Level         LogLevel `mapstructure:"level"`
	Format        string   `mapstructure:"format"`
	DisableCaller bool     `mapstructure:"disablecaller"` // Adds the caller (file:line) to log entries
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.AutomaticEnv()             // read in environment variables that match
	viper.SetEnvPrefix("GAMESERVER") // prefix for environment variables
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set defaults
	setDefaults()

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}

		fmt.Println("No config file found - using defaults and environment variables")
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return config, nil
}

func setDefaults() {
	defaultServerPort := 8080
	defaultDbPort := 5432
	defaultServerTimeout := 5 // seconds

	viper.SetDefault("server.port", defaultServerPort)
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.timeoutinseconds", defaultServerTimeout) // default timeout for server requests

	viper.SetDefault("database.name", "in_memory_db.db")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", defaultDbPort)

	viper.SetDefault("logging.level", string(LogLevelInfo))
	viper.SetDefault("logging.format", "json")       // or "console" for development
	viper.SetDefault("logging.disablecaller", false) // adds file:line to logs
}
