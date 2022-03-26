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
	log.Println("[INFO] - Initializing Database...")

	connection, err := infrastructure.InitializeDatabase()
	if err != nil {
		panic(err)
	}

	phoneRepository := repository.NewPhoneRepository(connection)
	phoneUsecase := usecase.NewphoneUsecase(phoneRepository)
	phoneHandler := handler.NewPhoneHandler(phoneUsecase)

	router := api.NewRouter(phoneHandler)
	router.DefineRoutes().
		Run()
}
