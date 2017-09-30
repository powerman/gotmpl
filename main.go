package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

var (
	app = strings.TrimSuffix(path.Base(os.Args[0]), ".test")
	ver = "v?"
	cfg struct {
		version bool
	}
)

func init() {
	flag.BoolVar(&cfg.version, "version", false, "print version")
}

func main() {
	flag.Parse()

	if cfg.version {
		fmt.Println(app, ver, runtime.Version())
		os.Exit(0)
	}
}
