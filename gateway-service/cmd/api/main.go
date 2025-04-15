package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/microservice-monorepo/gateway-service/internal/gateway/application"
	"github.com/microservice-monorepo/gateway-service/internal/gateway/infrastructure/events"
	"github.com/microservice-monorepo/gateway-service/internal/gateway/interfaces"
	"github.com/nats-io/nats.go"
)

func main() {

	// load environment vars from file
	if err := godotenv.Load("./.env"); err != nil {
		// if in docker container
		if err := godotenv.Load("./app/.env"); err != nil {
			log.Panic(err)
		}
	}

	// nats pub/sub service
	nc, err := nats.Connect(os.Getenv("NATS"), nats.UserCredentials(os.Getenv("NATS_CREDS")))
	if err != nil {
		log.Println(err) // TODO: change to panic
	}

	// event publisher singleton
	eventPublisher := events.NewEventPublisher(nc)
	// use cases
	signupUC := application.NewSignupUseCase(eventPublisher)

	app := &system{
		Config: &config{
			Addr: fmt.Sprintf(":%s", os.Getenv("PORT")),
		},
		Handler: &handlers{
			Signup: *interfaces.NewSignupHandler(signupUC),
		},
	}

	log.Printf("| GATEWAY => [@%s] is running ...\n", app.Config.Addr)
	// mount handler, and start server
	if err := app.run(app.mount()); err != nil {
		log.Panic(err)
	}
}
