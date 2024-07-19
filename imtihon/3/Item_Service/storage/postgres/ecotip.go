package postgres

import (
	pb "github.com/imtihon/3/Item_Service/generated/generated/items_service"
)

func (r *ItemRepo) CreateAddEcoTips(request *pb.CreateEcoTipRequest) (*pb.CreateEcoTipResponse, error) {
	var response pb.EcoTip

	query := `INSERT INTO eco_tips (title, content, created_at) VALUES ($1, $2, NOW())
              RETURNING id, title, content, created_at;`

	err := r.DB.QueryRow(query, request.Title, request.Content).
		Scan(&response.Id, &response.Title, &response.Content, &response.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &pb.CreateEcoTipResponse{EcoTip: &response}, nil
}

func (r *ItemRepo) GetAddEcoTips(request *pb.GetAddEcoTipsRequest) (*pb.GetAddEcoTipsResponse, error) {
	var response pb.GetAddEcoTipsResponse
	var tips []*pb.EcoTip

	query := `SELECT id, title, content, created_at FROM eco_tips ORDER BY created_at DESC LIMIT $1 OFFSET $2;`
	rows, err := r.DB.Query(query, request.Limit, (request.Page-1)*request.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tip pb.EcoTip
		if err := rows.Scan(&tip.Id, &tip.Title, &tip.Content, &tip.CreatedAt); err != nil {
			return nil, err
		}
		tips = append(tips, &tip)
	}

	countQuery := `SELECT COUNT(*) FROM eco_tips;`
	var total int32
	err = r.DB.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, err
	}

	response.Tips = tips
	response.Total = total
	response.Page = request.Page
	response.Limit = request.Limit

	return &response, nil
}
