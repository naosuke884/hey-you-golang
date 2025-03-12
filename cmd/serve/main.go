package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	server "github.com/naosuke884/hey-you"
	"github.com/naosuke884/hey-you/handler"
	"github.com/rs/zerolog/log"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.HomeHandler)
	http.Handle("/", router)

	srv := &http.Server{
		Handler:      router,
		Addr:         server.Envs.Host + ":" + server.Envs.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Info().Str("address", server.Envs.Host+":"+server.Envs.Port).Msg("server start")
	log.Fatal().Err(srv.ListenAndServe()).Msg("server stop")
}
