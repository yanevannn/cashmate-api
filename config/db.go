package config

import (
	"context"
	"log"
	"os"
	"fmt"

	"github.com/jackc/pgx/v5"

)

func ConnectDB() (*pgx.Conn, error){
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	databaseURL := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		user,
		password,
		host,
		port,
		dbname,
	)

	log.Println("Connecting to database Successfully...")

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
        return nil, err
    }

	 return conn, nil

}