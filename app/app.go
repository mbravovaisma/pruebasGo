package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/go-meli-toolkit/gingonic/mlhandlers"
)

var (
	router *gin.Engine
)

func mapRoutesController() {
	router.GET("/ping", Ping)
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

//Start configures all required dependencies and bind the port.
func Start() {
	ConfigureRouter()
	router.Run(":8080")
}

//ConfigureRouter configure router
func ConfigureRouter() {
	router = mlhandlers.DefaultMeliRouter()
	router.RedirectFixedPath = false
	router.RedirectTrailingSlash = true

	mapRoutesController()
}

func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	router.ServeHTTP(w, req)
}
