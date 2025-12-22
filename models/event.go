package models

import (
	"time"

	"github.com/event-backend-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int       `json:"user_id"`
}

func (e *Event) Save() error {

	querty := `INSERT INTO events (title, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)`

	stm, err := db.DB.Prepare(querty)

	if err != nil {
		return err
	}

	result, err := stm.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserID)

	defer stm.Close()

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id

	return err

}

func GetAllEvents() ([]Event, error) {

	query := `SELECT id, title, description, location, date_time, user_id FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events := []Event{}
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}
	return events, nil

}
func GetEventById(id int64) (*Event, error) {
	query := `SELECT id, title, description, location, date_time, user_id FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	var e Event
	err := row.Scan(&e.ID, &e.Title, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
func (e *Event) Update() error {
	query := `UPDATE events SET title = ?, description = ?, location = ?, date_time = ? WHERE id = ?`

	stm, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stm.Exec(e.Title, e.Description, e.Location, e.DateTime, e.ID)
	defer stm.Close()

	if err != nil {
		return err
	}
	return nil
}

func DeleteEvent(id int64) error {
	query := `DELETE FROM events WHERE id = ?`

	stm, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stm.Exec(id)
	defer stm.Close()
	if err != nil {
		return err
	}
	return nil
}
