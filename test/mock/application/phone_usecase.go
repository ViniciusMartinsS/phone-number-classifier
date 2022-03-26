package application

import (
	"strings"

	"github.com/ViniciusMartinss/phone-number-handler/src/domain"
)

type phoneUsecaseMock struct {
	repository domain.PhoneRepository
}

func NewphoneUsecaseMock(repository domain.PhoneRepository) domain.PhoneUsecase {
	return phoneUsecaseMock{repository}
}

func (p phoneUsecaseMock) List(filters ...map[string]string) []domain.PhoneReturnee {
	phoneNumbers, _ := p.repository.Find("")
	phoneNumber := strings.ReplaceAll(phoneNumbers[0], "(212)", "")

	phoneClassified := domain.PhoneReturnee{
		Country:     "Morocco",
		State:       "OK",
		CountryCode: "+212",
		PhoneNumber: strings.Trim(phoneNumber, " "),
	}

	return []domain.PhoneReturnee{phoneClassified}
}
