package main

import (
	"net/http"
    "github.com/gin-gonic/gin"
	"time"
	"example/myco-api/models"
)

func main() {
    router := gin.Default()
    router.GET("/grainJars", getGrainJars)
	router.GET ("/grainJar/:id", getGrainJar)
	router.POST("/grainJar", postGrainJar)

    router.Run("localhost:8080")
}

var grainJars = []models.GrainJar {
	{
		ID: "1", 
		Grain: "Rye", 
		SporeSyringe: models.SporeSyringe {
			ID: "1",
			Species: "Oyster",
			Supplier: "North Spore",
			Received: time.Now(),
		},
	},
	{
		ID: "2", 
		Grain: "Brown Rice", 
		SporeSyringe: models.SporeSyringe {
			ID: "1",
			Species: "Oyster",
			Supplier: "North Spore",
			Received: time.Now(),
		},
	},
	{
		ID: "3", 
		Grain: "Cracked Corn", 
		SporeSyringe: models.SporeSyringe {
			ID: "1",
			Species: "Oyster",
			Supplier: "North Spore",
			Received: time.Now(),
		},
	},
}

func getGrainJars(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, grainJars)
}

func getGrainJar(c *gin.Context) {
	id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range grainJars {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "grain jar not found"})
}

func postGrainJar(c *gin.Context) {
    var newGrainJar models.GrainJar

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newGrainJar); err != nil {
        return
    }

    // Add the new album to the slice.
    grainJars = append(grainJars, newGrainJar)
    c.IndentedJSON(http.StatusCreated, newGrainJar)
}
