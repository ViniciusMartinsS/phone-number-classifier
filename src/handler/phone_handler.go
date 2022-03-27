package handler

import (
	"github.com/ViniciusMartinss/phone-number-handler/src"
	"github.com/ViniciusMartinss/phone-number-handler/src/domain"
	"github.com/ViniciusMartinss/phone-number-handler/src/helper"
)

type phoneHandler struct {
	phoneUsecase domain.PhoneUsecase
}

func NewPhoneHandler(phoneUsecase domain.PhoneUsecase) domain.PhoneHandler {
	return &phoneHandler{phoneUsecase}
}

func (p *phoneHandler) List(filters ...map[string]string) []domain.PhoneReturnee {
	if len(filters) == 0 {
		return p.phoneUsecase.
			List()
	}
	country := helper.GetCountryFilter(filters...)

	countryCode, existCountry := src.Codes[country]
	if !existCountry && country != "" {
		return []domain.PhoneReturnee{}
	}

	filters[src.FILTER_POSITION][src.COUNTRY] = countryCode
	return p.phoneUsecase.
		List(filters...)
}
