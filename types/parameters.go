package types

import "os"

//Config all configurations variable used to connect to services
type Config struct {
	GrpcAddress   string
	MysqlAddr     string
	MysqlUsername string
	MysqlPassword string
}

//SetupConfig Using environment variables to setup all the configuration's variables
func SetupConfig() Config {
	config := Config{}
	if os.Getenv("GRPC_ADDRESS") == "" {
		panic("No gRPC server address informed.")
	} else {
		config.GrpcAddress = os.Getenv("GRPC_ADDRESS")
	}

	if os.Getenv("MYSQL_ADDR") == "" {
		config.MysqlAddr = "172.17.0.1:3306"
	} else {
		config.MysqlAddr = os.Getenv("MYSQL_ADDR")
	}

	if os.Getenv("MYSQL_USERNAME") == "" {
		panic("No mysql username informed.")
	} else {
		config.MysqlUsername = os.Getenv("MYSQL_USERNAME")
	}

	if os.Getenv("MYSQL_PASSWORD") == "" {
		panic("No mysql password informed.")
	} else {
		config.MysqlPassword = os.Getenv("MYSQL_PASSWORD")
	}

	return config
}
