package infrastructure

import (
	"fmt"

	"github.com/ViniciusMartinss/phone-number-handler/src/domain"
)

type phoneRepositoryMock struct {
	fail bool
}

func NewPhoneRepositoryMock(fail bool) domain.PhoneRepository {
	return phoneRepositoryMock{fail: fail}
}

func (p phoneRepositoryMock) Find(countryCode string) ([]string, error) {
	if p.fail {
		return nil, fmt.Errorf("an error has happened")
	}

	return []string{"(212) 6007989253", "(212) 698054317"}, nil
}
