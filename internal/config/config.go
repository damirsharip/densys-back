package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Build information -ldflags .
var (
	version    = "dev"
	commitHash = "-"
)

// Project - contains all parameters project information.
type Project struct {
	Debug       bool   `yaml:"debug"`
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	ServiceName string `yaml:"serviceName"`
	Version     string
	CommitHash  string
}

type Rest struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type PgxPool struct {
	// MaxConnLifetime is the duration since creation after which a connection will be automatically closed.
	MaxConnLifetime int64 `yaml:"max_conn_lifetime"`
	// MaxConnIdleTime is the duration after which an idle connection will be automatically closed by the health check.
	MaxConnIdleTime int64 `yaml:"max_conn_idle_time"`
	// MaxConns is the maximum size of the pool. The default is the greater of 4 or runtime.NumCPU().
	MaxConns int32 `yaml:"max_conns"`
	// MinConns is the minimum size of the pool. The health check will increase the number of connections to this
	// amount if it had dropped below.
	MinConns int32 `yaml:"min_conns"`
}

type Database struct {
	Dsn string `yaml:"dsn"`
}

type Token struct {
	SymmetricKey string `yaml:"symmetric_key"`
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Project  Project  `yaml:"project"`
	Rest     Rest     `yaml:"rest"`
	Database Database `yaml:"database"`
	Token    Token    `yaml:"token"`
	PgxPool  PgxPool  `yaml:"pgxpool"`
}

var cfg *Config

// ReadConfigYML - read configurations from file and init instance Config.
func ReadConfigYML(configYML string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(configYML)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&cfg); err != nil {
		return err
	}

	cfg.Project.Version = version
	cfg.Project.CommitHash = commitHash

	return nil
}

func Get() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}
