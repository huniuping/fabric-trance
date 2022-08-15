package assemble

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	router.GET("/query", Query)
	router.GET("/queryAssembleDetailByID", QueryAssembleDetailByID)
	router.GET("/queryByworkOrderId", QueryByworkOrderId)
	router.POST("/delete", Delete)
	router.POST("/set", Set)
	router.POST("/update", Update)
}
