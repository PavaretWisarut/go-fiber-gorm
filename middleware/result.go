package middleware

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Code    int         `json:"Code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  bool        `json:"Status"`
}

func Submit(data interface{}, msg string) ([]byte, error) {
	result := Result{
		Code:    200,
		Data:    data,
		Message: msg,
		Status:  true,
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %v", err)
	}

	return jsonData, nil
}
