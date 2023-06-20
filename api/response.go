package api

import (
	"encoding/json"
	"net/http"
)

type Reader interface {
	Records() [][]string
}

type Response struct {
	Status struct {
		Success bool   `json:"success"`
		Message string `json:"message,omitempty"`
		Error   string `json:"error,omitempty"`
	} `json:"status"`
}

func NewResponse(success bool, message string, err error) Response {
	response := Response{}
	response.Status.Success = success
	response.Status.Message = message
	if err != nil {
		response.Status.Error = err.Error()
	}
	return response
}

func Write(w http.ResponseWriter, status int, resp interface{}) {
	jsonResponse(w, status, resp)
}

func jsonResponse(w http.ResponseWriter, status int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(status)
	_, _ = w.Write(response)
}
