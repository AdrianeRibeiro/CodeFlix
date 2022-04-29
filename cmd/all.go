package cmd

import (
	"fmt"
	"os"

	"github.com/AdrianeRibeiro/CodePix/application/grpc"
	"github.com/AdrianeRibeiro/CodePix/application/kafka"
	"github.com/AdrianeRibeiro/CodePix/infrastructure/db"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
)

var gRPCPortNumber int

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run gRPC and kafka consumer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("all called")
		database := db.ConnectDB(os.Getenv("env"))
		go grpc.StartGrpcServer(database, portNumber)

		deliveryChan := make(chan ckafka.Event)
		producer := kafka.NewKafkaProducer()
		go kafka.DeliveryReport(deliveryChan)

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
	allCmd.Flags().IntVarP(&gRPCPortNumber, "grpc-port", "p", 500051, "gRPC Port")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
