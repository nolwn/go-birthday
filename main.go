package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nolwn/go-birthday/data"
)

func main() {
	db, err := data.New()
	if err != nil {
		log.Fatal("Could not open a new databse connection: ", err)
	}
	defer closeDatabase(db)

	args := os.Args[1:]

	fmt.Println(parseArguments(db, args))
}

func closeDatabase(db data.Database) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
