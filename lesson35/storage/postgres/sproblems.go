package postgres

import (
	"database/sql"
	"github.com/Azamjonov_Ozodjon/lesson35/model"
)

type SolvedProblemsRepo struct {
	Db *sql.DB
}

func (s *SolvedProblemsRepo) Create(solvedProblem model.SolvedProblem) error {
	// Execute the SQL INSERT statement
	_, err := s.Db.Exec("INSERT INTO Solved_Problems (user_id, problem_id, solve_date) VALUES ($1, $2, $3)",
		solvedProblem.UserID, solvedProblem.ProblemID, solvedProblem.SolveDate)
	if err != nil {
		return err
	}
	return nil
}

func (s *SolvedProblemsRepo) GetByID(id string) (*model.SolvedProblem, error) {
	row := s.Db.QueryRow("SELECT id, user_id, problem_id, solve_date FROM Solved_Problems WHERE id = $1", id)

	var sp model.SolvedProblem
	err := row.Scan(&sp.Id, &sp.UserID, &sp.ProblemID, &sp.SolveDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No record found
		}
		return nil, err
	}

	return &sp, nil
}

func (s *SolvedProblemsRepo) UpdateByID(id string, solvedProblem model.SolvedProblem) error {
	_, err := s.Db.Exec("UPDATE Solved_Problems SET user_id = $1, problem_id = $2, solve_date = $3 WHERE id = $4",
		solvedProblem.UserID, solvedProblem.ProblemID, solvedProblem.SolveDate, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SolvedProblemsRepo) DeleteByID(id string) error {
	_, err := s.Db.Exec("DELETE FROM Solved_Problems WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
