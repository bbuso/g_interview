package main

// I arbitrarily picked a framework, in this case, gin
// it looked lightweight and this is a quick one-off project

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Database struct {
	Name  string `json:"Name"`
	Code  string `json:"Code"`
	Price int    `json:"Price"`
}

var database = []Database{
	{Name: "A12T-4GH7-QPL9-3N4M	", Code: "Lettuce", Price: 346},
	{Name: "E5T6-9UI3-TH15-QR88	", Code: "Peach", Price: 299},
	{Name: "YRT6-72AS-K736-L4AR	", Code: "Green Pepper", Price: 79},
	{Name: "TQ4C-VV6T-75ZX-1RMR	", Code: "Gala Apple", Price: 359},
}

func getAllProduce(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, database)
}
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getOneProduce() {}

func addProduce() {}

func main() {

	router := gin.Default()

	router.GET("/produce", getAllProduce)
	router.GET("/produce/:id", getAllProduce)

	fmt.Println("test works")
	router.Run("localhost:8080")

}
