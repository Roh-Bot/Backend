package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func PostgresConnectionPool() *pgxpool.Pool {
	const ConnectString = `host=localhost port=5432 user=postgres password=admin dbname=User sslmode=disable`
	pool, err := pgxpool.New(context.Background(), ConnectString)
	if err != nil {
		fmt.Println("Connection Failed")
	}

	errPing := pool.Ping(context.Background())
	if errPing != nil {
		log.Fatal(fmt.Println("Connection failed to Databse"))
	} else {
		fmt.Println("DB Connected")
	}
	return pool
}
