package config

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
}

func Load() Config {
	cfg := Config{}

	cfg.PostgresHost = "localhost"
	cfg.PostgresPort = "5432"
	cfg.PostgresUser = "sevinch"
	cfg.PostgresPassword = "1218"
	cfg.PostgresDB = "shop"

	return cfg
}
