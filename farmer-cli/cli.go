package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/ashurai/fap-back/farmer-service/proto/farmer"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFileName = "farmer.json"
)

func parseFile(file string) (*pb.Farmer, error) {
	var farmer *pb.Farmer
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &farmer)
	return farmer, err
}

func main() {
	// Connecting to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewFarmerServiceClient(conn)

	file := defaultFileName
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	farmer, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}
	r, err := client.CreateFarmer(context.Background(), farmer)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetFarmer(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not find Farmer: %v", err)
	}

	for _, v := range getAll.Farmers {
		log.Println(v)
	}
}
