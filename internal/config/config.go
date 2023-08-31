package config

import "time"

// Config defines expected yaml structure for conf file.
type Config struct {
	Port                  string `yaml:"port"`
	Debug                 bool   `yaml:"debug"`
	VersionPrefixedRoutes bool   `yaml:"version_prefixed_routes"`
	Installpath           string `yaml:"installpath"`
	Basepath              string `yaml:"basepath"`

	DatabaseConfig struct {
		Engine                 string        `yaml:"engine"`
		ConnectionString       string        `yaml:"connection_string"`
		MaxIdleConns           int           `yaml:"max_idle_conns"`
		MaxOpenConns           int           `yaml:"max_open_conns"`
		ConnMaxLifetime        time.Duration `yaml:"conn_max_lifetime"`
		ConnectionRetriesCount uint          `yaml:"connection_retries_count"`
	} `yaml:"database"`

	Log struct {
		ConsoleLevel   string `yaml:"consolelevel"`
		UseFile        bool   `yaml:"usefile"`
		FileLevel      string `yaml:"filelevel"`
		FilePath       string `yaml:"filepath"`
		FileMaxSize    int    `yaml:"filemaxsize"`
		FileMaxBackups int    `yaml:"filemaxbackup"`
		FileMaxAge     int    `yaml:"filemaxage"`
	} `yaml:"log"`

	Cors struct {
		AllowedOrigins string `yaml:"allowed_origins"`
		AllowedMethods string `yaml:"allowed_methods"`
		AllowedHeaders string `yaml:"allowed_headers"`
	} `yaml:"cors"`

	Caching struct {
		TTL      int `yaml:"default_ttl"`
		ErrorTTL int `yaml:"error_ttl"`
	} `yaml:"caching"`
}
