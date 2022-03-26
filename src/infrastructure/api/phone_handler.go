package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *routes) phoneListHandler(c *gin.Context) {
	filters := setFilters(c)

	response := r.phoneHandler.
		List(filters)

	c.JSON(http.StatusOK, response)
}

func setFilters(context *gin.Context) map[string]string {
	country := context.Query("country")
	state := context.Query("state")

	return map[string]string{"country": country, "state": state}
}
