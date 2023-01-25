package database

import (
	db "WGameBack/database/db"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

//var DB *pgxpool.Pool = Connection()

//func Connection() *pgxpool.Pool {
//	poolConfig, err := newPoolConfig()
//	if err != nil {
//		log.Fatal(e.Wrap("pool config error", err))
//	}
//
//	poolConfig.MaxConns = 5
//
//	conn, err := newConnection(poolConfig)
//	if err != nil {
//		log.Fatal(e.Wrap("connection to database failed", err))
//	}
//
//	_, err = conn.Exec(context.Background(), ";")
//	if err != nil {
//		log.Fatal(e.Wrap("ping failed", err))
//	}
//	log.Printf("Ping OK!\n")
//
//	return conn
//}
//
//func newPoolConfig() (*pgxpool.Config, error) {
//	connStr := os.Getenv("DB_CONNECT_STRING")
//
//	poolConfig, err := pgxpool.ParseConfig(connStr)
//	if err != nil {
//		return nil, err
//	}
//
//	return poolConfig, nil
//}
//
//func newConnection(poolConfig *pgxpool.Config) (*pgxpool.Pool, error) {
//	conn, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
//	if err != nil {
//		return nil, err
//	}
//
//	return conn, nil
//}

func Connection() *db.Queries {
	connStr := os.Getenv("DB_CONNECT_STRING")
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	queries := db.New(DB)
	return queries
}
