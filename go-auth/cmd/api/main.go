package main

import (
	"database/sql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/karankumarshreds/GoMicroservices/go-auth/data"
	"log"
	"net/http"
	"os"
	"time"
	//_ "github.com/jackc/"
)

const PORT = ":8001"

type App struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Printf("Starting auth service on port %v", PORT)

	pgClient := ConnectToDatabase()
	if pgClient != nil {
		log.Panic("Cannot connect to postgres")
	}

	app := App{
		DB:     pgClient,
		Models: data.New(pgClient),
	}

	server := http.Server{
		Addr:    PORT,
		Handler: app.routes(),
	}
	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

// InitPostgres initializes the connection to the postgres database
func InitPostgres(dataSourceName string) (*sql.DB, error) {
	driverName := "pgx"
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Panic(err)
		return nil, err
	} else {
		if err := db.Ping(); err != nil {
			log.Panic(err)
			return nil, err
		}
	}
	return db, nil
}

// ConnectToDatabase tries to connect to the database for 20 seconds max
func ConnectToDatabase() *sql.DB {
	connString := os.Getenv("DSN")
	if connString == "" {
		log.Panic("DATABASE CONNECTION STRING NOT FOUND")
	}
	count := 0
	// keep connecting to the db
	for {
		if count > 10 {
			log.Panic("Cannot connect to the database")
			return nil
		}
		connection, err := InitPostgres(connString)
		if err != nil {
			log.Println("Postgres not yet ready, trying again...")
			count++
		} else {
			log.Println("Connected to the Postgres database successfully")
			return connection
		}
		log.Println("Waiting for two seconds to try again")
		time.Sleep(time.Second * 2)
		continue
	}
}
