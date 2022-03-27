package usecase

import (
	"fmt"
	"regexp"

	"github.com/ViniciusMartinss/phone-number-handler/src"
	"github.com/ViniciusMartinss/phone-number-handler/src/domain"
	"github.com/ViniciusMartinss/phone-number-handler/src/helper"
)

type phoneUsecase struct {
	repository domain.PhoneRepository
}

func NewphoneUsecase(repository domain.PhoneRepository) domain.PhoneUsecase {
	return &phoneUsecase{repository}
}

func (p *phoneUsecase) List(filters ...map[string]string) []domain.PhoneReturnee {
	country := helper.GetCountryFilter(filters...)
	state := helper.GetStateFilter(filters...)

	phoneNumbers, err := p.repository.
		Find(country)

	if err != nil {
		return []domain.PhoneReturnee{}
	}

	return p.classifyPhoneNumbers(phoneNumbers, state)
}

func (p *phoneUsecase) classifyPhoneNumbers(phoneNumbers []string, stateFilter string) []domain.PhoneReturnee {
	var phoneClassification = make([]domain.PhoneReturnee, src.DEFAULT_SLICE_SIZE)

	for _, p := range phoneNumbers {
		phoneNumberSplit := REGEX_SLIPT_PHONE.FindAllStringSubmatch(p, 1)[PHONE_NUMBER_CONTENT]

		countryCode := phoneNumberSplit[COUNTRY_CODE_POSITION]
		phoneNumber := phoneNumberSplit[PHONE_NUMBER_POSITION]

		country := countries[countryCode]

		stateRegexResult, _ := regexp.MatchString(phoneRegex[country], p)
		state := phoneStatus[stateRegexResult]

		if stateFilter != "" && state != stateFilter {
			continue
		}

		phoneClassified := domain.PhoneReturnee{
			Country:     country,
			State:       state,
			CountryCode: fmt.Sprintf("+%s", countryCode),
			PhoneNumber: phoneNumber,
		}

		phoneClassification = append(phoneClassification, phoneClassified)
	}

	return phoneClassification
}
