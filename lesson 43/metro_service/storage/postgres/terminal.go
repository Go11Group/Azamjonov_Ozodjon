package postgres

import (
	"atto/models"
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type TerminalRepo struct {
	Db *sql.DB
}

func NewTerminalRepo(db *sql.DB) *TerminalRepo {
	return &TerminalRepo{Db: db}
}

func (t *TerminalRepo) Create(terminal *models.CreateTerminal) error {
	_, err := t.Db.Exec("INSERT INTO terminal (id, station_id) VALUES ($1, $2)",
		uuid.NewString(), terminal.StationId)
	return err
}

func (t *TerminalRepo) GetById(id string) (*models.Terminal, error) {
	var terminal models.Terminal
	err := t.Db.QueryRow("SELECT id, station_id FROM terminal WHERE id = $1", id).Scan(&terminal.Id, &terminal.StationId)
	if err != nil {
		return nil, err
	}
	return &terminal, nil
}

func (t *TerminalRepo) Update(terminal *models.Terminal) error {
	_, err := t.Db.Exec("UPDATE terminal SET station_id = $2 WHERE id = $1", terminal.Id, terminal.StationId)
	return err
}

func (t *TerminalRepo) Delete(id string) error {
	_, err := t.Db.Exec("UPDATE terminal SET deleted_at = $2 WHERE id = $1", id, time.Now())
	return err
}

func (t *TerminalRepo) GetAll() ([]models.Terminal, error) {
	rows, err := t.Db.Query("SELECT id, station_id FROM terminal")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var terminals []models.Terminal
	for rows.Next() {
		var terminal models.Terminal
		if err := rows.Scan(&terminal.Id, &terminal.StationId); err != nil {
			return nil, err
		}
		terminals = append(terminals, terminal)
	}
	return terminals, nil
}
