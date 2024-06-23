package handler

import (
	"atto/storage/postgres"
	"database/sql"
)

type Handler struct {
	User        *postgres.UserRepo
	Card        *postgres.CardRepo
	Station     *postgres.StationRepo
	Terminal    *postgres.TerminalRepo
	Transaction *postgres.TransactionRepo
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		User:        postgres.NewUserRepo(db),
		Card:        postgres.NewCardRepo(db),
		Station:     postgres.NewStationRepo(db),
		Terminal:    postgres.NewTerminalRepo(db),
		Transaction: postgres.NewTransactionRepo(db),
	}
}
