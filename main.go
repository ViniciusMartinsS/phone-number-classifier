package main

import (
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
	server.Stop()
}
