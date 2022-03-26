package unit

import (
	"testing"

	"github.com/ViniciusMartinss/phone-number-handler/src/domain"
	"github.com/ViniciusMartinss/phone-number-handler/src/handler"
	"github.com/ViniciusMartinss/phone-number-handler/test/mock/application"
	"github.com/ViniciusMartinss/phone-number-handler/test/mock/infrastructure"
	. "github.com/onsi/gomega"
)

func TestPhoneHandler(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Expect to list all countries", func(it *testing.T) {
		filtersMock := make(map[string]string, 0)

		phoneRepositoryMock := infrastructure.NewPhoneRepositoryMock(false)
		phoneUseCaseMock := application.NewphoneUsecaseMock(phoneRepositoryMock)
		phoneHandler := handler.NewPhoneHandler(phoneUseCaseMock)

		result := phoneHandler.List(filtersMock)

		g.Expect(result).ToNot(BeNil())
		g.Expect(result[0].Country).To(Equal("Morocco"))
		g.Expect(result[0].State).To(Equal("OK"))
		g.Expect(result[0].CountryCode).To(Equal("+212"))
		g.Expect(result[0].PhoneNumber).To(Equal("6007989253"))
	})

	t.Run("Expect to receive empty slice of PhoneReturnee with invalid country on filter", func(it *testing.T) {
		filtersMock := map[string]string{"country": "Portugal"}

		phoneRepositoryMock := infrastructure.NewPhoneRepositoryMock(false)
		phoneUseCaseMock := application.NewphoneUsecaseMock(phoneRepositoryMock)
		phoneHandler := handler.NewPhoneHandler(phoneUseCaseMock)

		result := phoneHandler.List(filtersMock)

		g.Expect(result).To(Equal([]domain.PhoneReturnee{}))
	})
}
