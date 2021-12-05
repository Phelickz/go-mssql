package database

import (
	// "database/sql"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

func ConnectDb() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panic(err)
	}

	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbServer := os.Getenv("DATABASE_SERVER")
	dbName := os.Getenv("DATABASE_NAME")
	dbPort := os.Getenv("DATABASE_PORT")

	//converting the dbPort to integer
	portInt, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatal(err)
	}

	query := url.Values{}
	query.Add("database", dbName)
	query.Add("connection timeout", "60")

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(dbUsername, dbPassword),
		Host:   fmt.Sprintf("%s:%d", dbServer, portInt),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	fmt.Println(u.String())
	db, err := sql.Open("sqlserver", u.String())

	if err != nil {
		log.Panic(err)
	} else {
		log.Println("Database connection successful")
	}

	// db.Close()

	return db
}

var Client *sql.DB = ConnectDb()

func OpenDb() *sql.DB {
	return Client
}
