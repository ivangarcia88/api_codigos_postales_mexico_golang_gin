package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	//Read Postal Code Database (JSON FILE)

	// Open our jsonFile
	jsonFile, err := os.Open("codigos_postales_reduced.json")
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// Parsing JSON File
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	//Running API endpoint
	r := gin.Default()
	r.GET("/code/:code", func(c *gin.Context) {
		c.JSON(http.StatusOK, result[c.Param("code")])
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
