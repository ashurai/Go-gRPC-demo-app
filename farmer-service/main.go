package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ashurai/fap-back/farmer-service/proto/farmer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(*pb.Farmer) (*pb.Farmer, error)
	GetAll() []*pb.Farmer
}

type Repository struct {
	farmer []*pb.Farmer
}

func (repo *Repository) Create(farmer *pb.Farmer) (*pb.Farmer, error) {
	updated := append(repo.farmer, farmer)
	repo.farmer = updated
	return farmer, nil
}

func (repo *Repository) GetAll() []*pb.Farmer {
	return repo.farmer
}

type service struct {
	repo IRepository
}

func (s *service) CreateFarmer(ctx context.Context, req *pb.Farmer) (*pb.Response, error) {
	farmer, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Farmer: farmer}, nil
}

func (s *service) GetFarmer(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	farmer := s.repo.GetAll()
	return &pb.Response{Farmers: farmer}, nil
}

func main() {
	repo := &Repository{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFarmerServiceServer(s, &service{repo})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}
