package integration

import (
	"os"
	"testing"

	"github.com/ViniciusMartinss/phone-number-handler/src/handler"
	"github.com/ViniciusMartinss/phone-number-handler/src/infrastructure"
	"github.com/ViniciusMartinss/phone-number-handler/src/infrastructure/repository"
	"github.com/ViniciusMartinss/phone-number-handler/src/usecase"
	. "github.com/onsi/gomega"
)

func TestPhoneHandler(t *testing.T) {
	os.Chdir("../..")

	g := NewGomegaWithT(t)

	connection, err := infrastructure.InitializeDatabase()
	if err != nil {
		panic(err)
	}

	t.Run("Expect to list all phone numbers classified", func(it *testing.T) {
		phoneRepository := repository.NewPhoneRepository(connection)
		phoneUsecase := usecase.NewphoneUsecase(phoneRepository)
		phoneHandler := handler.NewPhoneHandler(phoneUsecase)

		result := phoneHandler.List()

		g.Expect(result).ToNot(BeNil())
		g.Expect(len(result)).To(Equal(41))
		g.Expect(result[0].Country).To(Equal("Morocco"))
		g.Expect(result[0].State).To(Equal("NOK"))
		g.Expect(result[0].CountryCode).To(Equal("+212"))
		g.Expect(result[0].PhoneNumber).To(Equal("6007989253"))
	})

	t.Run("Expect to list all phones of Ethiopia", func(it *testing.T) {
		filtersMock := map[string]string{"country": "Ethiopia"}

		phoneRepository := repository.NewPhoneRepository(connection)
		phoneUsecase := usecase.NewphoneUsecase(phoneRepository)
		phoneHandler := handler.NewPhoneHandler(phoneUsecase)

		result := phoneHandler.List(filtersMock)
		g.Expect(len(result)).To(Equal(9))
		g.Expect(result[0].Country).To(Equal("Ethiopia"))
		g.Expect(result[0].State).To(Equal("NOK"))
		g.Expect(result[0].CountryCode).To(Equal("+251"))
		g.Expect(result[0].PhoneNumber).To(Equal("9773199405"))
	})

	t.Run("Expect to list all OK phones", func(it *testing.T) {
		filtersMock := map[string]string{"state": "OK"}

		phoneRepository := repository.NewPhoneRepository(connection)
		phoneUsecase := usecase.NewphoneUsecase(phoneRepository)
		phoneHandler := handler.NewPhoneHandler(phoneUsecase)

		result := phoneHandler.List(filtersMock)
		g.Expect(len(result)).To(Equal(27))
		g.Expect(result[0].Country).To(Equal("Morocco"))
		g.Expect(result[0].State).To(Equal("OK"))
		g.Expect(result[0].CountryCode).To(Equal("+212"))
		g.Expect(result[0].PhoneNumber).To(Equal("698054317"))
	})
}
