package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
)

type Config struct {
	Port               int
	DBConnectionString string
	args               *args
}

type args struct {
	Port     string `json:"APP_PORT" env:"PORT"`

	DBPort     string `json:"DB_PORT" env:"DB_PORT"`
	DBServer   string `json:"DB_SERVER" env:"DB_SERVER"`
	DBInstance string `json:"DB_INSTANCE" env:"DB_NAME"`
	DBUsername string `json:"DB_USERNAME" env:"DB_USERNAME"`
	DBPassword string `json:"DB_PASSWORD" env:"DB_PASSWORD"`
	DBTimeout  string `json:"DB_TIMEOUT" env:"DB_TIMEOUT"`
}

func LoadConfigFromFile(configFilePath string) (*Config, []error) {
	args := &args{}

	contents, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, []error{fmt.Errorf("failed to open configuration file '%s': %s", configFilePath, err)}
	}

	err = json.Unmarshal(contents, &args)
	if err != nil {
		return nil, []error{fmt.Errorf("error parsing JSON configuration: %s", err)}
	}

	return newConfig(args)
}

func LoadConfigFromEnv() (*Config, []error) {
	args := &args{}
	// TODO:: Read args from environment variable
	return newConfig(args)
}

func newConfig(args *args) (*Config, []error) {
	var err error
	var errs []error

	args.setDefaultConfigs()

	connectionString, err := buildConnectionString(
		args.DBServer,
		args.DBPort,
		args.DBInstance,
		args.DBUsername,
		args.DBPassword,
		args.DBTimeout,
	)
	if err != nil {
		errs = append(errs, fmt.Errorf("cannot build connection string: %s", err))
	}
	dbConnectionString := connectionString

	port, err := strconv.Atoi(args.Port)
	if err != nil {
		errs = append(errs, fmt.Errorf("APP_PORT %q is not a number", args.Port))
	}

	config := Config{
		Port:               port,
		DBConnectionString: dbConnectionString,
		args:               args,
	}

	return &config, errs
}

func (conf Config) Loggable() string {
	return fmt.Sprintf("")
}

func (args *args) setDefaultConfigs() {
	args.DBServer = defaultIfNotConfigured(args.DBServer, "localhost")
	args.DBPort = defaultIfNotConfigured(args.DBPort, "5432")
	args.DBTimeout = defaultIfNotConfigured(args.DBTimeout, "30000")
	args.Port = defaultIfNotConfigured(args.Port, "80")
}

func buildConnectionString(dbServer string, dbPort string, dbInstance string, dbUsername string, dbPassword string, dbTimeout string) (string, error) {
	format := "postgres://%s:%s/%s?sslmode=disable&statement_timeout=%s"
	connectionString := fmt.Sprintf(format, dbServer, dbPort, dbInstance, dbTimeout)

	u, err := url.Parse(connectionString)
	if err != nil {
		return "", err
	}

	u.User = url.UserPassword(dbUsername, dbPassword)
	return u.String(), nil
}

func defaultIfNotConfigured(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
