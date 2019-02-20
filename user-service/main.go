package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/ashurai/fap-back/user-service/proto/user"
	micro "github.com/micro/go-micro"
)

func main() {
	// Create connection with the DB
	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	fmt.Println("host=>", os.Getenv("DB_HOST"))

	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	// Add user service to micro config list
	// so that it will be accessible in by name / id
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	// This will start the server by parsing cli flag
	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
