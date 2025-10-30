package config

import "os"

type Config struct {
	MailAPIKey      string
	Database        databaseConfig
	TransportConfig transportConfig
	LogLevel        string `env:"LOG_LEVEL" envDefault:"info"`
}
type databaseConfig struct {
	ReadHost      string `env:"DATABASE_HOST_READ,required"`
	ReadUser      string `env:"DATABASE_USER_READ,required"`
	ReadPassword  string `env:"DATABASE_PASSWORD_READ,required"`
	WriteHost     string `env:"DATABASE_HOST_WRITE,required"`
	WriteUser     string `env:"DATABASE_USER_WRITE,required"`
	WritePassword string `env:"DATABASE_PASSWORD_WRITE,required"`
	Name          string `env:"DATABASE_NAME"`
	Migrate       struct {
		SkipNotFoundError bool `env:"DATABASE_MIGRATE_SKIP_NOT_FOUND_ERROR"`
	}
}
type transportConfig struct {
	Address        string `env:"TRANSPORT_ADDRESS" envDefault:":8080"`
	HealthAddress  string `env:"TRANSPORT_HEALTH_ADDRESS" envDefault:":8081"`
	MetricsAddress string `env:"TRANSPORT_METRICS_ADDRESS" envDefault:":8082"`
	MaxBodySize    int    `env:"TRANSPORT_MAX_BODY_SIZE" envDefault:"10485760"`
}

func LoadConfig() *Config {
	return &Config{
		MailAPIKey: os.Getenv("MAIL_API_KEY"),
	}
}
