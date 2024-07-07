package main

import (
	"context"
	"log"
	"path/filepath"

	"jy.org/harvest/src/config"
	"jy.org/harvest/src/db"
	"jy.org/harvest/src/files"
	"jy.org/harvest/src/logging"
)

var cfg = config.Config
var logger = logging.Logger

func main() {
    // read config
    args := parseArgs()
    config.Override(args.config)
    err := config.Validate()
    if err != nil {
        log.Fatal(err)
        return
    }

    // init loggers
    logging.InitLogFiles()
    logger.INFO.Println("Starting harvest")
    defer logger.INFO.Println("Exiting harvest")
    logger.INFO.Printf("Config: %+v\n", cfg)

    db.Setup()
    dbconn := db.Conn
    defer dbconn.Close()

    reader, err := files.NewFileReader(cfg.Ingest.IndexFile)
    if err != nil {
        logger.ERROR.Printf("Error when creating index file reader")
        return
    }
    defer reader.Close()

    for {
        line, ok := reader.ReadNextLine()
        if !ok {
            break
        }
        ingest(line)
    }

    relateCollections()
}

func ingest(relpath string) {
    file := filepath.Join(cfg.Ingest.BaseDir, relpath)
    logger.INFO.Printf("[Ingest start] %s\n", file)

    if !files.VerifyFileExists(file) {
        logger.ERROR.Printf("[Ingest end][ERROR] File does not exist: %s\n", file)
        return
    }

    ctx := context.Background()
    conn := db.Conn

    // Get or insert collection
    dir := filepath.Dir(file)
    pcid := conn.GetCollectionId(ctx, dir)
    var cid int
    if pcid == nil {
        cid = conn.InsertCollection(ctx, dir, filepath.Base(dir))
    } else {
        cid = *pcid
    }

    // Insert or update entry
    stripped := filepath.Join(cfg.Ingest.ThumbDir, files.RemoveExt(relpath))
    png := stripped + ".png"
    gif := stripped + ".gif"
    var ppng, pgif *string
    if files.VerifyFileExists(png) {
        ppng = &png
    }
    if files.VerifyFileExists(gif) {
        pgif = &gif
    }
    if !conn.EntryExists(ctx, file) {
        conn.InsertEntry(ctx, file, filepath.Base(stripped), ppng, pgif, cid)
    } else {
        conn.UpdateEntry(ctx, file, ppng, pgif)
    }

    logger.INFO.Println("[Ingest end][ok]")
}

func relateCollections() {
    logger.INFO.Println("[Relate start]")

    ctx := context.Background()
    conn := db.Conn

    pathIds := conn.GetPathIds(ctx)
    for path, id := range pathIds {
        parent := filepath.Dir(path)
        pid, ok := pathIds[parent]
        if !ok {
            continue
        }
        logger.INFO.Printf("%s -> %s\n", path, parent)
        conn.SetParentCollection(ctx, id, pid)
    }

    logger.INFO.Println("[Relate end][ok]")
}

