package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PortalController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	AddUser(ctx *gin.Context)
	AddSubject(ctx *gin.Context)
}

type portalController struct{}

func PortalControllerImp() PortalController {
	return &portalController{}
}

func (c *portalController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "You have successfully logged in",
	})
}

func (c *portalController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "You have successfully registered new student",
	})
}

func (c *portalController) AddUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "You have successfully added a user",
	})
}

func (c *portalController) AddSubject(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "You have successfully added a subject",
	})
}
