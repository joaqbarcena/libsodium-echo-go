package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandleError interface {
	Status() int
	Error() string
}

type Handler struct {
	Path        string
	HandlerFunc func(*gin.Context) HandleError
	Method      string
}

func Wire(router *gin.Engine, routing ...*Handler) {
	for _, route := range routing {
		var methodHandler func(string, ...gin.HandlerFunc) gin.IRoutes
		switch route.Method {
		case http.MethodGet:
			methodHandler = router.GET
		case http.MethodHead:
			methodHandler = router.HEAD
		case http.MethodPost:
			methodHandler = router.POST
		case http.MethodPut:
			methodHandler = router.PUT
		case http.MethodPatch:
			methodHandler = router.PATCH
		case http.MethodDelete:
			methodHandler = router.DELETE
		case http.MethodOptions:
			methodHandler = router.OPTIONS
		default:
			panic("Http Method not handled")
		}

		methodHandler(route.Path, handleErrorWrapper(route.HandlerFunc))
	}
}

func handleErrorWrapper(handlerFunc func(*gin.Context) HandleError) func(*gin.Context) {
	return func(c *gin.Context) {
		if err := handlerFunc(c); err != nil {
			c.JSON(err.Status(), err.Error())
		}
	}
}
