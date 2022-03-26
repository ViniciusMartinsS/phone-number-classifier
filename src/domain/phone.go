package domain

type PhoneHandler interface {
	List(...map[string]string) []PhoneReturnee
}

type PhoneUsecase interface {
	List(...map[string]string) []PhoneReturnee
}

type PhoneRepository interface {
	Find(countryCode string) ([]string, error)
}

type PhoneReturnee struct {
	Country     string `json:"country"`
	State       string `json:"state"`
	CountryCode string `json:"countryCode"`
	PhoneNumber string `json:"PhoneNumber"`
}
