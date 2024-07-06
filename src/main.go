package main

import (
	"jy.org/harvest/src/config"
	"jy.org/harvest/src/db"
	"jy.org/harvest/src/logging"
)

var cfg = config.Config
var logger = logging.Logger

func main() {
    logger.INFO.Println("Starting harvester")
    defer logger.INFO.Println("Exiting harvester")

    dbconn := db.Conn
    defer dbconn.Close()

    logger.INFO.Printf("Config: %+v\n", cfg)
}

