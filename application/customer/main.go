package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	switch gin.Mode() {
	case gin.ReleaseMode:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true})

	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false})
	}
}

func main() {

	if s, err := di(8081); err != nil {
		log.Fatal().Any("failed to inject dependency", err)
		fmt.Println("error", err)
	} else {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal().Any("failed to inject dependency", err)
		}
	}
}
