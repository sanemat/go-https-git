package main

import (
	"log"
	"os"

	"github.com/sanemat/go-httpsgit"
)

func main() {
	log.SetFlags(0)
	err := httpsgit.Run(os.Args[1:], os.Stdout, os.Stderr)
	if err != nil {
		log.Println(err)
		exitCode := 1
		os.Exit(exitCode)
	}
}
