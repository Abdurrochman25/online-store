package config

import (
	"github.com/Abdurrochman25/online-store/internal/util"
)

type Config struct {
	Database Database
}

func NewConfig() Config {
	util.LoadEnv(".env")

	return Config{
		Database: Database{
			DatabaseName: util.GetEnv("PSQL_DBNAME", "postgres"),
			Host:         util.GetEnv("PSQL_HOST", "postgres"),
			Port:         util.GetEnvAsInt("PSQL_PORT", 5432),
			Username:     util.GetEnv("PSQL_USER", "admin"),
			Password:     util.GetEnv("PSQL_PASS", "admin"),
			Option: map[string]string{
				"sslmode": util.GetEnv("PSQL_SSLMODE", "disable"),
			},
		},
	}
}
