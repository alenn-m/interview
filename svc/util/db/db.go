package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/alenn-m/interview/svc/util"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDatabase() *sqlx.DB {
	port, err := strconv.Atoi(util.GetEnv("DB_PORT", os.Getenv("DB_PORT")))
	if err != nil {
		panic(err)
	}

	log.Println("db user", os.Getenv("DB_USER"))
	log.Println("db password", os.Getenv("DB_PASSWORD"))
	log.Println("db schema", os.Getenv("DB_SCHEMA"))
	log.Println("db host", os.Getenv("DB_HOST"))

	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		util.GetEnv("DB_USER", os.Getenv("DB_USER")),
		util.GetEnv("DB_PASSWORD", os.Getenv("DB_PASSWORD")),
		util.GetEnv("DB_SCHEMA", os.Getenv("DB_SCHEMA")),
		util.GetEnv("DB_HOST", os.Getenv("DB_HOST")),
		port,
	),
	)
	if err != nil {
		panic(err)
	}

	db.MapperFunc(func(s string) string {
		return s
	})

	return db
}
