package main

import (
	"os"

	"github.com/nolwn/go-birthday/data"
)

type option struct {
	property    string
	shortName   string
	longName    string
	variables   []string
	description string
}

var options = []option{
	{
		"month",
		"-m",
		"--month",
		[]string{"<month>"},
		"prints contacts whose birthdays are on <date>",
	},
}

func parseArguments(db data.Database, args []string) (response string) {
	args = os.Args[1:]

	if len(args) == 0 {
		response = birthdaysToday(db)
	}

	// var isVariable bool
	// var isRequired bool

	for _, arg := range args {
		for _, opt := range options {
			if arg == opt.shortName || arg == opt.longName {

			}
		}
	}

	return
}
