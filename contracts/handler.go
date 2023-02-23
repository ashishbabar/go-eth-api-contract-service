package contracts

import (
	"net/http"

	"go.uber.org/zap"
)

type Handler interface {
	CreateContract(writer http.ResponseWriter, reader *http.Request)
}

type handler struct {
	service IService
	logger  *zap.Logger
}

func NewHandler(service IService, logger *zap.Logger) Handler {
	return handler{service: service, logger: logger}
}

func (h handler) CreateContract(writer http.ResponseWriter, request *http.Request) {
	// Parse request and create contract struct var
	h.logger.Info("Handling CreateContract call")

	err := request.ParseForm()
	if err != nil {
		h.logger.Error(err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	h.logger.Info("Parsed request without any errors")
	contract, err := NewContract(request)
	if err != nil {
		h.logger.Error(err.Error())
		writer.WriteHeader(http.StatusBadRequest)
	}
	h.logger.Info("Created new contract", zap.Any("Contract", contract))
	// err = contract.Validate()
	// if err != nil {
	// 	h.logger.Error(err.Error())
	// 	writer.WriteHeader(http.StatusBadRequest)
	// }
	h.logger.Info("Validated contract")

	err = h.service.Create(contract)
	if err != nil {
		h.logger.Error(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
	}
	h.logger.Info("Successfully processed contract creation")

}
