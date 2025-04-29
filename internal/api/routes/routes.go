package routes

import (
	handler "github.com/Ahmeds-Library/Go-Jwt/internal/api/handlers"
	"github.com/Ahmeds-Library/Go-Jwt/internal/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RoutesHandler(r *gin.Engine) {

	r.POST("/signup", handler.Signup)
	r.POST("/login", handler.Login)
	r.POST("/upload", services.AuthMiddleware(), handler.Upload)
	r.GET("/getresuts", services.AuthMiddleware(), handler.GetResults)

	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

}
