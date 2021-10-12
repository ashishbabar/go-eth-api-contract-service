package main

import (
	"log"
	"net/http"

	"github.com/ashishbabar/go-eth-api-contract-service/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("port", "3000")
}
func main() {
	router := mux.NewRouter()
	port := viper.GetString("port")
	router.HandleFunc("/{contractAddress}", handlers.ContractHandler(ZapLogger))
	http.Handle("/", router)
	ZapLogger.Info("Starting routes API at " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
