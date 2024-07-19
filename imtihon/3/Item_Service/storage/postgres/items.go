package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/imtihon/3/Item_Service/generated/generated/items_service"
	"github.com/imtihon/3/Item_Service/logs"
	"log/slog"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type ItemRepo struct {
	pb.UnimplementedItemsServiceServer
	DB *sql.DB
	lg *slog.Logger
}

func NewItemRepo(db *sql.DB) *ItemRepo {

	return &ItemRepo{
		DB: db,
		lg: logs.InitLogger(),
	}
}

// CreateItem creates a new item in the database
func (r *ItemRepo) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	query := `
        INSERT INTO items (name, description, category_id, condition, swap_preference, status, owner_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id, name, description, category_id, condition, swap_preference, status, created_at;`

	swapPreference, err := json.Marshal(req.SwapPreference)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal swap preference: %v", err)
	}

	item := &pb.Item{}
	err = r.DB.QueryRowContext(ctx, query, req.Name, req.Description, req.CategoryId, req.Condition, swapPreference, req.Status, req.OwnerId).Scan(
		&item.Id, &item.Name, &item.Description, &item.CategoryId, &item.Condition, &item.SwapPreference, &item.Status, &item.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to create item: %v", err)
	}
	return &pb.CreateItemResponse{Item: item}, nil
}

// GetItem retrieves an item by ID from the database
func (r *ItemRepo) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	query := `SELECT id, name, description, category_id, condition, swap_preference, owner_id, status, created_at FROM items WHERE id = $1;`
	item := &pb.Item{}
	err := r.DB.QueryRowContext(ctx, query, req.Id).Scan(
		&item.Id, &item.Name, &item.Description, &item.CategoryId, &item.Condition, &item.SwapPreference, &item.OwnerId, &item.Status, &item.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("item not found")
	} else if err != nil {
		return nil, fmt.Errorf("failed to get item: %v", err)
	}
	return &pb.GetItemResponse{Item: item}, nil
}

// UpdateItem updates an existing item in the database
func (r *ItemRepo) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
	query := `
		UPDATE items
		SET name = $1,condition = $2,updated_at = CURRENT_TIMESTAMP
		WHERE id = $3
		RETURNING id, name, description, category_id, condition, swap_preference, owner_id, status;`

	item := &pb.Item{}
	err := r.DB.QueryRowContext(ctx, query, req.Name, req.Condition, req.Id).Scan(
		&item.Id, &item.Name, &item.Description, &item.CategoryId, &item.Condition, &item.SwapPreference, &item.OwnerId, &item.Status)

	if err == sql.ErrNoRows {
		return nil, errors.New("item not found")
	} else if err != nil {
		return nil, fmt.Errorf("failed to update item: %v", err)
	}
	return &pb.UpdateItemResponse{Item: item}, nil
}

func (r *ItemRepo) GetAllItems(ctx context.Context, req *pb.GetAllItemsRequest) (*pb.GetAllItemsResponse, error) {
	offset := (req.Page - 1) * req.Limit

	query := "SELECT id, name, category_id, condition, owner_id, status FROM items WHERE deleted_at = 0 LIMIT $1 OFFSET $2"
	fmt.Println(req.Limit)
	rows, err := r.DB.QueryContext(ctx, query, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*pb.Item
	for rows.Next() {
		var item pb.Item
		if err := rows.Scan(&item.Id, &item.Name, &item.CategoryId, &item.Condition, &item.OwnerId, &item.Status); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return &pb.GetAllItemsResponse{Items: items}, nil
}

// DeleteItem deletes an item by ID from the database
func (r *ItemRepo) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	query := `
		UPDATE 
			items 
		SET 
			deleted_at = EXTRACT(EPOCH FROM CURRENT_TIMESTAMP)
		WHERE 
			id = $1 AND deleted_at = 0
	`

	_, err := r.DB.Exec(query, req.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("item not found")
		}
		return nil, fmt.Errorf("failed to delete item: %v", err)
	}
	return &pb.DeleteItemResponse{Message: "Item successfully deleted"}, nil
}

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return namedQuery, args
}

func (r *ItemRepo) SearchItemsAndFilter(request *pb.SearchItemsAndFilterRequest) (*pb.SearchItemsAndFilterResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	filter := " WHERE deleted_at IS NULL"
	if len(request.Query) > 0 {
		params["query"] = "%" + request.Query + "%"
		filter += " AND (name ILIKE :query OR description ILIKE :query)"
	}
	if len(request.Category) > 0 {
		params["category"] = request.Category
		filter += " AND category_id = :category"
	}
	if len(request.Condition) > 0 {
		params["condition"] = request.Condition
		filter += " AND condition = :condition"
	}
	if request.Page > 0 {
		params["page"] = request.Page
		params["limit"] = request.Limit
		params["offset"] = (request.Page - 1) * request.Limit
		filter += " LIMIT :limit OFFSET :offset"
	}

	query := "SELECT id, name, description, category_id, condition, swap_preference, owner_id, status, listed_at, created_at, updated_at FROM items" + filter
	query, arr = ReplaceQueryParams(query, params)

	rows, err := r.DB.Query(query, arr...)
	if err != nil {
		r.lg.Error(fmt.Sprintf("message search items and filter error -> %v", err))
		return nil, err
	}
	defer rows.Close()

	var items []*pb.Item
	for rows.Next() {
		var item pb.Item
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.CategoryId, &item.Condition, &item.SwapPreference, &item.OwnerId, &item.Status, &item.ListedAt, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			r.lg.Error(fmt.Sprintf("message search items and filter scan error -> %v", err))
			return nil, err
		}
		items = append(items, &item)
	}
	return &pb.SearchItemsAndFilterResponse{Items: items}, nil
}
