package config

type (
	AppConfig struct {
		Host string
		Port int
	}

	DatabaseConfig struct {
		Driver string
		DSN    string
	}

	Config struct {
		App      AppConfig
		Database DatabaseConfig
	}
)
