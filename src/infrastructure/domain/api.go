package domain

import "github.com/gin-gonic/gin"

type Router interface {
	DefineRoutes() *gin.Engine
}
 