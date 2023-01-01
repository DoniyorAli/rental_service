package postgres

import (
	"MyProjects/RentCar_gRPC/rental_service/protogen/rental"
	"errors"
	"time"
)

func (psql Postgres) AddNewRental(id string, req *rental.CreateRentalRequest) error {

	_, err := psql.homeDB.Exec(`
	INSERT INTO "rentals" 
	(
		"rental_id", 
		"car_id", 
		"customer_id", 
		"start_date", 
		"end_date", 
		"payment", 
		"created_at"
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		now()
	)`,
		id,
		req.CarId,
		req.CustomerId,
		req.StartDate,
		req.EndDate,
		req.Payment,
	)
	if err != nil {
		return err
	}
	return nil
}

// *=========================================================================
func (psql Postgres) GetRentalById(id string) (*rental.GetRentalByIDResponse, error) {
	res := &rental.GetRentalByIDResponse{
		Car:      &rental.GetRentalByIDResponse_Car{},
		Customer: &rental.GetRentalByIDResponse_User{},
	}

	var deletedAt *time.Time
	var updatedAt *string

	err := psql.homeDB.QueryRow(`SELECT 
		"rental_id", 
		"car_id", 
		"customer_id", 
		"start_date", 
		"end_date",
		"payment", 
		"created_at", 
		"updated_at"
    FROM "rentals" WHERE "rental_id" = $1`, id).Scan(
		&res.RentalId,
		&res.CarId,
		&res.CustomerId,
		&res.StartDate,
		&res.EndDate,
		&res.Payment,
		&res.CreatedAt,
		&updatedAt,
	)
	if err != nil {
		return res, err
	}

	if updatedAt != nil {
		res.UpdatedAt = *updatedAt
	}

	if deletedAt != nil {
		return res, errors.New("rental not fount")
	}

	return res, nil
}

// *=========================================================================
func (psql Postgres) GetRentalList(limit, offset int, search string) (*rental.GetRentalListResponse, error) {
	resp := &rental.GetRentalListResponse{
		Rentals: make([]*rental.Rental, 0),
	}

	rows, err := psql.homeDB.Queryx(`SELECT 
		"rental_id", 
		"car_id", 
		"customer_id", 
		"start_date", 
		"end_date", 
		"payment", 
		"created_at", 
		"updated_at"
	FROM "rentals" WHERE deleted_at IS NULL AND (start_date ILIKE '%' || $1 || '%')
		LIMIT $2
		OFFSET $3
	`,
		search,
		limit,
		offset,
	)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var r = &rental.Rental{}
		var updatedAt *string

		err := rows.Scan(
			&r.RentalId,
			&r.CarId,
			&r.CustomerId,
			&r.StartDate,
			&r.EndDate,
			&r.Payment,
			&r.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return resp, err
		}

		if updatedAt != nil {
			r.UpdatedAt = *updatedAt
		}
		resp.Rentals = append(resp.Rentals, r)

	}
	return resp, err
}

// *=========================================================================
func (psql Postgres) UpdateRental(id string, box *rental.UpdateRentalRequest) error {

	res, err := psql.homeDB.NamedExec(`
	UPDATE "rentals"  
		SET "car_id"=:ca, 
			"customer_id"=:cu, 
			"start_date"=:s, 
			"end_date"=:e, 
			"payment"=:p, 
			updated_at=now()
		WHERE deleted_at IS NULL AND rental_id=:id`, map[string]interface{}{
		"id": box.RentalId,
		"ca": box.CarId,
		"cu": box.CustomerId,
		"s":  box.StartDate,
		"e":  box.EndDate,
		"p":  box.Payment,
	})
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect > 0 {
		return nil
	}
	return errors.New("rental not found")
}

// *=========================================================================
func (psql Postgres) DeleteRental(id string) error {
	res, err := psql.homeDB.Exec(`
	UPDATE "rentals" 
		SET deleted_at=now() 
			WHERE rental_id=$1 AND deleted_at IS NULL`, id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect > 0 {
		return nil
	}
	return errors.New("rental not found")
}
