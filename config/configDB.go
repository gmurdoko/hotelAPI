package config

import (
	"database/sql"
	"fmt"
	"hotelAPI/utils"
	"log"
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
	db, err := connDB(dbEngine, dbSource)
	if err != nil {
		log.Panic(err)
	}
	return db
	// return dbEngine, dbSource
}

func connCheck(db *sql.DB) (*sql.DB, error) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	return db, err
}

//connDB Connecting app to DB using variabel from EnvConn
func connDB(DbEngine, DbSource string) (db *sql.DB, err error) {
	db, _ = sql.Open(DbEngine, DbSource)
	db, err = connCheck(db)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
