package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/sataapon/btc/internal/sql"
)

type config struct {
	hostname    string
	servicePort string
}

var (
	cfg     *config
	onceCfg sync.Once
)

func GetHostName() string {
	onceCfg.Do(getConfig)
	return cfg.hostname
}

func GetServicePort() string {
	onceCfg.Do(getConfig)
	return cfg.servicePort
}

func getConfig() {
	cfg = &config{
		hostname:    envString("Hostname", "localhost"),
		servicePort: envString("ServicePort", "8080"),
	}
}

var (
	db     sql.DB
	onceDB sync.Once
)

func GetDB() sql.DB {
	onceDB.Do(getDB)
	return db
}

func getDB() {
	driverName := envString("DriverName", "postgres")

	postgreSqLServer := envString("PostgreSqLServer", "localhost")
	postgreSqlPort := envString("PostgreSqlPort", "5432")
	postgreSqlUser := envString("PostgreSqlUser", "postgres")
	postgreSqlPassword := envString("PostgreSqlPassword", "123456789")
	postgreSqlDatabase := envString("PostgreSqlDatabase", "wallet")

	var err error
	db, err = sql.Open(
		driverName,
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
			postgreSqLServer, postgreSqlPort, postgreSqlUser, postgreSqlPassword, postgreSqlDatabase),
	)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func envString(name, defaultValue string) string {
	env := os.Getenv(name)
	if env != "" {
		return env
	}
	return defaultValue
}
