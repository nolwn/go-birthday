package main

import (
	"log"
	"time"

	"github.com/nolwn/go-birthday/data"
)

func birthdaysToday(db data.Database) string {
	var punctuation string

	today := time.Now()
	month := uint8(today.Month())
	day := uint8(today.Day())

	contacts, err := db.GetByDate(month, day)
	if err != nil {
		log.Fatal(err)
	}

	if len(contacts) > 0 {
		punctuation = "!"
	} else {
		punctuation = "."
	}

	return formatBirthdayResult(contacts, 0, 0, punctuation)
}
