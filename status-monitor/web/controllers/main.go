package controllers

import (
	"status-monitor/web/serializer"
	"status-monitor/web/services"
	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	service := services.PingService{}

	context.String(200, service.Ping())
}

func Index(context *gin.Context) {
	service := services.IndexService{}

	context.HTML(200, "index.html",
		serializer.BuildIndexResponse(service.Index()))
}
