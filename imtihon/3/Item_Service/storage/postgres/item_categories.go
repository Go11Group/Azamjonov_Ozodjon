package postgres

import (
	"fmt"
	pb "github.com/imtihon/3/Item_Service/generated/generated/items_service"
)

func (r *ItemRepo) CreateItemCategory(request *pb.CreateItemCategoryManageRequest) (*pb.CreateItemCategoryManageResponse, error) {
	query := `INSERT INTO item_categories (name, description) VALUES ($1, $2) RETURNING id, name, description, created_at`
	row := r.DB.QueryRow(query, request.Name, request.Description)

	var response pb.ItemCategory
	err := row.Scan(&response.Id, &response.Name, &response.Description, &response.CreatedAt)
	if err != nil {
		r.lg.Error(fmt.Sprintf("Error adding item category: %v", err))
		return nil, err
	}

	return &pb.CreateItemCategoryManageResponse{ItemCategory: &response}, nil
}
func (r *ItemRepo) GetStatistics(request *pb.GetStatisticsRequest) (*pb.GetStatisticsResponse, error) {
	startDate, endDate := request.StartDate, request.EndDate

	var totalSwaps, totalRecycledItems, totalEcoPointsEarned int32

	// Get total number of swaps
	err := r.DB.QueryRow(`SELECT COUNT(*) FROM swaps WHERE created_at BETWEEN $1 AND $2`, startDate, endDate).Scan(&totalSwaps)
	if err != nil {
		r.lg.Error(fmt.Sprintf("Error getting total swaps: %v", err))
		return nil, err
	}

	// Get total number of recycled items
	err = r.DB.QueryRow(`SELECT COUNT(*) FROM recycled_items WHERE created_at BETWEEN $1 AND $2`, startDate, endDate).Scan(&totalRecycledItems)
	if err != nil {
		r.lg.Error(fmt.Sprintf("Error getting total recycled items: %v", err))
		return nil, err
	}

	// Get total eco points earned
	err = r.DB.QueryRow(`SELECT SUM(eco_points) FROM eco_points WHERE created_at BETWEEN $1 AND $2`, startDate, endDate).Scan(&totalEcoPointsEarned)
	if err != nil {
		r.lg.Error(fmt.Sprintf("Error getting total eco points earned: %v", err))
		return nil, err
	}

	// Get top categories based on swap counts
	rows, err := r.DB.Query(`SELECT id, name, COUNT(*) as swap_count FROM item_categories
	JOIN swaps ON item_categories.id = swaps.category_id WHERE swaps.created_at BETWEEN $1 AND $2
	GROUP BY id, name ORDER BY swap_count DESC LIMIT 10`, startDate, endDate)
	if err != nil {
		r.lg.Error(fmt.Sprintf("Error getting top categories: %v", err))
		return nil, err
	}
	defer rows.Close()

	var topCategories []*pb.TopCategory
	for rows.Next() {
		var category pb.TopCategory
		err := rows.Scan(&category.Id, &category.Name, &category.SwapCount)
		if err != nil {
			r.lg.Error(fmt.Sprintf("Error scanning category: %v", err))
			return nil, err
		}
		topCategories = append(topCategories, &category)
	}

	// Get top recycling centers based on submission counts
	rows, err = r.DB.Query(`SELECT id, name, COUNT(*) as submissions_count FROM recycling_centers
	JOIN recycled_items ON recycling_centers.id = recycled_items.center_id WHERE recycled_items.created_at BETWEEN $1 AND $2
	GROUP BY id, name ORDER BY submissions_count DESC LIMIT 10`, startDate, endDate)
	if err != nil {
		r.lg.Error(fmt.Sprintf("Error getting top recycling centers: %v", err))
		return nil, err
	}
	defer rows.Close()

	var topRecyclingCenters []*pb.TopRecyclingCenter
	for rows.Next() {
		var center pb.TopRecyclingCenter
		err := rows.Scan(&center.Id, &center.Name, &center.SubmissionCount)
		if err != nil {
			r.lg.Error(fmt.Sprintf("Error scanning center: %v", err))
			return nil, err
		}
		topRecyclingCenters = append(topRecyclingCenters, &center)
	}

	response := &pb.GetStatisticsResponse{
		TotalSwaps:           totalSwaps,
		TotalRecycledItems:   totalRecycledItems,
		TotalEcoPointsEarned: totalEcoPointsEarned,
		TopCategories:        topCategories,
		TopRecyclingCenters:  topRecyclingCenters,
	}

	return response, nil
}

func (r *ItemRepo) GetMonitoringUserActivity(request *pb.GetMonitoringUserActivityRequest) (*pb.GetMonitoringUserActivityResponse, error) {
	var swapInitiated, swapCompleted, itemListed, recyclingSubmissions, ecoPointsEarned int32

	err := r.DB.QueryRow(`SELECT COUNT(*) FROM swaps WHERE owner_id = $1 AND created_at BETWEEN $2 AND $3`, request.UserId, request.StartDate, request.EndDate).Scan(&swapInitiated)
	if err != nil {
		return nil, err
	}

	err = r.DB.QueryRow(`SELECT COUNT(*) FROM swaps WHERE owner_id = $1 AND completed_at BETWEEN $2 AND $3`, request.UserId, request.StartDate, request.EndDate).Scan(&swapCompleted)
	if err != nil {
		return nil, err
	}

	err = r.DB.QueryRow(`SELECT COUNT(*) FROM items WHERE owner_id = $1 AND listed_at BETWEEN $2 AND $3`, request.UserId, request.StartDate, request.EndDate).Scan(&itemListed)
	if err != nil {
		return nil, err
	}

	err = r.DB.QueryRow(`SELECT COUNT(*) FROM recycled_items WHERE user_id = $1 AND created_at BETWEEN $2 AND $3`, request.UserId, request.StartDate, request.EndDate).Scan(&recyclingSubmissions)
	if err != nil {
		return nil, err
	}

	err = r.DB.QueryRow(`SELECT SUM(eco_points) FROM eco_points WHERE user_id = $1 AND created_at BETWEEN $2 AND $3`, request.UserId, request.StartDate, request.EndDate).Scan(&ecoPointsEarned)
	if err != nil {
		return nil, err
	}

	response := &pb.GetMonitoringUserActivityResponse{
		UserId:               request.UserId,
		SwapInitiated:        swapInitiated,
		SwapCompleted:        swapCompleted,
		ItemListed:           itemListed,
		RecyclingSubmissions: recyclingSubmissions,
		EcoPointsEarned:      ecoPointsEarned,
	}

	return response, nil
}
