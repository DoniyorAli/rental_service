migrate create -ext sql -dir ./storage/migrations -seq -digits 2 create_rental_table

migrate -path ./storage/migrations -database 'postgres://admin:pswd123@localhost:9876/rental_service_db?sslmode=disable' up

migrate -path ./storage/migrations -database 'postgres://admin:pswd123@localhost:9876/rental_service_db?sslmode=disable' down