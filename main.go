package main

import (
	"MyProjects/RentCar_gRPC/rental_service/clients"
	"MyProjects/RentCar_gRPC/rental_service/config"
	"MyProjects/RentCar_gRPC/rental_service/protogen/rental"
	rentalService "MyProjects/RentCar_gRPC/rental_service/services/rental"
	
	"MyProjects/RentCar_gRPC/rental_service/storage"
	"MyProjects/RentCar_gRPC/rental_service/storage/postgres"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// * @license.name  Apache 2.0
// * @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

	cfg := config.Load()
	psqlAUTH := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	var interStg storage.StorageInter
	interStg, err := postgres.InitDB(psqlAUTH)
	if err != nil {
		panic(err)
	}

	log.Printf("\ngRPC server running port%s with tcp protocol!\n", cfg.GRPCPort)

	listener, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		panic(err)
	}

	grpcClients, err := clients.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	defer grpcClients.Close()

	c := &rentalService.RentalService{
		Stg: interStg,
	}

	newS := grpc.NewServer()

	rental.RegisterRentalServiceServer(newS, c)

	reflection.Register(newS)

	if err := newS.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
