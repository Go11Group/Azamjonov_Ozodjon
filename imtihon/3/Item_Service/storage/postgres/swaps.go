package postgres

import (
	"fmt"
	pb "github.com/imtihon/3/Item_Service/generated/generated/items_service"
	"time"
)

func (r *ItemRepo) CreateChangeSwaps(request *pb.CreateSwapRequest) (*pb.CreateSwapResponse, error) {
	_, err := r.DB.Exec("INSERT INTO swaps (offered_item_id, requested_item_id, requester_id, owner_id, status,message) VALUES ($1, $2, $3, $4, $5,$6)", request.OfferedItemId, request.RequestedItemId, request.RequesterId, request.OwnerId, request.Status, request.Message)
	if err != nil {
		r.lg.Error(fmt.Sprintf("message create Swaps error -> %v", err))
		return nil, err
	}
	var response pb.Swap
	err = r.DB.QueryRow("SELECT id, offered_item_id, requested_item_id, requester_id, owner_id, status, updated_at FROM swaps WHERE offered_item_id = $1", request.OfferedItemId).Scan(
		&response.Id, &response.OfferedItemId, &response.RequestedItemId, &response.RequesterId, &response.OwnerId, &response.Status, &response.UpdatedAt)
	if err != nil {
		r.lg.Error(fmt.Sprintf("message get updated swap error -> %v", err))
		return nil, err
	}
	return &pb.CreateSwapResponse{Swap: &response}, nil
}

func (r *ItemRepo) UpdateAcceptSwap(request *pb.UpdateSwapRequest) (*pb.UpdateSwapResponse, error) {
	_, err := r.DB.Exec("UPDATE swaps SET status = $1, updated_at = $2 WHERE id = $3", "accepted", time.Now(), request.Id)
	if err != nil {
		r.lg.Error(fmt.Sprintf("message update accept swap error -> %v", err))
		return nil, err
	}

	var response pb.Swap
	err = r.DB.QueryRow("SELECT id, offered_item_id, requested_item_id, requester_id, owner_id, status, updated_at FROM swaps WHERE id = $1", request.Id).Scan(
		&response.Id, &response.OfferedItemId, &response.RequestedItemId, &response.RequesterId, &response.OwnerId, &response.Status, &response.UpdatedAt)
	if err != nil {
		r.lg.Error(fmt.Sprintf("message get updated swap error -> %v", err))
		return nil, err
	}

	return &pb.UpdateSwapResponse{Swap: &response}, nil
}
func (r *ItemRepo) UpdateRejectSwap(request *pb.UpdateRejectSwapRequest) (*pb.UpdateRejectSwapResponse, error) {
	_, err := r.DB.Exec("UPDATE swaps SET status = $1, message = $2, updated_at = $3 WHERE id = $4", "rejected", request.Reason, time.Now(), request.SwapId)
	if err != nil {
		r.lg.Error(fmt.Sprintf("message update reject swap error -> %v", err))
		return nil, err
	}

	var response pb.Swap
	err = r.DB.QueryRow("SELECT id, offered_item_id, requested_item_id, requester_id, owner_id, status, message, updated_at FROM swaps WHERE id = $1", request.SwapId).Scan(
		&response.Id, &response.OfferedItemId, &response.RequestedItemId, &response.RequesterId, &response.OwnerId, &response.Status, &response.Message, &response.UpdatedAt)
	if err != nil {
		r.lg.Error(fmt.Sprintf("message get updated swap error -> %v", err))
		return nil, err
	}

	return &pb.UpdateRejectSwapResponse{Swap: &response}, nil
}
func (r *ItemRepo) GetChangeSwap(request *pb.GetChangeSwapRequest) (*pb.GetChangeSwapResponse, error) {
	offset := (request.Page - 1) * request.Limit
	rows, err := r.DB.Query("SELECT id, offered_item_id, requested_item_id, requester_id, owner_id, status, created_at FROM swaps WHERE status = $1 LIMIT $2 OFFSET $3", request.Status, request.Limit, offset)
	if err != nil {
		r.lg.Error(fmt.Sprintf("message get change swap error -> %v", err))
		return nil, err
	}
	defer rows.Close()

	var swaps []*pb.Swap
	for rows.Next() {
		var swap pb.Swap
		err := rows.Scan(&swap.Id, &swap.OfferedItemId, &swap.RequestedItemId, &swap.RequesterId, &swap.OwnerId, &swap.Status, &swap.UpdatedAt)
		if err != nil {
			r.lg.Error(fmt.Sprintf("message scan swap error -> %v", err))
			return nil, err
		}
		swaps = append(swaps, &swap)
	}

	if err = rows.Err(); err != nil {
		r.lg.Error(fmt.Sprintf("message rows error -> %v", err))
		return nil, err
	}

	var total int32
	err = r.DB.QueryRow("SELECT COUNT(*) FROM swaps WHERE status = $1", request.Status).Scan(&total)
	if err != nil {
		r.lg.Error(fmt.Sprintf("message count swaps error -> %v", err))
		return nil, err
	}

	response := &pb.GetChangeSwapResponse{
		Swaps: swaps,
		Total: total,
		Page:  request.Page,
		Limit: request.Limit,
	}

	return response, nil
}
