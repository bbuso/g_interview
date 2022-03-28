package main
// I arbitrarily picked a framework, in this case, gin 
// it looked lightweight and this is a quick one-off project

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


type Database struct {
    Name string `json:"Name"`
    Code string `json:"Code"`
    Price int `json:"Price"`
}

database = []Database{
	Database{Name: "A12T-4GH7-QPL9-3N4M	", Code: "Lettuce", Price: 346},
	Database{Name: "E5T6-9UI3-TH15-QR88	", Code: "Peach", Price: 299},
	Database{Name: "YRT6-72AS-K736-L4AR	", Code: "Green Pepper", Price: 79},
	Database{Name: "TQ4C-VV6T-75ZX-1RMR	", Code: "Gala Apple", Price: 359},
}

func getAllProduce(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, database)
}

func getOneProduce() {}

func addProduce() {}



func main() {


	router := gin.Default()

	router.GET("/produce", getAllProduce)

	fmt.Println("test works")
	router.Run("localhost:8080")

}