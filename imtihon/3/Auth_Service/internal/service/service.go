package service

import (
	"log/slog"

	pb "auth-service/genprotos/auth_pb"
	"auth-service/internal/storage"
)

type (
	AuthServiceSt struct {
		pb.UnimplementedAuthServiceServer
		service storage.AuthSt
		logger  *slog.Logger
	}
)

func New(service storage.AuthSt, logger *slog.Logger) *AuthServiceSt {
	return &AuthServiceSt{
		service: service,
		logger:  logger,
	}
}

// fuel_management
// func (s *FuelServiceSt) CreateFuel(ctx context.Context, req *genprotos.CreateFuelRequest) (*genprotos.FuelResponse, error) {
// 	s.logger.Println("get fuel request")
// 	return s.service.CreateFuel(ctx, req)
// }

// func (s *FuelServiceSt) GetFuel(ctx context.Context, req *genprotos.GetFuelRequest) (*genprotos.FuelResponse, error) {
// 	s.logger.Println("get fuel request")
// 	return s.service.GetFuel(ctx, req)
// }

// func (s *FuelServiceSt) UpdateFuel(ctx context.Context, req *genprotos.UpdateFuelRequest) (*genprotos.FuelResponse, error) {
// 	s.logger.Println("update fuel request")
// 	return s.service.UpdateFuel(ctx, req)
// }

// func (s *FuelServiceSt) DeleteFuel(ctx context.Context, req *genprotos.DeleteFuelRequest) (*genprotos.Empty, error) {
// 	s.logger.Println("delete fuel request")
// 	return s.service.DeleteFuel(ctx, req)
// }

// func (s *FuelServiceSt) GetFuelByChoice(ctx context.Context, req *genprotos.GetFuelsByChoiceRequest) (*genprotos.GetFuelsByChoiceResponse, error) {
// 	s.logger.Println("Get fuel by name request")
// 	return s.service.GetFuelByChoice(ctx, req)
// }

// func (s *FuelServiceSt) GetFuels(ctx context.Context, req *genprotos.Empty) (*genprotos.GetFuelsResponse, error) {
// 	s.logger.Println("Get fuels request")
// 	return s.service.GetFuels(ctx, req)
// }

// // fuel_history

// func (s *FuelServiceSt) CreateFuelHistory(ctx context.Context, req *genprotos.CreateFuelHistoryRequest) (*genprotos.FuelHistoryResponse, error) {
// 	s.logger.Println("create fuel history request")
// 	return s.service.CreateFuelHistory(ctx, req)
// }

// func (s *FuelServiceSt) GetFuelHistoriesByID(ctx context.Context, req *genprotos.GetFuelHistoriesByIdRequest) (*genprotos.GetFuelHistoriesByIdResponse, error) {
// 	s.logger.Println("update fuel history request")
// 	return s.service.GetFuelHistoriesByID(ctx, req)
// }

// func (s *FuelServiceSt) GetFuelHistoriesByChoice(ctx context.Context, req *genprotos.GetFuelHistoriesByChoiceRequest) (*genprotos.GetFuelHistoriesByChoiceResponse, error) {
// 	s.logger.Println("Get fuel histories by name request")
// 	return s.service.GetFuelHistoriesByChoice(ctx, req)
// }

// func (s *FuelServiceSt) GetFuelHistoriesByDate(ctx context.Context, req *genprotos.GetFuelHistoriesByDateRequest) (*genprotos.GetFuelHistoriesByDateResponse, error) {
// 	s.logger.Println("Get fuel histories by date request")
// 	return s.service.GetFuelHistoriesByDate(ctx, req)
// }

// func (s *FuelServiceSt) GetFuelHistories(ctx context.Context, req *genprotos.Empty) (*genprotos.GetFuelHistoriesResponse, error) {
// 	s.logger.Println("Get fuel histories request")
// 	return s.service.GetFuelHistories(ctx, req)
// }
