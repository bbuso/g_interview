package main

// I arbitrarily picked a framework, in this case, gin
// it looked lightweight and this is a quick one-off project

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type Produce struct {
	Name  string `json:"Name"`
	Code  string `json:"Code"`
	Price int    `json:"Price"`
}

var database = []Produce{
	{Code: "A12T-4GH7-QPL9-3N4M", Name: "Lettuce", Price: 346},
	{Code: "E5T6-9UI3-TH15-QR88", Name: "Peach", Price: 299},
	{Code: "YRT6-72AS-K736-L4AR", Name: "Green Pepper", Price: 79},
	{Code: "TQ4C-VV6T-75ZX-1RMR", Name: "Gala Apple", Price: 359},
}

func getAllProduce(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, database)
}
func getProduceByName(context *gin.Context) {
	name := strings.ToLower(context.Param("name"))

	for _, entry := range database {
		if strings.ToLower(entry.Name) == name {
			context.IndentedJSON(http.StatusOK, entry)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "food not found"})
}

func addProduce(context *gin.Context) {
	var newProduce Produce

	if err := context.BindJSON(&newProduce); err != nil {
		return
	}
	if !validateProduceCode(newProduce) {
		context.IndentedJSON(http.StatusBadRequest, "Invalid Produce Code")
		return
	}

	database = append(database, newProduce)
	context.IndentedJSON(http.StatusCreated, database)

}

func validateProduceName(produce Produce) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9]*$", produce.Code)
	return match
}

func validateProduceCode(produce Produce) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}$", produce.Code)
	return match
}

func deleteProduce(context *gin.Context) {
	name := strings.ToLower(context.Param("name"))

	for index, entry := range database {
		if strings.ToLower(entry.Name) == name {
			database = RemoveIndex(database, index)
			return
		}
	}
	return
}

func RemoveIndex(s []Produce, index int) []Produce {
	return append(s[:index], s[index+1:]...)
}

func main() {

	router := gin.Default()

	router.GET("/produce", getAllProduce)
	router.GET("/produce/:name", getProduceByName)
	router.POST("/produce", addProduce)
	router.DELETE("/produce/:name", deleteProduce)
	router.Run("localhost:8080")

}
