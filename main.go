package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Song struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Genre  string `json:"genre"`
}

var songs []Song
var carnetNumber = "201700672"

func main() {
	// Configuración del router Gin
	router := gin.Default()

	// Endpoint POST para crear una nueva canción
	router.POST("/song", func(c *gin.Context) {
		var newSong Song
		// Bind JSON recibido en el cuerpo de la solicitud a la estructura Song
		if err := c.BindJSON(&newSong); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Agregar la nueva canción a la lista de canciones
		songs = append(songs, newSong)

		// Responder con los detalles de la canción recién creada
		c.JSON(http.StatusCreated, newSong)
	})
	router.GET("/carnet", func(c *gin.Context) {
		c.String(http.StatusOK, "Número de carnet: %s", carnetNumber)
	})

	// Correr el servidor en el puerto 8080
	router.Run(":8080")
}
