// TO PREVENT FROM SQL INJECTIONS WE ARE SETTING VALUES ????
// USING QUERY INSTEAD OF EXEC ->> QUERY IS TYPICALLY USED IF WE HAVE A QUERY WHERE WE WANT TO GET BUNCH OF ROWS. WHICH WE WANT TO USE. FETCHES DATA
// WHEREAS EXEC IS USED WHEN WE HAVE A QUERY THAT CHANGES DATA IN THE DATABASE. INSERTS DATA UPDATES DATA AND SO-ON. CHANGES DATA.

package models

import (
	"rest-api/m/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int64
}

var events = []Event{} // var events added to the slice Event

// methods for the above struct
func (e *Event) Save() error {
	// for db
	query := `
	INSERT INTO events(name, description, location, dataTime, user_id) 
	VALUES (?,?,?,?,?)`
	stmnt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmnt.Close()                                                                // close it after the execution or creation by using defer keyword.
	result, err := stmnt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID) // insertion of values in sql in a safe way.
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

// fetching rows from the event table
func GetAllWEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err // returing nil value for the slice of events and error if we have any err.
	}
	// when func is done close it.
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

// sending query to db for single event.
func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id =? " // setting it to ? to prevent from sql injections
	row := db.DB.QueryRow(query, id)             // queryrow method as we know in id theer will be only a single row. so we can take advantage by using queryrow()

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	// in case of success
	return &event, nil
}

// update method.
func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

// delete method
func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	return err
}

// method for new registration of table.
func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

// method for deletion of the registering table.

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ? "
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}
