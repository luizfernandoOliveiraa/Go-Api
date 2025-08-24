package router

import "github.com/gin-gonic/gin"

// Para ter uma "classe"/"funcao" publica, somente precisamos inicar a variavel com maiusculo 

func Initialize() {
	// Inicializa o Router utilizando as configurações Default do gin
	router := gin.Default()
	// Defino uma rota em GO
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run("localhost:8080")
}
