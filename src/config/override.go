package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type ConfigOverride struct {
    YmlPath string
    BaseDir string
    ThumbDir string
    IndexFile string
}

func loadYml(configPath string) {
    f, err := os.Open(configPath)
    if err != nil {
        log.Fatal(fmt.Errorf("Error loading config yml file: %v", err))
    }
    defer f.Close()

    s, _ := f.Stat()
    if s.Size() == 0 {
        return
    }

    decoder := yaml.NewDecoder(f)
    err = decoder.Decode(&Config)
    if err != nil {
        log.Fatal(err)
    }

    Config.Ingest.IgnoreMap = stringToMap(Config.Ingest.IgnoreStr)
    Config.Ingest.ExtMap = stringToMap(Config.Ingest.ExtStr)

}

func Override(override ConfigOverride) {
    loadYml(override.YmlPath)
    doIfNotEmpty(override.BaseDir, func(s string) { Config.Ingest.BaseDir = s })
    doIfNotEmpty(override.ThumbDir, func(s string) { Config.Ingest.ThumbDir = s })
}

func doIfNotEmpty(s string, f func(string)) {
    if s != "" {
        f(s)
    }
}

func stringToMap(s string) map[string]bool {
    m := make(map[string]bool)
    if s == "" {
        return m
    }

    for _, v := range strings.Split(s, ":") {
        m[v] = true
    }
    return m
}


