package tr

import (
	"context"
	pb "homeworks/homework_47/genproto/transport"
)

func (t *TransportRepo) GetBusSchedule(ctx context.Context, in *pb.BusNumber) (*pb.BusSchudle, error) {
	location := []*pb.Location{
		{
			Location: "any where",
			Time:     "every 27 minutes",
		},
	}
	return &pb.BusSchudle{BusNumber: in.Bus, Stations: location}, nil
}
