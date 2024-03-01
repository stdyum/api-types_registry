package main

import (
	"log"

	"github.com/stdyum/api-types-registry/internal"
)

func main() {
	log.Fatal("error launching web server", internal.App().Run())
}
