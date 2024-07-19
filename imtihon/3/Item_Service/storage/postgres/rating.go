package postgres

import (
	"context"
	"fmt"
	pb "github.com/imtihon/3/Item_Service/generated/generated/items_service"
)

func (r *ItemRepo) CreateAddUserRating(request *pb.CreateRatingRequest) (*pb.CreateRatingResponse, error) {
	query := `INSERT INTO ratings (user_id, rater_id, rating, comment, swap_id) 
	          VALUES ($1, $2, $3, $4,$5) RETURNING id, user_id, rater_id, rating, comment, swap_id, created_at`
	row := r.DB.QueryRow(query, request.UserId, request.RaterId, request.Rating, request.Comment, request.SwapId)

	var response pb.Rating
	err := row.Scan(&response.Id, &response.UserId, &response.RaterId, &response.Rating, &response.Comment, &response.SwapId, &response.CreatedAt)
	if err != nil {
		r.lg.Error(fmt.Sprintf("Error adding user rating: %v", err))
		return nil, err
	}

	return &pb.CreateRatingResponse{Rating: &response}, nil
}

func (r *ItemRepo) GetRatings(ctx context.Context, request *pb.GetRatingRequest) (*pb.GetRatingResponse, error) {
	offset := (request.Page - 1) * request.Limit
	rows, err := r.DB.QueryContext(ctx, "SELECT id, rater_id, rating, comment, swap_id, created_at FROM ratings WHERE user_id = $1 LIMIT $2 OFFSET $3", request.UserId, request.Limit, offset)
	if err != nil {
		r.lg.Error(fmt.Sprintf("Error fetching ratings: %v", err))
		return nil, err
	}
	defer rows.Close()

	var ratings []*pb.Rating
	for rows.Next() {
		var rating pb.Rating
		err := rows.Scan(&rating.Id, &rating.RaterId, &rating.Rating, &rating.Comment, &rating.SwapId, &rating.CreatedAt)
		if err != nil {
			r.lg.Error(fmt.Sprintf("Error scanning rating: %v", err))
			return nil, err
		}
		ratings = append(ratings, &rating)
	}

	if err = rows.Err(); err != nil {
		r.lg.Error(fmt.Sprintf("Error iterating over ratings rows: %v", err))
		return nil, err
	}

	// Prepare response
	response := &pb.GetRatingResponse{
		Ratings: ratings,
		Page:    request.Page,
		Limit:   request.Limit,
	}

	return response, nil
}
