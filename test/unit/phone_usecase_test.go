package unit

import (
	"testing"

	"github.com/ViniciusMartinss/phone-number-handler/src/domain"
	"github.com/ViniciusMartinss/phone-number-handler/src/usecase"
	"github.com/ViniciusMartinss/phone-number-handler/test/mock/infrastructure"
	. "github.com/onsi/gomega"
)

func TestPhoneUsecase(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Expect to list all phone numbers classified", func(it *testing.T) {
		filtersMock := make(map[string]string, 0)

		phoneRepositoryMock := infrastructure.NewPhoneRepositoryMock(false)
		phoneUsecase := usecase.NewphoneUsecase(phoneRepositoryMock)

		result := phoneUsecase.List(filtersMock)
		g.Expect(result).ToNot(BeNil())

		g.Expect(result).ToNot(BeNil())
		g.Expect(result[0].Country).To(Equal("Morocco"))
		g.Expect(result[0].State).To(Equal("NOK"))
		g.Expect(result[0].CountryCode).To(Equal("+212"))
		g.Expect(result[0].PhoneNumber).To(Equal("6007989253"))
	})

	t.Run("Expect to list all phone numbers classified as OK", func(it *testing.T) {
		filtersMock := map[string]string{"state": "OK"}

		phoneRepositoryMock := infrastructure.NewPhoneRepositoryMock(false)
		phoneUsecase := usecase.NewphoneUsecase(phoneRepositoryMock)

		result := phoneUsecase.List(filtersMock)

		g.Expect(result[0].Country).To(Equal("Morocco"))
		g.Expect(result[0].State).To(Equal("OK"))
		g.Expect(result[0].CountryCode).To(Equal("+212"))
		g.Expect(result[0].PhoneNumber).To(Equal("698054317"))
	})

	t.Run("Expect to receive empty slice of PhoneReturnee when there is an database error", func(it *testing.T) {
		filtersMock := make(map[string]string, 0)

		phoneRepositoryMock := infrastructure.NewPhoneRepositoryMock(true)
		phoneUsecase := usecase.NewphoneUsecase(phoneRepositoryMock)

		result := phoneUsecase.List(filtersMock)

		g.Expect(result).To(Equal([]domain.PhoneReturnee{}))
	})
}
