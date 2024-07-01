package service

import (
	"context"
	pb "github.com/Azamjonov_Ozodjon/lesson46/generate/transport"
	"github.com/Azamjonov_Ozodjon/lesson46/storage"
)

type TranpostRepo struct {
	*pb.UnimplementedTranportServiceServer
	St storage.StorageRepo
}

func (t *TranpostRepo) GetBusSchedule(ctx context.Context, in pb.BusRequest) (*pb.BusScheduleRespons, error) {
	s, err := t.St.GetScheduleByNumber(int(in.BusNumber))
	return &pb.BusScheduleRespons{BusNumber: in.BusNumber, BusSchedule: s}, err
}

func (t *TranpostRepo) TrackBusLocation(ctx context.Context, in pb.BusRequest) (*pb.TrucLocation, error) {
	_, err := t.St.GetScheduleByNumber(int(in.BusNumber))
	return &pb.TrucLocation{Location: "every 25 minute = nowhere"}, err
}

func (t *TranpostRepo) ReportTraficJam(ctx context.Context, in pb.TrucLocation) (*pb.TraficJamRespons, error) {
	s, err := t.St.IsTraficJam(in.Location)
	return &pb.TraficJamRespons{IsTraficJam: s}, err
}
