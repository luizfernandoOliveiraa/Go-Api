package router

import "github.com/gin-gonic/gin"

// Para ter uma "classe"/"funcao" publica, somente precisamos inicar a variavel com maiusculo

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Run the server
	router.Run("localhost:8080")
}
