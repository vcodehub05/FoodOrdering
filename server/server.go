package server

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine, port int) *http.Server {
	if os.Getenv("GIN_MODE") == "release" {
		port = 8080
	}
	return &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%d", port),
	}
}

type Response struct {
	State  string      `json:"state" enums:"success,failure"`
	Data   interface{} `json:"data,omitempty"`
}
