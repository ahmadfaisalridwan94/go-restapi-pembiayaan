package gorm

import (
	"fmt"
	"os"
	"time"

	gormMySql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type MySql struct{
	host string
	port string
	user string
	password string
	database string
	logMode int
	maxIdleConnection int
	maxOpenConnection int
	connectionMaxLifetimeInSeconds int
}

type mysqlOption func (*MySql)

func Connect() (*gorm.DB, error) {
	logMode := 0
	if os.Getenv("DB_DEBUG") == "TRUE" {
		logMode = 3
	}

	dbOptions := &MySql{
		host:                          os.Getenv("DB_HOST"),
		port:                          os.Getenv("DB_PORT"),
		user:                          os.Getenv("DB_USER"),
		password:                      os.Getenv("DB_PASSWORD"),
		database:                      os.Getenv("DB_DATABASE"),
		logMode:                       logMode,
		maxIdleConnection:             5,
		maxOpenConnection:             10,
		connectionMaxLifetimeInSeconds: 60,
	}

	return connect(dbOptions)
}

func connect(param *MySql) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		param.user, param.password, param.host, param.port, param.database)

	db, err := gorm.Open(gormMySql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(param.logMode)),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	// set configuration pooling connection
	mysqlDB, _ := db.DB()
	mysqlDB.SetMaxOpenConns(param.maxOpenConnection)
	mysqlDB.SetConnMaxLifetime(time.Duration(param.connectionMaxLifetimeInSeconds) * time.Minute)
	mysqlDB.SetMaxIdleConns(param.maxIdleConnection)

	migrateAllTables(db)

	return db, nil
}

func SetMaxIdleConns(conns int) mysqlOption {
	return func(c *MySql) {
		if conns > 0 {
			c.maxIdleConnection = conns
		}
	}
}

func SetMaxOpenConns(conns int) mysqlOption {
	return func(c *MySql) {
		if conns > 0 {
			c.maxOpenConnection = conns
		}
	}
}

func SetConnMaxLifetime(conns int) mysqlOption {
	return func(c *MySql) {
		if conns > 0 {
			c.connectionMaxLifetimeInSeconds = conns
		}
	}
}

func migrateAllTables(db *gorm.DB) {
	// models.MigrateAccountBank(db)
}
