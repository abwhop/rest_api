package rest_api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerInterface interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
	SetService(ServiceInterface)
	GetService() ServiceInterface
}

type Controller struct {
	service ServiceInterface
}

func (cn *Controller) SetService(s ServiceInterface) {
	cn.service = s
}

func (cn *Controller) GetService() ServiceInterface {
	return cn.service
}

func (cn *Controller) Index(c *gin.Context) {
	if result, err := cn.GetService().GetList(c.Request.Context(), func(obj interface{}) error { return c.ShouldBind(obj) }, func(user interface{}) error { return cn.getUser(c, &user) }); err != nil {
		ErrorHandler(err, c)
		return
	} else {
		c.JSON(http.StatusOK, result)
	}
	return
}

func (cn *Controller) Show(c *gin.Context) {
	if result, err := cn.GetService().GetById(c.Request.Context(), c.Param("id"), func(obj interface{}) error { return c.Bind(obj) }, func(user interface{}) error {
		return cn.getUser(c, &user)
	}); err != nil {
		ErrorHandler(err, c)
		return
	} else {
		c.JSON(http.StatusOK, result)
	}
	return
}

func (cn *Controller) Store(c *gin.Context) {
	if result, err2 := cn.GetService().Create(c.Request.Context(), func(obj interface{}) error { return c.ShouldBind(obj) }, func(obj interface{}) error { return c.Bind(obj) }, func(user interface{}) error { return cn.getUser(c, &user) }); err2 != nil {
		ErrorHandler(err2, c)
		return
	} else {
		c.JSON(http.StatusOK, result)
	}
	return
}

func (cn *Controller) Update(c *gin.Context) {
	if result, err2 := cn.GetService().Update(c.Request.Context(), c.Param("id"), func(obj interface{}) error { return c.ShouldBind(obj) }, func(obj interface{}) error { return c.Bind(obj) }, func(user interface{}) error { return cn.getUser(c, &user) }); err2 != nil {
		ErrorHandler(err2, c)
		return
	} else {
		c.JSON(http.StatusOK, result)
	}
	return
}

func (cn *Controller) Destroy(c *gin.Context) {
	if result, err := cn.GetService().Delete(c.Request.Context(), c.Param("id"), func(obj interface{}) error { return c.ShouldBind(obj) }, func(user interface{}) error { return cn.getUser(c, &user) }); err != nil {
		ErrorHandler(err, c)
		return
	} else {
		c.JSON(http.StatusOK, result)
	}
	return
}

func (cn *Controller) getUser(c *gin.Context, user interface{}) error {
	err := json.Unmarshal(c.MustGet("user").([]byte), &user)
	if err != nil {
		return fmt.Errorf(`error unmarshel jwt payload %w`, err)
	}
	return nil
}
