package service

import (
	pb "auth-service/genprotos/auth_pb"
	"context"
)

// 8
func (s *AuthServiceSt) CreateKitchen(ctx context.Context, in *pb.CreateKitchenRequest) (*pb.CreateKitchenResponse, error) {
	s.logger.Info("create kitchen request")
	return s.service.CreateKitchen(ctx, in)
}

// 9
func (s *AuthServiceSt) UpdateKitchen(ctx context.Context, in *pb.UpdateKitchenRequest) (*pb.UpdateKitchenResponse, error) {
	s.logger.Info("update kitchen request")
	return s.service.UpdateKitchen(ctx, in)
}

// 10
func (s *AuthServiceSt) GetKitchen(ctx context.Context, in *pb.GetKitchenRequest) (*pb.GetKitchenResponse, error) {
	s.logger.Info("get kitchen request")
	return s.service.GetKitchen(ctx, in)
}

// 11
func (s *AuthServiceSt) ListKitchens(ctx context.Context, in *pb.ListKitchensRequest) (*pb.ListKitchensResponse, error) {
	s.logger.Info("list kitchens request")
	return s.service.ListKitchens(ctx, in)
}

// 12
func (s *AuthServiceSt) SearchKitchens(ctx context.Context, in *pb.SearchKitchensRequest) (*pb.SearchKitchensResponse, error) {
	s.logger.Info("search kitchens request")
	return s.service.SearchKitchens(ctx, in)
}