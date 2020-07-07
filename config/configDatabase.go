package config

import (
	"database/sql"
	"fmt"
	"hotelAPI/utils"
	"log"

	//for connection mysql db
	_ "github.com/go-sql-driver/mysql"
)

// var (
// 	dbUser,
// 	dbPassword,
// 	dbHost,
// 	dbPort,
// 	dbSchema string
// )

//EnvConn is function to get Environment Variabel for DB connection
func EnvConn() *sql.DB {
	dbEngine := utils.ViperGetEnv("DB_ENGINE", "mysql") //mysql
	dbUser := utils.ViperGetEnv("DB_USER", "root")      //root
	dbPassword := utils.ViperGetEnv("DB_PASSWORD", "password")
	dbHost := utils.ViperGetEnv("DB_HOST", "localhost") //localhost
	dbPort := utils.ViperGetEnv("DB_PORT", "3306")      //3306
	dbSchema := utils.ViperGetEnv("DB_SCHEMA", "schema")

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbSchema)
	db, err := sql.Open(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
