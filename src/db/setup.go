package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"jy.org/harvest/src/config"
	"jy.org/harvest/src/logging"
)

var cfg = config.Config.DB
var logger = logging.Logger

type DbConn struct {
    conn *pgx.Conn
}

func Setup() *DbConn {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
    url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode)
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
        logger.ERROR.Printf("Error connecting to database: %v\n", url)
        logger.ERROR.Fatalf("Error: %v\n", err)
		os.Exit(1)
	}
    logger.INFO.Printf("Connected to database: %v\n", url)
    return &DbConn{conn: conn}
}

func (db *DbConn) Close() {
    db.conn.Close(context.Background())
}

var Conn = Setup()

