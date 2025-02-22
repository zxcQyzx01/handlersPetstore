package config

type Config struct {
	ServerPort string
	DBConfig   DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewConfig() *Config {
	return &Config{
		ServerPort: "8080",
		DBConfig: DBConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "postgres",
			DBName:   "petstore",
		},
	}
}
