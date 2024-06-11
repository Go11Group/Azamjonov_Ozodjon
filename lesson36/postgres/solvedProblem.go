package postgres

import (
	"database/sql"
	"github.com/Azamjonov_Ozodjon/lesson36/model"
	_ "github.com/lib/pq"
)

type SolvedProblemRepo struct {
	db *sql.DB
}

func NewSolvedProblemRepo(db *sql.DB) *SolvedProblemRepo {
	return &SolvedProblemRepo{db: db}
}

func (s *SolvedProblemRepo) SolvedProblem_Get(id int) (model.SolvedProblem, error) {
	var solvedProblem model.SolvedProblem
	err := s.db.QueryRow("SELECT id, user_id, problem_id FROM solved_problems WHERE id = $1", id).Scan(
		&solvedProblem.ID, &solvedProblem.UserID, &solvedProblem.ProblemID)
	return solvedProblem, err
}

func (s *SolvedProblemRepo) SolvedProblem_Create(solvedProblem model.SolvedProblem) error {
	_, err := s.db.Exec("INSERT INTO solved_problems (user_id, problem_id) VALUES ($1, $2)",
		solvedProblem.UserID, solvedProblem.ProblemID)
	return err
}

func (s *SolvedProblemRepo) SolvedProblem_Update(id int, solvedProblem model.SolvedProblem) error {
	_, err := s.db.Exec("UPDATE solved_problems SET user_id = $1, problem_id = $2 WHERE id = $3",
		solvedProblem.UserID, solvedProblem.ProblemID, id)
	return err
}

func (s *SolvedProblemRepo) SolvedProblem_Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM solved_problems WHERE id = $1", id)
	return err
}
