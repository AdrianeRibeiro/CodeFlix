package main

import (
	"os"

	"github.com/codeedu/imersao/codepix-go/infrastructure/grpc"
	"gorm.io/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDb(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
