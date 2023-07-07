package middleware

import (
	"encoding/json"
	"fmt"
	"go-fiber-app/utils"
	"os"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  bool        `json:"Status"`
}

type Token struct {
	Code      int    `json:"code"`
	Token     string `json:"token"`
	Expire_at int    `json:"expire_at"`
	Message   string `json:"message"`
	Status    bool   `json:"Status"`
}

type Register struct {
	Code    int    `json:"code"`
	Id      int    `json:"id"`
	Message string `json:"message"`
	Status  bool   `json:"Status"`
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

func GenerateToken(token string, msg string) ([]byte, error) {
	secretkey := []byte(os.Getenv("JWT_SECRET_KEY"))
	decodeToken, err := utils.DecodeToken(token, secretkey)

	if err != nil {
		panic(err.Error())
	}

	exp := decodeToken["ExpireTime"].(float64)

	Expire_time := int(exp)

	result := Token{
		Code:      200,
		Token:     token,
		Expire_at: Expire_time,
		Message:   msg,
		Status:    true,
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %v", err)
	}

	return jsonData, nil
}

func Registers(id int, msg string) ([]byte, error) {
	
	result := Register{
		Code:    200,
		Id:      id,
		Message: msg,
		Status:  true,
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %v", err)
	}

	return jsonData, nil
}
