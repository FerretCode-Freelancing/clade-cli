package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/ferretcode-freelancing/clade-cli/connect"
	"github.com/ferretcode-freelancing/clade-cli/container"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	appDir := filepath.FromSlash(homeDir + "/.clade")

	if _, err := os.Stat(appDir); err != nil {
		if !os.IsNotExist(err) {
			log.Fatal(err)
		}

		err := os.Mkdir(appDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = connect.Preload()
	if err != nil {
		log.Fatal(err)
	}

	insecure := flag.Bool("insecure", false, "whether to enforce https")
	flag.Parse()

	switch flag.Arg(0) {
	case "connect":
		err := connect.Connect(flag.Args()[1:], *insecure)
		if err != nil {
			log.Fatal(err)
		}
	case "container":
		switch flag.Arg(1) {
		case "create":
			err := container.Create(flag.Args()[2:], *insecure)
			if err != nil {
				log.Fatal(err)
			}
		case "delete":
		case "update":
		default:
			help()
		}
	case "registry":
	default:
		help()
	}
}

func help() {
	log.Fatalf("\nconnect\n\tconnect: connect to the clade server\ncontainer\n\tcreate: create a new container\n\tdelete: delete a container\n\tupdate: update a container\nregistry\n\tadd: configure a registry\n\tremove: remove an existing registry")
}
