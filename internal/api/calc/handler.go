package calc

import (
	"errors"
	"net/http"
	"api_calculate/internal/api"
)

type CalculateReq struct {
	FirstNumber   float64    `json:"first_number"`
	SecondNumber  float64    `json:"second_number"`
	OperationName string `json:"operation_name"`
}

func (req CalculateReq) Validate(_ *api.Context) error {
	if req.OperationName == "division" && req.SecondNumber == 0 {
		return errors.New("error: division by zero")
	}

	return nil
}

type CalculateResp struct {
	Result float64 `json:"result"`
}

func (c *Calc) Handle(_ *api.Context, req *CalculateReq) (*CalculateResp, int) {
	resp := &CalculateResp{}

	switch req.OperationName {
	case "plus":
		resp.Result = req.FirstNumber + req.SecondNumber
	case "multiply":
		resp.Result = req.FirstNumber * req.SecondNumber
	case "minus":
		resp.Result = req.FirstNumber - req.SecondNumber
	case "division":
		resp.Result = req.FirstNumber / req.SecondNumber
	}

	return resp, http.StatusOK
}