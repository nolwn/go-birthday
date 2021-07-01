package main

import (
	"fmt"
	"strings"

	"github.com/nolwn/go-birthday/data"
)

// format response beginning
const msgNoBirthdays = "No birthdays"
const msgBirthdays = "It's %s birthday"

// format response ending
const msgToday = "today"
const msgOnMonthDay = "on %s %s"
const msgInMonth = "in %s"

func formatBirthdayResult(contacts []data.Contact, month uint8, day uint8, punctuation string) string {
	var messageStart string
	var messageEnd string

	args := make([]interface{}, 0, 3)
	contactsString := formatNames(contacts)
	monthString := nameMonth(month)
	dayString := nameDay(day)

	if len(contacts) == 0 {
		messageStart = msgNoBirthdays
	} else {
		messageStart = msgBirthdays
		args = append(args, contactsString)
	}

	// To get birthdays for a specific time, a month is required. So if there is no month, assume
	// today.
	if month == 0 {
		messageEnd = msgToday
	} else if day == 0 {
		messageEnd = msgInMonth
		args = append(args, monthString)
	} else {
		messageEnd = msgOnMonthDay
		args = append(args, monthString, dayString)
	}

	formatString := fmt.Sprintf("%s %s%s", messageStart, messageEnd, punctuation)

	return fmt.Sprintf(formatString, args...)
}

func nameMonth(month uint8) string {
	if month == 0 {
		return ""
	}

	monthStruct := months[month-1]
	return strings.Title(monthStruct.fullName)
}

func nameDay(day uint8) string {
	var suffix string
	lastDigit := day % 10

	if day == 0 {
		return ""
	}

	switch lastDigit {
	case 1:
		suffix = "st"
	case 2:
		suffix = "nd"
	case 3:
		suffix = "rd"
	default:
		suffix = "th"
	}

	return fmt.Sprintf("%d%s", day, suffix)
}

func formatNames(contacts []data.Contact) string {
	var contactsString string

	var comma string
	for _, contact := range contacts {
		contactsString += comma + contact.Name
		comma = ", "
	}

	contactsString += "'s"

	return contactsString
}
