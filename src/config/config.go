package config

import (
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type ingestConfig struct {
    BaseDir   string `yaml:"baseDir"`
    ThumbDir  string `yaml:"thumbDir"`
    IndexFile string `yaml:"indexFile"`
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

var basePath = "conf/" // TODO no hardcoding
var configPath = path.Join(basePath, "config.yml")

func readYmlConfig(cfg *config) {
    f, err := os.Open(configPath)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    s, _ := f.Stat()
    if s.Size() == 0 {
        return
    }

    decoder := yaml.NewDecoder(f)
    err = decoder.Decode(&cfg)
    if err != nil {
        log.Fatal(err)
    }
}

func initConfig() *config {
    var cfg config
    readYmlConfig(&cfg)
    return &cfg
}
var Config = initConfig()

