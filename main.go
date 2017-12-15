package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Luzifer/rconfig"
)

var (
	cfg = struct {
		InFile         string `flag:"in,i" default:"-" description:"File to encode, by default stdin"`
		VersionAndExit bool   `flag:"version" default:"false" description:"Prints current version and exits"`
	}{}

	version = "dev"
)

func init() {
	if err := rconfig.Parse(&cfg); err != nil {
		log.Fatalf("Unable to parse commandline options: %s", err)
	}

	if cfg.VersionAndExit {
		fmt.Printf("str2hex %s\n", version)
		os.Exit(0)
	}
}

func main() {
	var (
		err    error
		inFile io.Reader
	)

	switch cfg.InFile {
	case "-":
		inFile = os.Stdin
	default:
		f, err := os.Open(cfg.InFile)
		if err != nil {
			log.Fatalf("Unable to open file: %s", err)
		}

		inFile = f
		defer f.Close()
	}

	buf := new(bytes.Buffer)
	if _, err = io.Copy(buf, inFile); err != nil {
		log.Fatalf("Unable to read from input: %s", err)
	}

	fmt.Printf("%x", buf.String())
}
