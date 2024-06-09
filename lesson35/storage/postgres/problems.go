package postgres

import (
	"database/sql"
	"github.com/Azamjonov_Ozodjon/lesson35/model"
)

type ProblemsRepo struct {
	Db *sql.DB
}

func (p *ProblemsRepo) Create(problems model.Problems) error {
	_, err := p.Db.Exec("INSERT INTO problems (id, description, difficulty) VALUES ($1, $2, $3)",
		problems.Id, problems.Description, problems.Difficulty)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProblemsRepo) GetByID(id string) (model.Problems, error) {
	row := p.Db.QueryRow("SELECT id, description, difficulty FROM problems WHERE id = $1", id)
	var problems model.Problems
	err := row.Scan(&problems.Id, &problems.Description, &problems.Difficulty)
	if err != nil {
		return problems, err
	}
	return problems, nil
}

func (p *ProblemsRepo) Update(problems model.Problems) error {
	_, err := p.Db.Exec("UPDATE problems SET description = $1, difficulty = $2 WHERE id = $3",
		problems.Description, problems.Difficulty, problems.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProblemsRepo) Delete(id string) error {
	_, err := p.Db.Exec("DELETE FROM problems WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
