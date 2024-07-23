package redis

import (
	"context"
	"encoding/json"
	"lesson62/models"
	"time"

	"github.com/redis/go-redis/v9"
)

type RestaurantRepo struct {
	DB *redis.Client
}

func NewRestaurantRepo(redisClient *redis.Client) *RestaurantRepo {
	return &RestaurantRepo{
		DB: redisClient,
	}
}

// Create
func (r *RestaurantRepo) Create(ctx context.Context, id, restaurant string) error {
	return r.DB.HSet(ctx, "restaurants", id, restaurant).Err()
}

// Get all
func (r *RestaurantRepo) GetAll(ctx context.Context) (*[]models.RestaurantGet, error) {
	res, err := r.DB.HGetAll(ctx, "restaurants").Result()

	resp := []models.RestaurantGet{}
	for _, val := range res {
		restaurant := models.RestaurantGet{}
		err := json.Unmarshal([]byte(val), &restaurant)
		if err != nil {
			return nil, err
		}
		resp = append(resp, restaurant)
	}
	return &resp, err
}

// Get by Id
func (r *RestaurantRepo) GetById(ctx context.Context, id string) (*models.RestaurantGet, error) {
	restaurant, err := r.DB.HGet(ctx, "restaurants", id).Result()
	if err != nil {
		return nil, err
	}
	resp := models.RestaurantGet{}
	err = json.Unmarshal([]byte(restaurant), &resp)
	return &resp, err
}

// Update
func (r *RestaurantRepo) Update(ctx context.Context, restaurant *models.RestaurantUpdate) error {
	resp, err := r.GetById(ctx, restaurant.Id)
	if err != nil {
		return err
	}

	restaurantToUpdate := models.RestaurantGet{
		Id:          resp.Id,
		Title:       restaurant.Title,
		Description: restaurant.Description,
		Address:     restaurant.Address,
		ProneNumber: restaurant.ProneNumber,
		Time: models.Time{
			CreatedAt: resp.CreatedAt,
			UpdatedAt: time.Now(),
		},
	}

	req, err := json.Marshal(restaurantToUpdate)
	if err != nil {
		return err
	}

	err = r.DB.HSet(ctx, "restaurants", restaurant.Id, req).Err()
	return err
}

// Delete
func (r *RestaurantRepo) Delete(ctx context.Context, id string) error {
	return r.DB.HDel(ctx, "restaurants", id).Err()
}
