package application

import "github.com/ViniciusMartinss/phone-number-handler/src/domain"

type phoneUsecaseMock struct {
	repository domain.PhoneRepository
}

func NewphoneUsecaseMock(repository domain.PhoneRepository) domain.PhoneUsecase {
	return phoneUsecaseMock{repository}
}

func (p phoneUsecaseMock) List(filters ...map[string]string) []domain.PhoneReturnee {
	phoneNumbers, _ := p.repository.Find("")

	phoneClassified := domain.PhoneReturnee{
		Country:     "Morocco",
		State:       "OK",
		CountryCode: "+212",
		PhoneNumber: phoneNumbers[0],
	}

	return []domain.PhoneReturnee{phoneClassified}
}
