package domain

import (
	"github.com/ViniciusMartinss/phone-number-handler/src/domain"
	"github.com/gin-gonic/gin"
)

type Router interface {
	DefineRoutes() *gin.Engine
}

type APIResponse struct {
	Status bool                   `json:"status"`
	Count  int                    `json:"count"`
	Result []domain.PhoneReturnee `json:"result"`
}
