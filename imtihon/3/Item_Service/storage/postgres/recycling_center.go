package postgres

//func (r *ItemRepo) CreateAddRecyclingCenter(request *pb.CreateRecyclingCenterRequest) (*pb.CreateRecyclingCenterResponse, error) {
//	query := `INSERT INTO recycling_centers (name, address, accepted_materials, working_hours, contact_number, created_at)
//	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, address, accepted_materials, working_hours, contact_number, created_at`
//	//acceptedMaterials, err := json.Marshal(request.AcceptedMaterials)
//	//if err != nil {
//	//	return nil, fmt.Errorf("Failed to marshal accepted_materials: %w", err)
//	//}
//
//	row := r.DB.QueryRow(query, request.Name, request.Address, request.AcceptedMaterials, request.WorkingHours, request.ContactNumber, time.Now())
//	fmt.Println(request)
//	var response pb.RecyclingCenter
//	err := row.Scan(&response.Id, &response.Name, &response.Address, &response.AcceptedMaterials, &response.WorkingHours, &response.ContactNumber, &response.CreatedAt)
//	if err != nil {
//		r.lg.Error(fmt.Sprintf("Error creating recycling center: %v", err))
//		return nil, err
//	}
//
//	return &pb.CreateRecyclingCenterResponse{RecyclingCenter: &response}, nil
//}

//func (r *ItemRepo) SearchRecyclingCenter(request *pb.SearchRecyclingCenterRequest) (*pb.SearchRecyclingCenterResponse, error) {
//	offset := (request.Page - 1) * request.Limit
//	rows, err := r.DB.Query("SELECT id, name, address, accepted_materials, working_hours, contact_number FROM recycling_centers WHERE accepted_materials ILIKE $1 LIMIT $2 OFFSET $3", "%"+request.Material+"%", request.Limit, offset)
//	if err != nil {
//		r.lg.Error(fmt.Sprintf("Error searching recycling centers: %v", err))
//		return nil, err
//	}
//	defer rows.Close()
//
//	var centers []*pb.RecyclingCenter
//	for rows.Next() {
//		var center pb.RecyclingCenter
//		err := rows.Scan(&center.Id, &center.Name, &center.Address, &center.AcceptedMaterials, &center.WorkingHours, &center.ContactNumber)
//		if err != nil {
//			r.lg.Error(fmt.Sprintf("Error scanning recycling center row: %v", err))
//			return nil, err
//		}
//		centers = append(centers, &center)
//	}
//
//	if err = rows.Err(); err != nil {
//		r.lg.Error(fmt.Sprintf("Error iterating over recycling center rows: %v", err))
//		return nil, err
//	}
//
//	var total int32
//	err = r.DB.QueryRow("SELECT COUNT(*) FROM recycling_centers WHERE accepted_materials ILIKE $1", "%"+request.Material+"%").Scan(&total)
//	if err != nil {
//		r.lg.Error(fmt.Sprintf("Error counting recycling centers: %v", err))
//		return nil, err
//	}
//
//	response := &pb.SearchRecyclingCenterResponse{
//		RecyclingCenters: centers,
//		Total:            total,
//		Page:             request.Page,
//		Limit:            request.Limit,
//	}
//
//	return response, nil
//}
