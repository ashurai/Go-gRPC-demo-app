package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/ashurai/fap-back/farmer-service/proto/farmer"
	productpb "github.com/ashurai/fap-back/product-service/proto/product"
	micro "github.com/micro/go-micro"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(*pb.Farmer) (*pb.Farmer, error)
	GetAll() []*pb.Farmer
}

type Repository struct {
	farmer  []*pb.Farmer
	product productpb.ProductServiceClient
	params  productpb.QueryParams
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
	repo          IRepository
	productClient productpb.ProductServiceClient
}

func (s *service) CreateFarmer(ctx context.Context, req *pb.Farmer, res *pb.Response) error {

	productResponse, err := s.productClient.FindFarmerProduct(context.Background(), &productpb.QueryParams{
		FarmerId: req.Id,
		Quantity: req.Quantity,
	})

	log.Printf("Found available product: %v", productResponse.Product.Name)

	if err != nil {
		return err
	}
	req.Id = productResponse.Product.Id
	farmer, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Farmer = farmer
	return nil
}

func (s *service) GetFarmer(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	farmer := s.repo.GetAll()

	res.Farmers = farmer
	return nil
}

func main() {
	repo := &Repository{}

	srvr := micro.NewService(
		micro.Name("go.micro.srv.farmer"),
		micro.Version("latest"),
	)

	productClient := productpb.NewProductServiceClient("go.micro.srv.product", srvr.Client())
	srvr.Init()

	data, err := productClient.GetMSG(context.Background(), &productpb.QueryParams{
		FarmerId: "abc123",
		Quantity: 3,
	})
	if err != nil {
		fmt.Printf("get sms error: %v", err)
	}
	fmt.Printf("data: %v", data)
	pb.RegisterFarmerServiceHandler(srvr.Server(), &service{repo, productClient})
	// Run the server
	if err := srvr.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}
