package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ViniciusMartinss/phone-number-handler/src/domain"
	"github.com/ViniciusMartinss/phone-number-handler/src/infrastructure"
	"github.com/gorilla/mux"
)

type Server interface {
	Set() server
	Start()
	Stop()
}

type server struct {
	phoneHandler domain.PhoneHandler
	server       *http.Server
}

func NewServer(phoneHandler domain.PhoneHandler) Server {
	return server{phoneHandler, nil}
}

func (s server) Set() server {
	port := infrastructure.GetConfig("api.port")

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: s.defineRoutes(),
	}

	return s
}

func (s server) defineRoutes() *mux.Router {
	router := mux.NewRouter().
		StrictSlash(true)

	router.Use(SetHeaders)

	router.HandleFunc("/phone", s.phoneListHandler).
		Methods("GET")

	return router
}

func (s server) Start() {
	s.server.
		ListenAndServe()
}

func (s server) Stop() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		os.Kill,
		syscall.SIGTERM,
	)

	<-ctx.Done()
	log.Println("[WARN] Server is shutting down!")

	err := s.server.
		Shutdown(context.Background())

	if err != nil {
		log.Fatalf("[FAIL] Fail on stop the server, reason: %s\n", err.Error())
	}

	stop()
}
