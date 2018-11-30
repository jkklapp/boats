package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
	"os"
	"strings"
)

// InitRoutes initializes routes for the backend API
// e.g. example.com/api/v1/healthcheck
func InitRoutes(router *gin.Engine) *gin.Engine {
	// Group of API calls
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/ingest", ingestor)
		}
	}

	return router
}

func ingestor(c *gin.Context) {
	file, err := os.Open("data/data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if i != 0 {
			// Iterate over the file and send data to datastore one line at a time.
			s := strings.Split(scanner.Text(), ",")
			values := map[string]string{"id": s[0], "name": s[1], "email": s[2], "mobile_number": s[3]}
			jsonValue, _ := json.Marshal(values)
			resp, err := http.Post("http://boats_datastore_1:8001/api/v1/process", "application/json", bytes.NewBuffer(jsonValue))
			fmt.Println(resp)
			fmt.Println(err)
			fmt.Println(jsonValue)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
