package service

import (
	"context"
	"log"

	pb "github.com/Azamjonov_Ozodjon/lesson46/genproto/generator/transport" // Import your generated transport proto package
	"github.com/Azamjonov_Ozodjon/lesson46/storage"
)

type TransportService struct {
	storage *storage.TransportStorage
}

func NewTransportService(db *storage.DB) *TransportService {
	return &TransportService{
		storage: storage.NewTransportStorage(db),
	}
}

func (s *TransportService) GetBusSchedule(ctx context.Context, req *pb.BusScheduleRequest) (*pb.BusScheduleResponse, error) {
	schedules, err := s.storage.GetBusSchedule(req.BusNumber)
	if err != nil {
		log.Printf("Failed to get bus schedule: %v", err)
		return nil, err
	}

	response := &pb.BusScheduleResponse{
		Schedules: schedules,
	}

	return response, nil
}

func (s *TransportService) TrackBusLocation(ctx context.Context, req *pb.BusLocationRequest) (*pb.BusLocationResponse, error) {
	location, err := s.storage.TrackBusLocation(req.BusNumber)
	if err != nil {
		log.Printf("Failed to track bus location: %v", err)
		return nil, err
	}

	response := &pb.BusLocationResponse{
		Location: location,
	}

	return response, nil
}

func (s *TransportService) ReportTrafficJam(ctx context.Context, req *pb.TrafficJamReportRequest) (*pb.TrafficJamReportResponse, error) {
	err := s.storage.ReportTrafficJam(req.Location, req.Description)
	if err != nil {
		log.Printf("Failed to report traffic jam: %v", err)
		return &pb.TrafficJamReportResponse{Success: false}, err
	}

	return &pb.TrafficJamReportResponse{Success: true}, nil
}
