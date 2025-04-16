package application

import (
	"context"
	"log"

	"github.com/microservice-monorepo/gateway-service/internal/gateway/infrastructure/events"
	"github.com/microservice-monorepo/gateway-service/internal/gateway/infrastructure/utils"
	"github.com/microservice-monorepo/gateway-service/internal/gateway/interfaces/dto"
)

type SignupUseCase struct {
	Validator utils.Validator
	// Publisher is the singleton NATS instance for
	// calling the implemented methods
	//
	// for example => Publisher.Request(topic, message)
	//
	Publisher events.EventPublisher
}

func NewSignupUseCase(p events.EventPublisher) *SignupUseCase {
	return &SignupUseCase{
		Validator: utils.NewValidator(),
		Publisher: p,
	}
}

func (uc *SignupUseCase) Execute(ctx context.Context, p *dto.Payload) (map[string]string, error) {

	// validate payload
	err := uc.Validator.Struct(p)
	if err != nil {
		errMsg, err := uc.Validator.Test(err)
		if err != nil {
			return errMsg, err
		}
	}

	//publish to listener
	msg, err := uc.Publisher.Request("signup", p)
	if err != nil {
		return nil, err
	}

	// TODO: put response to return value
	log.Println(string(msg))

	return nil, nil
}
