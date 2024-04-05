package routes

import (
	"boilerplate/app/http/middlewares"
	v1 "boilerplate/routes/v1"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.RequestLogger())
	v1.SetupRouter(r)
	return r
}
