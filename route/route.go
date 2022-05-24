package route

import (
	"golang_todo_app_ashrtech/auth"
	"golang_todo_app_ashrtech/controller"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	authMiddleware, err := auth.SetupAuth()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Ashr Todo App")
	})

	api := router.Group("/api")
	{
		api.POST("/login", authMiddleware.LoginHandler)
		api.POST("/register", controller.RegisterEndPoint)
		api.GET("/todos", authMiddleware.MiddlewareFunc(), controller.FetchAllTask)
		api.POST("/store", authMiddleware.MiddlewareFunc(), controller.CreateTask)
		api.PUT("/update/:todo_id", authMiddleware.MiddlewareFunc(), controller.UpdateTask)
		api.PUT("/updatestatus/:todo_id", authMiddleware.MiddlewareFunc(), controller.UpdateStatusTask)
		api.DELETE("/delete/:todo_id", authMiddleware.MiddlewareFunc(), controller.DeleteTask)

	}

	authorization := router.Group("/auth")
	authorization.GET("/refresh_token", authMiddleware.RefreshHandler)

	return router
}
