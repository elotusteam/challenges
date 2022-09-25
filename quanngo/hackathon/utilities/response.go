package utilities

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(w http.ResponseWriter, r *http.Request, message string, data interface{}) {
	res := Response{
		Message: message,
		Errors:  nil,
		Data:    data,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "n")
	res := Response{
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}
