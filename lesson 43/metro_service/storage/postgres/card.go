package postgres

import (
	"atto/models"
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type CardRepo struct {
	Db *sql.DB
}

func NewCardRepo(db *sql.DB) *CardRepo {
	return &CardRepo{Db: db}
}

func (c *CardRepo) Create(card *models.CreateCard) error {
	_, err := c.Db.Exec("INSERT INTO card (id, number, user_id) VALUES ($1, $2, $3)",
		uuid.NewString(), card.Number, card.UserId)
	return err
}

func (c *CardRepo) GetById(id string) (*models.Card, error) {
	var card models.Card
	err := c.Db.QueryRow("SELECT id, number, user_id FROM card WHERE id = $1 and dele", id).Scan(&card.Id, &card.Number, &card.UserId)
	if err != nil {
		return nil, err
	}
	return &card, nil
}

func (c *CardRepo) Update(card *models.Card) error {
	_, err := c.Db.Exec("UPDATE card SET number = $2, user_id = $3 WHERE id = $1",
		card.Id, card.Number, card.UserId)
	return err
}

func (c *CardRepo) Delete(id string) error {
	_, err := c.Db.Exec("UPDATE card SET deleted_at = $2 WHERE id = $1", id, time.Now())
	return err
}

func (c *CardRepo) GetAll() ([]models.Card, error) {
	rows, err := c.Db.Query("SELECT id, number, user_id FROM card")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []models.Card
	for rows.Next() {
		var card models.Card
		if err := rows.Scan(&card.Id, &card.Number, &card.UserId); err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}
