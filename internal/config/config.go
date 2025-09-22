package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logging  LoggingConfig
}

type ServerConfig struct {
	Port int
	Host string
}

type DatabaseConfig struct {
	Name     string
	Host     string
	Port     int
	Username string
	Password string
}

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
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %s", err)
		}

		fmt.Println("No config file found - using defaults and environment variables")
	}

	config := &Config{}
	err := viper.Unmarshal(config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %s", err)
	}

	return config, nil
}

func setDefaults() {
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.host", "0.0.0.0")

	viper.SetDefault("database.name", "in_memory_db.db")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)

	viper.SetDefault("logging.level", string(LogLevelInfo))
	viper.SetDefault("logging.format", "json")       // or "console" for development
	viper.SetDefault("logging.disablecaller", false) // adds file:line to logs
}
