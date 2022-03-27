package main

import (
	"log"

	"github.com/ViniciusMartinss/phone-number-handler/src/handler"
	"github.com/ViniciusMartinss/phone-number-handler/src/infrastructure"
	"github.com/ViniciusMartinss/phone-number-handler/src/infrastructure/api"
	"github.com/ViniciusMartinss/phone-number-handler/src/infrastructure/repository"
	"github.com/ViniciusMartinss/phone-number-handler/src/usecase"
)

func main() {
	connection, err := infrastructure.InitializeDatabase()
	if err != nil {
		panic(err)
	}

	phoneRepository := repository.NewPhoneRepository(connection)
	phoneUsecase := usecase.NewphoneUsecase(phoneRepository)
	phoneHandler := handler.NewPhoneHandler(phoneUsecase)

	server := api.NewServer(phoneHandler).
		Set()

	go server.Start()

	port := infrastructure.GetConfig("api.port")
	log.Printf("[INFO] Server Initialized Successfully. Running on: %s\n", port)

	server.Stop()
}
