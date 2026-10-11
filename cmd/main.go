package main

import (
	"log"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	h := api.mount()
	e := api.run(h)

	if e != nil {
		log.Printf("Server has failed to start, err : %s", e)
		os.Exit(1)
	}
}
