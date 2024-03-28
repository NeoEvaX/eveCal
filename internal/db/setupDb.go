package db

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DB *Queries
)

func Setup(logger *slog.Logger) {

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Error("Unable to create connection pool:", err)
		os.Exit(1)
	}

	DB = New(dbpool)

	// var err error
	// DB, err = sql.Open("sqlite3", "./test.db")
	// if err != nil {
	// 	panic("Could not connect to db")
	// }
	// err = DB.Ping()
	// if err != nil {
	// 	fmt.Println("Could not ping db", err)
	// }
}
