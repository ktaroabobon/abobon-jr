package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Hello, world!")

	log.Println(os.Getenv("MYSQL_DATABASE"))
}
