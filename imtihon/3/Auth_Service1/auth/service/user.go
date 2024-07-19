package service

import (
	pb "auth_service/generated/genproto"
	"auth_service/storage/postgres"
	"context"
	"log/slog"
)

type UserService struct {
	db        *postgres.UserAuthRepo
	logger    *slog.Logger
	ecoPoints *postgres.EcoPointsRepo
}

func NewUserService(db *postgres.UserAuthRepo, logger *slog.Logger, ecoPoints *postgres.EcoPointsRepo) *UserService {
	return &UserService{
		db:        db,
		logger:    logger,
		ecoPoints: ecoPoints,
	}
}

// User-related methods

func (s *UserService) RegisterUser(req *pb.RegisterUserReq) (*pb.Register, error) {
	return s.db.RegisterUser(req)
}

func (s *UserService) LoginUser(req *pb.LoginReq) (*pb.Token, error) {
	return s.db.LoginUser(req)
}

func (s *UserService) GetUser(req *pb.ById) (*pb.Register, error) {
	return s.db.GetUser(req)
}

func (s *UserService) UpdateUser(req *pb.UpdateUserReq) (*pb.Register, error) {
	return s.db.UpdateUser(req)
}

func (s *UserService) GetAllUsers(req *pb.PageLimit) (*pb.GetAllUsersRes, error) {
	return s.db.GetAllUsers(req)
}

func (s *UserService) DeleteUser(req *pb.ById) (*pb.Success, error) {
	return s.db.DeleteUser(req)
}

func (s *UserService) ResetPassword(req *pb.ByEmail) (*pb.Success, error) {
	return s.db.ResetPassword(req)
}

func (s *UserService) RefreshToken(req *pb.RefreshTokenReq) (*pb.Token, error) {
	return s.db.RefreshToken(req)
}

// Eco points-related methods

func (s *UserService) GetEcoPoints(req *pb.ById) (*pb.EcoPointsBalance, error) {
	return s.ecoPoints.GetEcoPoints(req)
}

func (s *UserService) AddEcoPoints(req *pb.AddEcoPointsReq) (*pb.AddEcoPointsRes, error) {
	return s.ecoPoints.AddEcoPoints(req)
}

func (s *UserService) GetEcoPointsHistory(req *pb.PageLimit1) (*pb.EcoPointsHistoryRes, error) {
	return s.ecoPoints.GetEcoPointsHistory(req)
}

func (s *UserService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.Success, error) {
	err := s.db.InvalidateToken(req.AccessToken)
	if err != nil {
		return nil, err
	}
	return &pb.Success{Message: "Successfully logged out"}, nil
}
