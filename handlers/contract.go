package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func ContractHandler(logger *zap.Logger) http.HandlerFunc {
	return func(httpWriter http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		contractAddress := vars["contractAddress"]
		logger.Info("Calling contract service " + contractAddress + " with paratmeter " + vars["contractAddress"])
		httpWriter.Write([]byte(contractAddress))
	}

}
