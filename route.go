package rest_api

import (
	"github.com/gin-gonic/gin"
)

// RouterInterface Base route interface
type RouterInterface interface {
	setController(ControllerInterface)
}

// Route Base route type
type Route struct {
	Controller ControllerInterface
	Router     *gin.RouterGroup
}

// SetController Method for adding a controller instance
func (r *Route) SetController(c ControllerInterface) {
	r.Controller = c
}

// Crud Base CRUD operations
func (r *Route) Crud() *gin.RouterGroup {
	//Get List of booking
	r.Router.GET("", r.Controller.Index)
	r.Router.GET("/", r.Controller.Index)
	//Get One booking
	r.Router.GET("/:id", r.Controller.Show)
	r.Router.GET("/:id/", r.Controller.Show)
	//Create One booking
	r.Router.POST("", r.Controller.Store)
	r.Router.POST("/", r.Controller.Store)
	//Update One booking
	r.Router.PUT("/:id", r.Controller.Update)
	r.Router.PUT("/:id/", r.Controller.Update)
	//Update one booking in a non-standard way
	r.Router.PUT("", r.Controller.Update)
	r.Router.PUT("/", r.Controller.Update)
	//DELETE One booking
	r.Router.DELETE("/:id", r.Controller.Destroy)
	r.Router.DELETE("/:id/", r.Controller.Destroy)
	//DELETE one booking in a non-standard way
	r.Router.DELETE("", r.Controller.Destroy)
	r.Router.DELETE("/", r.Controller.Destroy)
	return r.Router
}
