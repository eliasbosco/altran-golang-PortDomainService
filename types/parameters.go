package types

import "os"

//Config all configurations variable used to connect to services
type Config struct {
	GrpcAddress string
	SqlitePath  string
}

//SetupConfig Using environment variables to setup all the configuration's variables
func SetupConfig() Config {
	config := Config{}
	if os.Getenv("GRPC_ADDRESS") == "" {
		panic("No gRPC server address informed.")
	} else {
		config.GrpcAddress = os.Getenv("GRPC_ADDRESS")
	}

	if os.Getenv("SQLITE_PATH") == "" {
		config.SqlitePath = "/tmp/ports.db"
	} else {
		config.SqlitePath = os.Getenv("SQLITE_PATH")
	}

	return config
}
