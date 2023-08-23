package controlador

import (
	"github.com/gin-gonic/gin"
	models "new.com/events/Model"
	"net/http"
)

/*
func Resultados_totales(router *gin.Engine) {
	router.GET("/Previa/Ciudad/", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.Resultados_quiniela_Ciuad_1 + models.Resultados_quiniela_Provincia_1)
	})
}
*/
func Resultados_totales(router *gin.Engine) {
	router.GET("/Nacional/", func(c *gin.Context) {
		resultados := []models.User{
			models.Resultados_quiniela_Ciuad_1,
			models.Resultados_quiniela_Provincia_1,
			models.Resultados_quiniela_EntreRios_1,
			models.Resultados_quiniela_Cordoba_1,
			models.Resultados_quiniela_Mendoza_1,
			models.Resultados_quiniela_SantaFe_1,
		}
		c.JSON(http.StatusOK, resultados)
	})

	router.GET("/Nacional/:Quiniela", func(c *gin.Context) {
		quiniela := c.Param("Quiniela")

		var resultados []models.User
		for _, user := range []models.User{
			models.Resultados_quiniela_Ciuad_1,
			models.Resultados_quiniela_Provincia_1,
			models.Resultados_quiniela_EntreRios_1,
			models.Resultados_quiniela_Cordoba_1,
			models.Resultados_quiniela_Mendoza_1,
			models.Resultados_quiniela_SantaFe_1,
		} {
			resultado, exists := user.Resultado[quiniela]
			if exists {
				resultados = append(resultados, models.User{
					Posicion:  user.Posicion,
					Resultado: map[string]string{quiniela: resultado},
				})
			}
		}

		if len(resultados) > 0 {
			c.JSON(http.StatusOK, resultados)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Quiniela not found"})
		}
	})

	router.GET("/Previa/Ciudad/:Quiniela", func(c *gin.Context) {
		quiniela := c.Param("Quiniela")
		resultado := models.Resultados_quiniela_Ciuad_1.Resultado[quiniela]
		c.JSON(http.StatusOK, gin.H{"Quiniela": quiniela, "Resultado": resultado})
	})

}