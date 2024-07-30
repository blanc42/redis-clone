package main

import (
	"log"

	"redis-clone/internal/server"
)

func main() {
	s := server.NewServer(":6379")
	log.Fatal(s.Run())
}
