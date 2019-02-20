package main

import (
	"context"
	"log"
	"os"

	pb "github.com/ashurai/fap-back/user-service/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

func main() {
	cmd.Init()

	// Create new user greeter client
	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	// Defining flags to ask / get queries from cli
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "your full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "your email address",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "choose your password",
			},
			cli.StringFlag{
				Name:  "Language",
				Usage: "choose your language",
			},
		),
	)

	// Initiate service
	service.Init(
		micro.Action(func(c *cli.Context) {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			language := c.String("language")

			// Call user service to process data
			r, err := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
				Language: language,
			})

			if err != nil {
				log.Fatalf("Was unable to create user: %v", err)
			}

			log.Printf("User Created: %v", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("was not unable to list users: %v", err)
			}
			for _, v := range getAll.Users {
				log.Println(v)
			}

			os.Exit(0)
		}),
	)

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
