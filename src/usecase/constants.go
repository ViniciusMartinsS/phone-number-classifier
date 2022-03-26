package usecase

import "regexp"

var REGEX_SLIPT_PHONE = regexp.MustCompile(`\((.*?)\)\s(.+)`)

var countries = map[string]string{
	"237": "Cameroon",
	"251": "Ethiopia",
	"212": "Morocco",
	"258": "Mozambique",
	"256": "Uganda",
}

var phoneRegex = map[string]string{
	"Cameroon":   `\(237\)\ ?[2368]\d{7,8}$`,
	"Ethiopia":   `\(251\)\ ?[1-59]\d{8}$`,
	"Morocco":    `\(212\)\ ?[5-9]\d{8}$`,
	"Mozambique": `\(258\)\ ?[28]\d{7,8}$`,
	"Uganda":     `\(256\)\ ?\d{9}$`,
}

var phoneStatus = map[bool]string{
	true:  "OK",
	false: "NOK",
}

const (
	PHONE_NUMBER_CONTENT = iota
	COUNTRY_CODE_POSITION
	PHONE_NUMBER_POSITION
)
