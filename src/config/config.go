package config

import (
	"errors"
	"os"
	"path/filepath"
)

type ingestConfig struct {
    BaseDir   string `yaml:"baseDir"`
    ThumbDir  string `yaml:"thumbDir"`
    IndexFile string `yaml:"indexFile"`
    ExtStr string `yaml:"exts"`
    ExtMap map[string]bool
    IgnoreStr string `yaml:"ignore"`
    IgnoreMap map[string]bool
    DotFiles bool `yaml:"dotfiles"`
}

type dbConfig struct {
    Host     string `yaml:"host"`
    User     string `yaml:"user"`
    Password string `yaml:"password"`
    Port     int    `yaml:"port"`
    Database string `yaml:"database"`
    Schema   string `yaml:"schema"`
    SSLMode  string `yaml:"sslmode"`
}

type logConfig struct {
    Path string `yaml:"path"`
}

type config struct {
    Ingest ingestConfig `yaml:"ingest"`
    DB dbConfig `yaml:"db"`
    Log logConfig `yaml:"log"`
}

var Config = &config{
    Ingest: ingestConfig{},
    DB: dbConfig{},
    Log: logConfig{
        Path: "/tmp/harvest.log",
    },
}

func Validate() error {
    ingest := Config.Ingest
    if !dirExists(ingest.BaseDir) {
        return errors.New("Input directory does not exist")
    }
    if !dirExists(ingest.ThumbDir) {
        return errors.New("Thumbnail directory does not exist")
    }
    if !fileExists(ingest.IndexFile) {
        return errors.New("Index file does not exist")
    }
    if len(ingest.ExtMap) == 0 {
        return errors.New("No file extensions specified")
    }

    db := Config.DB
    if db.Host == "" || db.User == "" || db.Password == "" || db.Database == "" || db.Schema == "" || db.SSLMode == "" || db.Port == 0 {
        return errors.New("Invalid database configuration")
    }

    logs := Config.Log
    if logs.Path != "" && !parentDirExists(logs.Path) {
        return errors.New("Invalid log file path")
    }

    return nil
}

func fileExists(path string) bool {
    _, err := os.Stat(path)
    return err == nil
}

func parentDirExists(path string) bool {
    parent := filepath.Dir(path)
    return dirExists(parent)
}

func dirExists(dir string) bool {
    if dir == "" {
        return false
    }
    _, err := os.Stat(dir)
    return err == nil
}

