package api

import (
	"github.com/ViniciusMartinss/phone-number-handler/src/domain"
	"github.com/gin-gonic/gin"
)

type router interface {
	DefineRoutes() *gin.Engine
}

type routes struct {
	phoneHandler domain.PhoneHandler
}

func NewRouter(phoneHandler domain.PhoneHandler) router {
	return &routes{phoneHandler}
}

func (r *routes) DefineRoutes() *gin.Engine {
	router := gin.Default()

	router.
		Group("phone").
		GET("/", r.phoneListHandler)

	return router
}
