package types

import "os"

//Config all configurations variable used to connect to services
type Config struct {
	GrpcAddress string
	GrpcPort    string
	SQLitePath  string
}

//SetupConfig Using environment variables to setup all the configuration's variables
func SetupConfig() Config {
	config := Config{}
	if os.Getenv("GRPC_ADDRESS") == "" {
		panic("No gRPC server address informed.")
	} else {
		config.GrpcAddress = os.Getenv("GRPC_ADDRESS")
	}

	if os.Getenv("GRPC_PORT") == "" {
		panic("No gRPC server port address informed.")
	} else {
		config.GrpcPort = os.Getenv("GRPC_PORT")
	}

	if os.Getenv("SQLITE_PATH") == "" {
		config.SQLitePath = "/tmp/ports.db"
	} else {
		config.SQLitePath = os.Getenv("SQLITE_PATH")
	}

	return config
}
