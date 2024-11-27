package main

import (
	routes "apisimples/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Registrar rotas
	routes.RegisterRoutes(r)

	// Iniciar o servidor
	r.Run(":8080")
}
