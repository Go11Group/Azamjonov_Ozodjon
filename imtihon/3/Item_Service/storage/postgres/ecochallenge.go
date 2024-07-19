package postgres

import (
	pb "github.com/imtihon/3/Item_Service/generated/generated/items_service"
	"time"
)

func (r *ItemRepo) CreateEcoChallenge(request *pb.CreateEcoChallengeRequest) (*pb.CreateEcoChallengeResponse, error) {
	var challenge pb.EcoChallenge

	query := `INSERT INTO eco_challenges (title, description, start_date, end_date, reward_points, created_at)
              VALUES ($1, $2, $3, $4, $5, $6)
              RETURNING id, title, description, start_date, end_date, reward_points, created_at;`

	err := r.DB.QueryRow(query, request.Title, request.Description, request.StartDate, request.EndDate, request.RewardPoints, time.Now().UTC()).
		Scan(&challenge.Id, &challenge.Title, &challenge.Description, &challenge.StartDate, &challenge.EndDate, &challenge.RewardPoints, &challenge.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &pb.CreateEcoChallengeResponse{
		EcoChallenge: &challenge,
	}, nil
}
