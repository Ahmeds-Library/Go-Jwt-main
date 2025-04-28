package routes

import (
	handler "github.com/Ahmeds-Library/Go-Jwt/internal/api/handlers"
	"github.com/Ahmeds-Library/Go-Jwt/internal/api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RoutesHandler(r *gin.Engine) {

	r.POST("/signup", handler.Signup)
	r.POST("/login", handler.Login)
	r.POST("/upload", middleware.AuthMiddleware(), handler.Upload)
	r.GET("/getresuts", middleware.AuthMiddleware(), handler.GetResults)

	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

}
