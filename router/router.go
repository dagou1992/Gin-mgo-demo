package router

import (
	"github.com/gin-gonic/gin"
	"../controller"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.POST("/insert", controller.Insert)
	r.POST("/find", controller.Find)
	r.POST("/update", controller.Update)
	r.POST("/delete", controller.Delete)

	return r
}