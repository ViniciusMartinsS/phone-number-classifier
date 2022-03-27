package helper

import (
	"strings"

	"github.com/ViniciusMartinss/phone-number-handler/src"
)

func GetCountryFilter(filters ...map[string]string) string {
	if len(filters) == 0 {
		return ""
	}

	country := filters[src.FILTER_POSITION][src.COUNTRY]
	if country != "" {
		return strings.ToLower(country)
	}

	return ""
}

func GetStateFilter(filters ...map[string]string) string {
	if len(filters) == 0 {
		return ""
	}

	state := filters[src.FILTER_POSITION][src.STATE]
	if state != "" {
		return strings.ToUpper(state)
	}

	return ""
}
