package controller

import "github.com/gin-gonic/gin"

type BaseController interface{
	RegisterRoutes(e *gin.Engine)
}