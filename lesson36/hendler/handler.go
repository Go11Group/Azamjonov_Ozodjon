package handler

import "github.com/Azamjonov_Ozodjon/lesson36/postgres"

type Handler struct {
	User          *postgres.UserRepo
	Problem       *postgres.ProblemRepo
	SolvedProblem *postgres.SolvedProblemRepo
}
