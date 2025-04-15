package interfaces

import (
	"net/http"

	"github.com/microservice-monorepo/gateway-service/internal/gateway/application"
	"github.com/microservice-monorepo/gateway-service/internal/gateway/infrastructure/utils"
	"github.com/microservice-monorepo/gateway-service/internal/gateway/interfaces/dto"
)

type SignupHandler struct {
	Signup application.Registrar
	JSON   utils.JSONHelper
}

func NewSignupHandler(s application.Registrar) *SignupHandler {
	return &SignupHandler{
		Signup: s,
		JSON:   &utils.JSON{},
	}
}

func (uc SignupHandler) Handle(w http.ResponseWriter, r *http.Request) {

	payload := &dto.Payload{}

	err := uc.JSON.Read(w, r, payload)
	if err != nil {
		uc.JSON.Error(w, http.StatusUnprocessableEntity, dto.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "unprocessable entity",
			Data:       nil,
			Error:      true,
		})
		return
	}

	// signup use case
	data, err := uc.Signup.Execute(r.Context(), payload)
	if err != nil {
		uc.JSON.Error(w, http.StatusBadRequest, dto.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "bad request",
			Data: []map[string]any{
				{"result": data},
			},
			Error: true,
		})
		return
	}

	_ = uc.JSON.Write(w, http.StatusOK, dto.Response{
		StatusCode: http.StatusOK,
		Message:    "user created",
		Data: []map[string]any{
			{"result": data},
		},
		Error: false,
	})
}
