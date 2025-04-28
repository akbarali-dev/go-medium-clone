package main

import (
	"fmt"
	"log"
	"mediumclone/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load(".")
	psqlUrl := fmt.Sprintf(
		"host=%s, port=%s, user=%s, password=%s, dbname=%s, sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	psqlConn, err := sqlx.Open("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("Failed to conect to postgresql: %v", err)
	}
	log.Println("Postgres Connection successfull 👌")
	_ = psqlConn
}
