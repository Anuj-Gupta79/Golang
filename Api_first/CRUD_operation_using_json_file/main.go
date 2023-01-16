package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostPayload struct {
	Name string `json:"name"`
}

type PutPayload struct {
	Branch string `json:"branch"`
}

func main() {
	r := gin.Default()

	// Basic Check Api
	chloFunc := func(c *gin.Context) {
		res := gin.H{
			"Sandesh": "Sabh Theek Chal Rha Hai Apna",
		}
		c.JSON(http.StatusOK, res)
	}
	r.GET("/chlo", chloFunc)

	mp := map[string]string{
		"Anuj":   "CSE Student",
		"Ramesh": "ECE Student",
		"Ishika": "Boitech Student",
	}

	//Get request
	getNameFunc := func(c *gin.Context) {
		res := gin.H{
			"records": mp,
		}
		c.JSON(http.StatusOK, res)
	}
	r.GET("/name", getNameFunc)

	//Post request
	postNameFunc := func(c *gin.Context) {
		var data PostPayload
		err := c.ShouldBindJSON(&data)
		if err != nil {
			log.Print("got error", err)
		}
		mp[data.Name] = "CSE Student"
		res := gin.H{
			"records": data,
		}
		c.JSON(http.StatusOK, res)
	}
	r.POST("/name", postNameFunc)

	// Delete Request
	deleteNameFunc := func(c *gin.Context) {

		// delete request using param
		// key := c.Params.ByName("id")

		//delete request using query
		key := c.Query("id")

		delete(mp, key)
		res := gin.H{
			"message": key + " deleted",
		}
		c.JSON(http.StatusOK, res)
	}
	// r.DELETE("/name/:id", deleteNameFunc)
	r.DELETE("/name", deleteNameFunc)

	//Put request
	putNameFunc := func(c *gin.Context){
		key := c.Params.ByName("id")
		var data PutPayload
		err := c.ShouldBindJSON(&data)
		if err != nil {
			log.Print("Got an error: ", err)
		}
		mp[key] = data.Branch

		res := gin.H{
			"record" : data,
		}
		c.JSON(http.StatusOK, res)
	}
	r.PUT("name/:id", putNameFunc)

	r.Run()
}
