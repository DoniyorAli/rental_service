package storage

import (
	"MyProjects/RentCar_gRPC/rental_service/protogen/rental"
)

type StorageInter interface {
	//* Rental
	AddNewRental(id string, req *rental.CreateRentalRequest) error
	GetRentalById(id string) (*rental.GetRentalByIDResponse, error)
	GetRentalList(offset, limit int, search string) (*rental.GetRentalListResponse, error)
	UpdateRental(entity *rental.UpdateRentalRequest) error
	DeleteRental(id string) error

}
