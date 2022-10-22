package main

import (
	"github.com/Nanthat/sa-65-example/controller"
	"github.com/Nanthat/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	r := gin.Default()
	//r.Use(CORSMiddleware())

	// Cart Routes
	r.GET("/carts", controller.ListCart)
	r.GET("/cart/:id", controller.GetCart)
	r.POST("/carts", controller.CreateCart)
	r.PATCH("/carts", controller.UpdateCart)
	r.DELETE("/carts/:id", controller.DeleteCart)

	// Order Routes
	r.GET("/orders", controller.ListOrder)
	r.GET("/order/:id", controller.GetOrder)
	r.POST("/orders", controller.CreateOrder)
	r.PATCH("/orders", controller.UpdateOrder)
	r.DELETE("/orders/:id", controller.DeleteOrder)

	// Run the server
	r.Run()

}

// func CORSMiddleware() gin.HandlerFunc {

// 	return func(c *gin.Context) {

// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

// 		if c.Request.Method == "OPTIONS" {

// 			c.AbortWithStatus(204)

// 			return

// 		}

// 		c.Next()

// 	}

// }
