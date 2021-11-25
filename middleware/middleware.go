package middleware

import (
	"GOLANG/todo/routes"

	"github.com/gin-gonic/gin"
)

//InitMiddleWare initializes middleware on all types of requests
func InitMiddleWare(g *gin.Engine) {

	open := g.Group("/o")
	open.Use(OpenMiddleWare())
	// Implement other middleware for restricted and role based routes. Adding an dummy middleware for open routes
	routes.InitRoutes(open)
}

//OpenMiddleWare applies for all the Open request
func OpenMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
