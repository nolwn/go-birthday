package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const databaseName = "./birthdays.db"
const driverName = "sqlite3"

const sqlCreateContactsTable = `
	CREATE TABLE IF NOT EXISTS contacts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE,
		month INTEGER,
		day INTEGER,
		year INTEGER
	);
`
const sqlInsertNewContact = "INSERT INTO contacts (name, month, day, year) VALUES (?, ?, ?, ?);"
const sqlGetByContactName = "SELECT * FROM contacts WHERE name = ?;"
const sqlGetByMonth = "SELECT * FROM contacts WHERE month=?;"
const sqlGetByMonthAndDay = "SELECT * FROM contacts WHERE (month=? AND day=?);"

const (
	// Unknown represents a blank value, any part of the date we don't know
	Unknown uint8 = iota

	// January because remembering date numbers is hard
	January

	// February because remembering date numbers is hard
	February

	// March because remembering date numbers is hard
	March

	// April because remembering date numbers is hard
	April

	// May because remembering date numbers is hard
	May

	// June because remembering date numbers is hard
	June

	// July because remembering date numbers is hard
	July

	// August because remembering date numbers is hard
	August

	// September because remembering date numbers is hard
	September

	// October because remembering date numbers is hard
	October

	// November because remembering date numbers is hard
	November

	// December because remembering date numbers is hard
	December
)

// Contact represents a contact record.
type Contact struct {
	ID            int64
	Name          string
	Month         int
	Day           int
	Year          int
	isInitialized bool
}

func (c Contact) IsNull() bool {
	return !c.isInitialized
}

// Database represents a database connection. It includes methods from database/sql but,
// more importantly, it add methods specific to this program.
type Database struct {
	*sql.DB
}

// New returns a new Database object with an open connection to the sqlite file. Returns
// an error if the databse cannot be opened.
func New() (Database, error) {
	db, err := sql.Open(driverName, databaseName)
	if err != nil {
		return Database{}, err
	}

	return Database{db}, nil
}

// InitializeDatabse creates the needed tables in the databse if they don't already
// exist.
func (d *Database) InitializeDatabase() error {
	stmt, err := d.Prepare(sqlCreateContactsTable)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	return nil
}

// Add takes a name which is the unique name for the contact and a dateID which is
// the id for the date
func (d *Database) Add(name string, month int, day int, year int) (id int64, err error) {
	stmt, err := d.Prepare(sqlInsertNewContact)
	if err != nil {
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(name, month, day, year)
	if err != nil {
		return
	}

	id, err = res.LastInsertId()
	if err != nil {
		return
	}

	return
}

// GetByName takes a name and searches the database for that contact.
func (d *Database) GetByName(name string) (contact Contact, err error) {
	stmt, err := d.Prepare(sqlGetByContactName)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(name)
	if err != nil {
		return
	}
	defer rows.Close()

	// There should only be one record with this name, so assume no more than one row.
	hasRow := rows.Next()

	if !hasRow {
		err = rows.Err()

		if err != nil {
			log.Fatal(err)
		}

		return
	}

	err = rows.Scan(&contact.ID, &contact.Name, &contact.Month, &contact.Day, &contact.Year)
	if err != nil {
		return
	}

	contact.isInitialized = true

	return
}

// GetByDate takes some day and month and figures out whose birthday it is on that day.
// Month must not be Unknown (0). If day is set to Unknown, then anyone whose birthday
// falls in that month will be returned.
//
// GetByDate assumes that you are using a correct date int, it does not check to make
// sure that birthMonth or birthDay are out of range. If at all possible, use the
// provided month constants to ensure you are setting correct months.
func (d *Database) GetByDate(birthMonth uint8, birthDay uint8) ([]Contact, error) {
	var query string
	var args []interface{}
	var id int64
	var name string
	var month int
	var day int
	var year int

	contacts := make([]Contact, 0, 2)

	if birthMonth == Unknown {
		return nil, fmt.Errorf("birthMonth must be known")
	}

	if birthDay == Unknown {
		query = sqlGetByMonth
		args = []interface{}{birthMonth}
	} else {
		query = sqlGetByMonthAndDay
		args = []interface{}{birthMonth, birthDay}

	}

	stmt, err := d.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &name, &month, &day, &year)
		if err != nil {
			return nil, err
		}

		contacts = append(contacts, Contact{id, name, month, day, year, true})
	}

	return contacts, nil
}
