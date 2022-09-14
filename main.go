package main

import (
	"github.com/A-u-usman/portal-go_clean_api.git/controllers"
	"github.com/gin-gonic/gin"
)

var (
	portalController controllers.PortalController = controllers.PortalControllerImp()
)

func main() {

	r := gin.Default()

	portalRoutes := r.Group("user/")
	{
		portalRoutes.POST("/login", portalController.Login)
		portalRoutes.POST("/register", portalController.Register)
		portalRoutes.POST("/adduser", portalController.AddUser)
		portalRoutes.POST("/addsubject", portalController.AddSubject)

	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
