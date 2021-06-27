package main

import (
	"fmt"
	"log"

	"github.com/nolwn/go-birthday/data"
)

func main() {
	db, err := data.New()
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal("Could not create a connection to the databse: ", err)
	}

	err = db.InitializeDatabase()
	if err != nil {
		log.Fatal("Could not initialize the database: ", err)
	}

	if err != nil {
		log.Fatal("Could not add contact: ", err)
	}

	id, err := db.Add("Edu", data.June, 20, 1989)
	if err != nil {
		log.Fatal("Could not add entry: ", err)
	}

	fmt.Printf("Added Edu with an id of %d", id)

	contact, err := db.GetByName("Edu")
	if err != nil {
		log.Fatal("Could not retrieve Edu! ", err)
	}

	if contact.IsNull() {
		fmt.Print("There is not contact by that name")
	}

	fmt.Printf("Recovered a record: %v", contact)
}
