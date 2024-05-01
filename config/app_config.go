package config

type ApplicationConfig struct {
	AmqpConfig  AmqpConfig
	Application AppConfig
}

type AmqpConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Protocol string
}

type AppConfig struct {
	Host string
	Port int
}
