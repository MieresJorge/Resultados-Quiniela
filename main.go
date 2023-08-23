package main

import (
	"github.com/gin-gonic/gin"
	controller "new.com/events/Controller"
	//"fmt"
)

func main() {

	router := gin.Default()

	controller.Resultados_totales(router)
    router.Run(":8080")
}

