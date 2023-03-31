package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Nombre   string
	Apellido string
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.POST("/saludo", func(ctx *gin.Context) {
		var p Persona
		if err := json.NewDecoder(ctx.Request.Body).Decode(&p); err != nil {
			return
		}
		//if err := ctx.BindJSON(&p); err != nil {
		//	return
		//}
		ctx.String(200, "hola "+p.Nombre+" "+p.Apellido)
	})

	router.Run()
}
