package clients

import (
	"MyProjects/RentCar_gRPC/rental_service/config"
	"MyProjects/RentCar_gRPC/rental_service/protogen/authorization"
	"MyProjects/RentCar_gRPC/rental_service/protogen/car"
	"MyProjects/RentCar_gRPC/rental_service/protogen/rental"

	"google.golang.org/grpc"
)

type GrpcClients struct {
	Car    car.CarServiceClient
	Auth   authorization.AuthServiceClient
	Rental rental.RentalServiceClient
	conns  []*grpc.ClientConn
}

func NewGrpcClients(cfg config.Config) (*GrpcClients, error) {
	connectCar, err := grpc.Dial(cfg.CarServiceGrpcHost+cfg.CarServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	car := car.NewCarServiceClient(connectCar)

	connectAuth, err := grpc.Dial(cfg.AuthorizationServiceGrpcHost+cfg.AuthorizationServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	auth := authorization.NewAuthServiceClient(connectAuth)

	connRental, err := grpc.Dial(cfg.RentalServiceGrpcHost+cfg.RentalServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	rental := rental.NewRentalServiceClient(connRental)

	conns := make([]*grpc.ClientConn, 0)
	return &GrpcClients{
		Car:    car,
		Auth:   auth,
		Rental: rental,
		conns:  append(conns, connectCar, connectAuth, connRental),
	}, nil
}

func (c *GrpcClients) Close() {
	for _, v := range c.conns {
		v.Close()
	}
}
