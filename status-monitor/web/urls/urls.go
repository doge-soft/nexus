package urls

import (
	"github.com/gin-gonic/gin"
	"status-monitor/web/controllers"
)

type URL struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}

var Urls = []URL{
	URL{"/", "GET", controllers.Index},
}
