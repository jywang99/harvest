package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func panicIfError(err error) {
    if err != nil && err != pgx.ErrNoRows {
        panic(err)
    }
}

func (db *DbConn) GetCollectionId(ctx context.Context, path string) *int {
    var id int
    err := db.conn.QueryRow(ctx, "SELECT id FROM collection WHERE path = $1", path).Scan(&id)
    if err == pgx.ErrNoRows {
        return nil
    }
    panicIfError(err)
    return &id
}

func (db *DbConn) InsertCollection(ctx context.Context, path, name string) int {
    var id int
    err := db.conn.QueryRow(ctx, "INSERT INTO collection (path, disp_name) VALUES ($1, $2) RETURNING id", path, name).Scan(&id)
    if err != nil {
        panic(err)
    }
    return id
}

func (db *DbConn) EntryExists(ctx context.Context, path string) bool {
    row := db.conn.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM entry WHERE path = $1)", path)
    var exists bool
    err := row.Scan(&exists)
    panicIfError(err)
    return exists
}

func (db *DbConn) InsertEntry(ctx context.Context, path, name string, png, gif *string, parentId int) int {
    var dthumbs []*string
    if gif != nil {
        dthumbs = []*string{gif}
    }

    var id int
    err := db.conn.QueryRow(ctx, 
        "INSERT INTO entry (path, disp_name, thumb_static, thumb_dynamic, parent) VALUES ($1, $2, $3, $4, $5) RETURNING id", 
        path, name, png, dthumbs, parentId,
    ).Scan(&id)
    if err != nil {
        panic(err)
    }
    return id
}

func (db *DbConn) UpdateEntry(ctx context.Context, path string, png, gif *string) {
    var dthumbs []*string
    if gif != nil {
        dthumbs = []*string{gif}
    }
    _, err := db.conn.Exec(ctx, "UPDATE entry SET thumb_static = $2, thumb_dynamic = $3 WHERE path = $1", path, png, dthumbs)
    if err != nil {
        panic(err)
    }
}

func (db *DbConn) GetPathIds(ctx context.Context) map[string]int {
    rows, err := db.conn.Query(ctx, "SELECT id, path FROM collection ORDER BY path")
    panicIfError(err)
    defer rows.Close()

    ids := make(map[string]int)
    for rows.Next() {
        var path string
        var id int
        err := rows.Scan(&id, &path)
        if err != nil {
            panic(err)
        }
        ids[path] = id
    }

    return ids
}

func (db *DbConn) SetParentCollection(ctx context.Context, child, parent int) {
    _, err := db.conn.Exec(ctx, "UPDATE collection SET parent = $2 WHERE id = $1", child, parent)
    if err != nil {
        panic(err)
    }
}
