package redis

import (
	"context"
	"encoding/json"
	"lesson62/models"
	"time"

	"github.com/redis/go-redis/v9"
)

type UsersRepo struct {
	DB *redis.Client
}

func NewUsersClient(redisClient *redis.Client) *UsersRepo {
	return &UsersRepo{
		DB: redisClient,
	}
}

// Create
func (u *UsersRepo) Create(ctx context.Context, id, user string) error {
	return u.DB.HSet(ctx, "users", id, user).Err()
}

// Get all
func (u *UsersRepo) GetAll(ctx context.Context) (*[]models.UserGet, error) {
	res, err := u.DB.HGetAll(ctx, "users").Result()

	resp := []models.UserGet{}
	for _, val := range res {
		user := models.UserGet{}
		err := json.Unmarshal([]byte(val), &user)
		if err != nil {
			return nil, err
		}
		resp = append(resp, user)
	}
	return &resp, err
}

// Get by Id
func (u *UsersRepo) GetById(ctx context.Context, id string) (*models.UserGet, error) {
	user, err := u.DB.HGet(ctx, "users", id).Result()
	if err != nil {
		return nil, err
	}
	resp := models.UserGet{}
	err = json.Unmarshal([]byte(user), &resp)
	return &resp, err
}

// Update
func (u *UsersRepo) Update(ctx context.Context, user *models.UserUpdate) error {
	resp, err := u.GetById(ctx, user.Id)
	if err != nil {
		return err
	}

	userToUpdate := models.UserGet{
		Id:        resp.Id,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Age:       user.Age,
		Role:      user.Role,
		Time: models.Time{
			CreatedAt: resp.CreatedAt,
			UpdatedAt: time.Now(),
		},
	}

	req, err := json.Marshal(userToUpdate)
	if err != nil {
		return err
	}

	err = u.DB.HSet(ctx, "users", user.Id, req).Err()
	return err
}

// Delete
func (u *UsersRepo) Delete(ctx context.Context, id string) error {
	return u.DB.HDel(ctx, "users", id).Err()
}
