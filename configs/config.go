package configs

import (
	"database/sql"
	"log"

	"github.com/muhfaris/lib-go/psql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config is config app
type Config struct {
	Based
	HTTP HTTP
	Persistance
	Connection *Connection
}

// Based is based app config
type Based struct {
	Name string
	Port int
}

// HTTP is http config
type HTTP struct {
	ReadTimeout int
}

// Persistance is wrap all type store
type Persistance struct {
	Database PersistanceInSQL
}

// PersistanceInSQL is database
type PersistanceInSQL struct {
	Name     string
	Host     string
	Port     int
	Username string
	Password string
	SSLMode  string
}

type Connection struct {
	DB     *sql.DB
	Logger *logrus.Logger
}

// NewConfig is initialize app config
func NewConfig() *Config {
	log.Println(viper.GetString("persistence.database.name"))
	return &Config{
		Based: Based{
			Name: viper.GetString("app.name"),
			Port: viper.GetInt("app.port"),
		},
		HTTP: HTTP{
			ReadTimeout: viper.GetInt("app.http.read_timeout"),
		},
		Persistance: Persistance{
			Database: PersistanceInSQL{
				Name:     viper.GetString("persistence.database.name"),
				Host:     viper.GetString("persistence.database.host"),
				Port:     viper.GetInt("persistence.database.port"),
				Username: viper.GetString("persistence.database.username"),
				Password: viper.GetString("persistence.database.password"),
				SSLMode:  viper.GetString("persistence.database.ssl_mode"),
			},
		},
		Connection: &Connection{
			Logger: logrus.New(),
		},
	}
}

func (config *Config) InitializeConnectionPSQL() {
	dbOptions := psql.DBOptions{
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		Username: config.Database.Username,
		Password: config.Database.Password,
		DBName:   config.Database.Name,
		SSLMode:  config.Database.SSLMode,
	}

	conn, err := psql.Connect(&dbOptions)
	if err != nil {
		log.Fatalln("Database:", err)
	}

	log.Println("Database connected ...")
	config.Connection.DB = conn
}
