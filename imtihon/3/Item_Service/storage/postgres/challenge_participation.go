package postgres

import (
	"time"

	pb "github.com/imtihon/3/Item_Service/generated/generated/items_service"
)

// ParticipateChallenge handles inserting a new participation record into the challenge_participation table
func (r *ItemRepo) ParticipateChallenge(req *pb.CreateParticipateChallengeRequest) (*pb.CreateParticipateChallengeResponse, error) {
	var res pb.CreateParticipateChallengeResponse

	query := `INSERT INTO challenge_participation (challenge_id, user_id, status, joined_at)
              VALUES ($1, $2, $3, $4)
              RETURNING challenge_id, user_id, status, joined_at;`

	err := r.DB.QueryRow(query, req.ChallengeId, req.UserId, req.Status, time.Now().UTC()).
		Scan(&res.ChallengeId, &res.UserId, &res.Status, &res.JoinedAt)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// UpdateEcoChallengeResult handles updating the result of an eco challenge in the challenge_participation table
func (r *ItemRepo) UpdateEcoChallengeResult(req *pb.UpdateEcoChallengeResultRequest) (*pb.UpdateEcoChallengeResultResponse, error) {
	var res pb.UpdateEcoChallengeResultResponse

	query := `UPDATE challenge_participation 
              SET recycled_items_count = $1, updated_at = $2 
              WHERE challenge_id = $3 AND deleted_at = 0
              RETURNING challenge_id, user_id, status, recycled_items_count, updated_at;`

	err := r.DB.QueryRow(query, req.RecycledItemsCount, time.Now().UTC(), req.ChallengeId).
		Scan(&res.ChallengeId, &res.UserId, &res.Status, &res.RecycledItemsCount, &res.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
