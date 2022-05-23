package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	envFile = ".env"
)

type Config struct {
	DatabaseConfig struct {
		Host     string
		Port     string
		Name     string
		User     string
		Password string
	}
	WebApplicationConfig struct {
		Port string
	}
	TinkoffInvestConfig struct {
		Mod             string
		ProductionToken string
		SandboxToken    string
		GrpcUrl         string
		GrpcPort        string
	}
	ApplicationConfig struct {
		PeriodMonthForGetHistoricalData int
	}
}

func NewConfig() (*Config, error) {
	config := new(Config)
	err := config.loadConfiguration()
	if err != nil {
		return nil, err
	}

	return config, nil

}

func (c *Config) loadConfiguration() error {
	err := c.loadEnvFileIfExist()
	if err != nil {
		return err
	}

	c.loadDatabaseConfiguration()
	c.loadWebApplicationConfiguration()
	c.loadTinkoffInvestConfiguration()
	c.loadApplicationConfiguration()

	return nil
}

func (c *Config) loadEnvFileIfExist() error {
	_, err := os.Stat(envFile)
	if err == nil {
		err := godotenv.Load(envFile)
		if err == nil {
			return err
		}
	}

	return nil
}

func (c *Config) loadDatabaseConfiguration() {
	c.DatabaseConfig.Host = os.Getenv("DATABASE_HOST")
	c.DatabaseConfig.Port = os.Getenv("DATABASE_PORT")
	c.DatabaseConfig.Name = os.Getenv("DATABASE_NAME")
	c.DatabaseConfig.User = os.Getenv("DATABASE_USER")
	c.DatabaseConfig.Password = os.Getenv("DATABASE_PASSWORD")
}

func (c *Config) loadWebApplicationConfiguration() {
	c.WebApplicationConfig.Port = os.Getenv("WEB_APPLICATION_PORT")
}

func (c *Config) loadTinkoffInvestConfiguration() {
	c.TinkoffInvestConfig.Mod = os.Getenv("TRADING_MOD")
	c.TinkoffInvestConfig.ProductionToken = os.Getenv("PRODUCTION_TOKEN")
	c.TinkoffInvestConfig.SandboxToken = os.Getenv("SANDBOX_TOKEN")
	c.TinkoffInvestConfig.GrpcUrl = os.Getenv("TINKOFF_INVEST_URL")
	c.TinkoffInvestConfig.GrpcPort = os.Getenv("TINKOFF_INVEST_PORT")
}

func (c *Config) loadApplicationConfiguration() {
	period, err := strconv.Atoi(os.Getenv("APP_PERIOD_MONTH_FOR_GET_HISTORICAL_DATA"))
	if err != nil {
		period = 6
	}

	c.ApplicationConfig.PeriodMonthForGetHistoricalData = period
}
