package main

import (
	"github.com/Ahmeds-Library/Go-Jwt/internal/api/routes"
	"github.com/Ahmeds-Library/Go-Jwt/internal/database"
	"github.com/gin-gonic/gin"
)

// @title User Auth Application APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @in header
// @name token
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8001
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	database.ConnectDatabase()
	r := gin.Default()

	routes.RoutesHandler(r)

	r.Run(":8001")
}
