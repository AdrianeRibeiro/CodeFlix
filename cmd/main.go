package main

import (
	"os"

	"github.com/AdrianeRibeiro/CodePix/application/grpc"
	"github.com/AdrianeRibeiro/CodePix/infrastructure/db"
	"gorm.io/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDb(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
