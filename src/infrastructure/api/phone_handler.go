package api

import (
	"net/http"

	"github.com/ViniciusMartinss/phone-number-handler/src"
	"github.com/ViniciusMartinss/phone-number-handler/src/domain"
	"github.com/gin-gonic/gin"
)

type apiResponse struct {
	Status bool                   `json:"status"`
	Count  int                    `json:"count"`
	Result []domain.PhoneReturnee `json:"result"`
}

func (r *routes) phoneListHandler(c *gin.Context) {
	filters := setFilters(c)

	result := r.phoneHandler.
		List(filters)

	response := apiResponse{
		Status: true,
		Count:  len(result),
		Result: result,
	}

	c.JSON(http.StatusOK, response)
}

func setFilters(context *gin.Context) map[string]string {
	country := context.Query(src.COUNTRY)
	state := context.Query(src.STATE)

	return map[string]string{src.COUNTRY: country, "state": state}
}
