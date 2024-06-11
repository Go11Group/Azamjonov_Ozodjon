package postgres

import (
	"database/sql"
	"github.com/Azamjonov_Ozodjon/lesson36/model"
	_ "github.com/lib/pq"
)

type ProblemRepo struct {
	db *sql.DB
}

func NewProblemRepo(db *sql.DB) *ProblemRepo {
	return &ProblemRepo{db: db}
}

func (p *ProblemRepo) Problem_Get(id int) (model.Problem, error) {
	var problem model.Problem
	err := p.db.QueryRow("SELECT id, title, description, difficulty FROM problems WHERE id = $1", id).Scan(
		&problem.ID, &problem.Title, &problem.Description, &problem.Difficulty)
	return problem, err
}

func (p *ProblemRepo) Problem_Create(problem model.Problem) error {
	_, err := p.db.Exec("INSERT INTO problems (title, description, difficulty) VALUES ($1, $2, $3)",
		problem.Title, problem.Description, problem.Difficulty)

	return err
}

func (p *ProblemRepo) Problem_Update(id int, problem model.Problem) error {
	_, err := p.db.Exec("UPDATE problems SET title = $1, description = $2, difficulty = $3 WHERE id = $4",
		problem.Title, problem.Description, problem.Difficulty, id)
	return err
}

func (p *ProblemRepo) Problem_Delete(id int) error {
	_, err := p.db.Exec("DELETE FROM problems WHERE id = $1", id)
	return err
}
