package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open(viper.GetString("DB_DATABASE"), viper.GetString("DB_USERNAME")+":@tcp("+viper.GetString("DB_HOST")+":"+viper.GetString("DB_PORT")+")/"+viper.GetString("DB_NAME")+"?parseTime=True&loc=Asia%2FJakarta&charset=utf8&autocommit=false")

	if err != nil {
		return nil, err
	}

	// db.SetMaxOpenConns(5) sets the maximum number of open connections to the database.
	// This helps to manage the number of connections and prevent resource exhaustion.
	// A value of 5 is chosen as a reasonable starting point, but it can be adjusted based on the specific requirements and resources available.
	db.SetMaxOpenConns(5)
	// db.SetMaxIdleConns(1) sets the maximum number of idle connections in the connection pool.
	// Idle connections are kept open to be reused, reducing the overhead of establishing new connections.
	// A value of 1 is chosen as a minimum to maintain at least one idle connection, but it can be adjusted based on the application's needs.
	db.SetMaxIdleConns(1)
	// db.SetConnMaxIdleTime(60_000) sets the maximum amount of time (in milliseconds) that an idle connection can remain open.
	// Idle connections that exceed this time will be closed to free up resources.
	// A value of 60,000 milliseconds (1 minute) is chosen as a reasonable default, but it can be adjusted based on the specific requirements.
	db.SetConnMaxIdleTime(60_000)
	// db.SetConnMaxLifetime(10 * 60_000) sets the maximum lifetime of a connection (in milliseconds).
	// Connections that exceed this time will be closed to prevent resource leaks.
	// A value of 10 minutes is chosen as a reasonable default, but it can be adjusted based on the specific requirements.
	db.SetConnMaxLifetime(10 * 60_000)

	return db, nil
}
