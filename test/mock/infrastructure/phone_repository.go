package infrastructure

import "github.com/ViniciusMartinss/phone-number-handler/src/domain"

type phoneRepositoryMock struct {
	fail bool
}

func NewPhoneRepositoryMock(fail bool) domain.PhoneRepository {
	return phoneRepositoryMock{fail: fail}
}

func (p phoneRepositoryMock) Find(countryCode string) ([]string, error) {
	return []string{"(212) 6007989253"}, nil
}
