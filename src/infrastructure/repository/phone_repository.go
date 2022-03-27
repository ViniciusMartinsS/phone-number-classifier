package repository

import (
	"fmt"

	"github.com/ViniciusMartinss/phone-number-handler/src"
	"github.com/ViniciusMartinss/phone-number-handler/src/domain"
	"github.com/jmoiron/sqlx"
)

type phoneRepository struct {
	database *sqlx.DB
}

type customer struct {
	Phone string `db:"phone"`
}

func NewPhoneRepository(connection *sqlx.DB) domain.PhoneRepository {
	return phoneRepository{database: connection}
}

func (p phoneRepository) Find(countryCode string) ([]string, error) {
	var customers []customer

	query := "SELECT phone FROM customer"

	if countryCode != "" {
		query = fmt.Sprintf(`%s WHERE phone LIKE '(%s)%%'`, query, countryCode)
	}

	err := p.database.
		Select(&customers, query)

	if err != nil {
		return nil, err
	}

	phoneNumbers := make([]string, src.DEFAULT_SLICE_SIZE)
	for _, c := range customers {
		phoneNumbers = append(phoneNumbers, c.Phone)
	}

	return phoneNumbers, nil
}
