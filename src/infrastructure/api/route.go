package api

import (
	"github.com/ViniciusMartinss/phone-number-handler/src/domain"
	domainInfrastructure "github.com/ViniciusMartinss/phone-number-handler/src/infrastructure/domain"
	"github.com/gin-gonic/gin"
)

type routes struct {
	phoneHandler domain.PhoneHandler
}

func NewRouter(phoneHandler domain.PhoneHandler) domainInfrastructure.Router {
	return &routes{phoneHandler}
}

func (r *routes) DefineRoutes() *gin.Engine {
	router := gin.Default()

	router.
		Group("phone").
		GET("/", r.phoneListHandler)

	return router
}
