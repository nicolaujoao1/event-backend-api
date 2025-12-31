package models

import (
	"github.com/event-backend-api/db"
)

func (e Event) Register(userId int64) error {

	querty := `INSERT INTO registrations (event_id, user_id) VALUES (?, ?)`

	stm, err := db.DB.Prepare(querty)

	if err != nil {
		return err
	}

	result, err := stm.Exec(e.ID, userId)

	defer stm.Close()

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id

	return err
}
func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id=? AND user_id=?"
	stm, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stm.Close()

	if _, err = stm.Exec(e.ID, userId); err != nil {
		return err
	}

	return err
}
