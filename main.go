package main

import (
	"context"
	"example/myco-api/ent"
	"example/myco-api/ent/grainjar"
	"log"

	"example/myco-api/models"
	"fmt"
	"net/http"
	"time"
	"example/myco-api/keys"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	keys := keys.GetKeys()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", keys.User, keys.Password, keys.Host, keys.Port, keys.Database)
	client, err := ent.Open("mysql", connectionString)
    if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
    }
    defer client.Close()
    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
	postGrainJar(context.Background(), client)

    router := gin.Default()
    router.GET("/grainJars", getGrainJars)
	router.GET ("/grainJar/:id", getGrainJar)
//	router.POST("/grainJar", postGrainJar)

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

func postGrainJar(ctx context.Context, client *ent.Client) (*ent.GrainJar, error) {
    u, err := client.GrainJar.
        Create().
        SetGrain("Rye Malt Blend").
		SetHarvestDate(time.Now()).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed creating user: %w", err)
    }
    log.Println("user was created: ", u)
    return u, nil
}

func QueryGrainJar(ctx context.Context, client *ent.Client) (*ent.GrainJar, error) {
    u, err := client.GrainJar.
        Query().
        Where(grainjar.ID(1)).
        // `Only` fails if no user found,
        // or more than 1 user returned.
        Only(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed querying user: %w", err)
    }
    log.Println("user returned: ", u)
    return u, nil
}
