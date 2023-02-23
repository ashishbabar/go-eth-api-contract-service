package contracts

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

type Contract struct {
	ID       string `json:"id"`
	ByteCode string `json:"byteCode" validate:"required"`
	ABI      string `json:"ABI" validate:"required"`
}

func NewContract(r *http.Request) (Contract, error) {
	var contract Contract
	err := json.NewDecoder(r.Body).Decode(&contract)
	if err != nil {
		return Contract{}, err
	}
	return contract, nil
}

func (contract Contract) Validate() error {
	err := validate.Struct(contract)
	if err != nil {
		return err
	}
	return nil
}
