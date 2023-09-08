package main

import (
	"api-supermarket/cmd/server/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// creo el server principal
	server := gin.Default()
	// creo un grupo para productos
	group := server.Group("/products")
	// llamo a la creacion de un router.
	router := handlers.NewProductRouter(group)

	// ejecuto el metodo que crea las rutas para que esten registradas
	router.ProductRoutes()

	// corro mi server
	server.Run(":8080")
}
