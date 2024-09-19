package config

type Config struct {
	DatabaseURL string
}

func LoadConfig() *Config {
	// Load configuration settings from environment variables or a configuration file
	// For example:
	// databaseURL := os.Getenv("DATABASE_URL")
	// return &Config{
	//     DatabaseURL: databaseURL,
	// }
	return &Config{
		DatabaseURL: "postgres://username:password@localhost:5432/database",
	}
}