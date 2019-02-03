package main

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/ashurai/fap-back/product-service/proto/product"
	micro "github.com/micro/go-micro"
)

type Repository interface {
	FindFarmerProduct(*pb.QueryParams) (*pb.Product, error)
}

type ProductRepository struct {
	products []*pb.Product
}

func (repo ProductRepository) FindFarmerProduct(params *pb.QueryParams) (*pb.Product, error) {
	for _, product := range repo.products {
		if params.Quantity <= product.Available {
			return product, nil
		}
	}

	return nil, errors.New("Not Available product")
}

type service struct {
	repo Repository
}

func (s *service) FindFarmerProduct(ctx context.Context, req *pb.QueryParams, res *pb.Response) error {
	product, err := s.repo.FindFarmerProduct(req)
	if err != nil {
		return err
	}

	res.Product = product
	return nil
}

func main() {
	products := []*pb.Product{
		&pb.Product{Id: "abc123", Name: "Tomoto", Available: 6, Quantity: 70, FarmerId: "ddfcb"},
	}

	repo := &ProductRepository{products}

	srvr := micro.NewService(
		micro.Name("go.micro.srv.prodct"),
		micro.Version("latest"),
	)

	srvr.Init()

	pb.RegisterProductServiceHandler(srvr.Server(), &service{repo})
	if err := srvr.Run(); err != nil {
		fmt.Println(err)
	}
}
