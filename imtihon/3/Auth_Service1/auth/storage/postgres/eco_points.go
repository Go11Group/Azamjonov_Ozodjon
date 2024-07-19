package postgres

import (
	pb "auth_service/generated/genproto"
	"database/sql"
	"log"
)

type EcoPointsRepo struct {
	db *sql.DB
}

func NewEcoPointsRepo(db *sql.DB) *EcoPointsRepo {
	return &EcoPointsRepo{db: db}
}

func (r *EcoPointsRepo) GetEcoPoints(req *pb.ById) (*pb.EcoPointsBalance, error) {
	query := `SELECT user_id, SUM(points) as eco_points, MAX(timestamp) as last_updated 
              FROM ecopoints 
              WHERE user_id = $1 
              GROUP BY user_id`

	row := r.db.QueryRow(query, req.Id)

	res := &pb.EcoPointsBalance{}
	err := row.Scan(&res.UserId, &res.EcoPoints, &res.LastUpdated)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (r *EcoPointsRepo) AddEcoPoints(req *pb.AddEcoPointsReq) (*pb.AddEcoPointsRes, error) {
	query := `INSERT INTO ecopoints (user_id, points, reason, type, timestamp) 
              VALUES ($1, $2, $3, 'earned', CURRENT_TIMESTAMP)
              RETURNING user_id, points, reason, timestamp`

	row := r.db.QueryRow(query, req.UserId, req.Points, req.Reason)

	res := &pb.AddEcoPointsRes{}
	err := row.Scan(&res.UserId, &res.AddedPoints, &res.Reason, &res.Timestamp)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Update the total eco_points in the response
	balance, err := r.GetEcoPoints(&pb.ById{Id: req.UserId})
	if err != nil {
		return nil, err
	}
	res.EcoPoints = balance.EcoPoints

	return res, nil
}

func (r *EcoPointsRepo) GetEcoPointsHistory(req *pb.PageLimit1) (*pb.EcoPointsHistoryRes, error) {
	offset := (req.Page - 1) * req.Limit

	query := `SELECT id, points, type, reason, timestamp 
              FROM ecopoints 
              WHERE user_id = $1 
              ORDER BY timestamp DESC 
              LIMIT $2 OFFSET $3`

	rows, err := r.db.Query(query, req.Userid, req.Limit, offset)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var history []*pb.EcoPointsHistoryItem
	for rows.Next() {
		item := &pb.EcoPointsHistoryItem{}
		err := rows.Scan(&item.Id, &item.Points, &item.Type, &item.Reason, &item.Timestamp)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		history = append(history, item)
	}

	query = `SELECT COUNT(*) FROM ecopoints WHERE user_id = $1`
	var total int64
	err = r.db.QueryRow(query, req.Userid).Scan(&total)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &pb.EcoPointsHistoryRes{
		History: history,
		Total:   total,
		Page:    req.Page,
		Limit:   req.Limit,
	}
	return res, nil
}
