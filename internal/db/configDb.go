package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitDB(connStr string) {
    var err error
    Pool, err = pgxpool.New(context.Background(), connStr)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v", err)
    }
    fmt.Println("Database connection established")
}
