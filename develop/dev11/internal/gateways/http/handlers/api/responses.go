package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

// Ответ удачный
type SuccessResponse struct {
	Result interface{} `json:"result"`
}

// Ответ с ошибкой
type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteJSONResponse(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response.Body)
}
