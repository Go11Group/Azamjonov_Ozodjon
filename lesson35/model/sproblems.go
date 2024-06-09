package model

type SolvedProblem struct {
	Id        string `json:"id"`
	UserID    int    `json:"user_id"`
	ProblemID int    `json:"problem_id"`
	SolveDate string `json:"solve_date"`
}
