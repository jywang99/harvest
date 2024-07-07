package main

import (
	"flag"

	"jy.org/harvest/src/config"
)

type args struct {
    config config.ConfigOverride
}

func parseArgs() args {
    var a args
    flag.StringVar(&a.config.YmlPath, "f", "", "Path to configuration file")
    flag.StringVar(&a.config.BaseDir, "i", "", "Path containing original images/videos")
    flag.StringVar(&a.config.ThumbDir, "t", "", "Path containing thumbnails")
    flag.StringVar(&a.config.IndexFile, "x", "", "Path to index file")
    flag.Parse()
    return a
}

