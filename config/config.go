package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	App         string
	AppVersion  string
	Environment string // devlopment, staging, production
	
	GRPCPort    string

	DefaultOffset string
	DefaultLimit  string

	CarServiceGrpcHost string
	CarServiceGrpcPort string

	RentalServiceGrpcHost string
	RentalServiceGrpcPort string

	AuthorizationServiceGrpcHost string
	AuthorizationServiceGrpcPort string

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("APP", "RentCarApp"))
	config.AppVersion = cast.ToString(getOrReturnDefaultValue("APP_VERSION", "1.0.0"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "development"))

	config.GRPCPort = cast.ToString(getOrReturnDefaultValue("GRPC_PORT", ":5000"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	config.CarServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("CAR_SERVICE_GRPC_HOST", "localhost"))
	config.CarServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("CAR_SERVICE_GRPC_PORT", ":7000"))

	config.RentalServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("RENTAL_SERVICE_GRPC_HOST", "localhost"))
	config.RentalServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("RENTAL_SERVICE_GRPC_PORT", ":7002"))

	config.AuthorizationServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("AUTHORIZATION_SERVICE_GRPC_HOST", "localhost"))
	config.AuthorizationServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("AUTHORIZATION_SERVICE_GRPC_PORT", ":7001"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "127.0.0.1"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "your_service_db"))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "you_are_user"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "your_db_pswd"))
	
	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
