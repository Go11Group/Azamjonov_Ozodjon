package postgres

import (
	"fmt"
	"time"

	pb "github.com/imtihon/3/Item_Service/generated/generated/items_service"
)

func (r *ItemRepo) CreateRecyclingSubmission(request *pb.CreateRecyclingSubmissionsRequest) (*pb.CreateRecyclingSubmissionsResponse, error) {
	query := `INSERT INTO recycling_submissions (center_id, user_id, eco_points_earned, created_at)
	          VALUES ($1, $2, $3, $4) RETURNING id, center_id, user_id, eco_points_earned, created_at`
	row := r.DB.QueryRow(query, request.CenterId, request.UserId, calculateEcoPoints(request.Itemes), time.Now())

	var response pb.CreateRecyclingSubmissionsResponse
	err := row.Scan(&response.Id, &response.CenterId, &response.UserId, &response.EcoPointsEarned, &response.CreatedAt)
	if err != nil {
		r.lg.Error(fmt.Sprintf("Error creating recycling submission: %v", err))
		return nil, err
	}

	for _, item := range request.Itemes {
		itemQuery := `INSERT INTO recycling_submission_items (submission_id, item_id, weight, material) VALUES ($1, $2, $3, $4)`
		_, err := r.DB.Exec(itemQuery, response.Id, item.ItemId, item.Weight, item.Material)
		if err != nil {
			r.lg.Error(fmt.Sprintf("Error inserting item: %v", err))
			return nil, err
		}
	}
	response.Itemes = request.Itemes
	return &response, nil
}

func calculateEcoPoints(items []*pb.Itemes) int32 {
	var totalPoints int32
	for _, item := range items {
		totalPoints += int32(item.Weight * 10)
	}
	return totalPoints
}
