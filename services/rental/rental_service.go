package rental

import (
	"MyProjects/RentCar_gRPC/rental_service/clients"
	"MyProjects/RentCar_gRPC/rental_service/protogen/rental"
	"MyProjects/RentCar_gRPC/rental_service/storage"

	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RentalService struct {
	Stg storage.StorageInter
	rental.UnimplementedRentalServiceServer
	grpcClients *clients.GrpcClients
}

// ?===============================================================================================================
func (rs *RentalService) CreateRental(ctx context.Context, req *rental.CreateRentalRequest) (*rental.Rental, error) {
	id := uuid.New()
	err := rs.Stg.AddNewRental(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "rs.Stg.CreateRental: %s", err.Error())
	}

	res, err := rs.Stg.GetRentalById(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "rs.Stg.GetBrandById: %s", err.Error())
	}

	return &rental.Rental{
		RentalId:   res.RentalId,
		CarId:      res.Car.CarId,
		CustomerId: res.Customer.Id,
		StartDate:  res.StartDate,
		EndDate:    res.EndDate,
		Payment:    res.Payment,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
	}, nil
}

// ?===============================================================================================================
func (rs *RentalService) GetRentalByID(ctx context.Context, req *rental.GetRentalByIDRequest) (*rental.GetRentalByIDResponse, error) {
	res, err := rs.Stg.GetRentalById(req.RentalId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "rs.Stg.GetRentalById: %s", err.Error())
	}
	return res, nil
}

// ?===============================================================================================================
func (rs *RentalService) GetRentalList(ctx context.Context, req *rental.GetRentalListRequest) (*rental.GetRentalListResponse, error) {
	res, err := rs.Stg.GetRentalList(int(req.Limit), int(req.Offset), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "rs.Stg.GetRentalList: %s", err.Error())
	}
	return res, nil
}

// ?===============================================================================================================
func (rs *RentalService) UpdateRental(ctx context.Context, req *rental.UpdateRentalRequest) (*rental.Rental, error) {
	err := rs.Stg.UpdateRental(req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "rs.Stg.UpdateRental: %s", err.Error())
	}

	res, err := rs.Stg.GetRentalById(req.RentalId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "rs.Stg.UpdateRental: %s", err.Error())
	}

	return &rental.Rental{
		RentalId:   res.RentalId,
		CarId:      res.Car.CarId,
		CustomerId: res.Customer.Id,
		StartDate:  res.StartDate,
		EndDate:    res.EndDate,
		Payment:    res.Payment,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
	}, nil
}

// ?===============================================================================================================
func (rs *RentalService) DeleteRental(ctx context.Context, req *rental.DeleteRentalRequest) (*rental.Rental, error) {
	res, err := rs.Stg.GetRentalById(req.RentalId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "rs.Stg.GetRentalById: %s", err.Error())
	}

	err = rs.Stg.DeleteRental(req.RentalId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "rs.Stg.DeleteRental: %s", err.Error())
	}

	return &rental.Rental{
		RentalId:   res.RentalId,
		CarId:      res.Car.CarId,
		CustomerId: res.Customer.Id,
		StartDate:  res.StartDate,
		EndDate:    res.EndDate,
		Payment:    res.Payment,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
	}, nil
}
