package main

import (
	"context"
	"log"
	"os"
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
    abspath := filepath.Join(cfg.Ingest.BaseDir, relpath)
    logger.INFO.Printf("[Ingest start] %s\n", abspath)

    stat, e := os.Stat(abspath)
    if e != nil {
        logger.ERROR.Printf("[Ingest end][ERROR] Error trying to stat file %s\n", abspath)
        return
    }

    ctx := context.Background()
    conn := db.Conn

    // Get or insert parent collection
    dir := filepath.Dir(abspath)
    pcid := conn.GetCollectionId(ctx, dir)
    var cid int
    if pcid == nil {
        cid = conn.InsertCollection(ctx, dir, filepath.Base(dir))
    } else {
        cid = *pcid
    }

    // Thumbnails
    stripped := filepath.Join(cfg.Ingest.ThumbDir, relpath)
    if !stat.IsDir() {
        stripped = files.RemoveExt(stripped)
    }
    ppng, _ := files.VerifyAndGetBasename(stripped + ".png")
    pgif, _ := files.VerifyAndGetBasename(stripped + ".gif")

    // content files (relative paths)
    var pfiles *[]string
    if stat.IsDir() {
        files, err := files.GetFilesInDir(abspath)
        if err != nil {
            logger.ERROR.Printf("[Ingest end][ERROR] Error when getting files in directory %s\n", abspath)
            return
        }
        pfiles = &files
    }

    // Insert or update entry
    if !conn.EntryExists(ctx, relpath) {
        conn.InsertEntry(ctx, relpath, filepath.Base(stripped), ppng, pgif, cid, pfiles)
    } else {
        conn.UpdateEntry(ctx, relpath, ppng, pgif, pfiles)
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

