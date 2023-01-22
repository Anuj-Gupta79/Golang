package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	id int `json:"id" json`
	name string `json:"name" json`
	year int `json:"year" json`
	color string `json:"color" json`
	patonValues string `json:"pantone_value" json`
}

func main(){
	r := gin.Default()
	user := []User{} 

	getDataFunc := func(c *gin.Context){
		data, err1 := ioutil.ReadFile("./dummy.json")
		if err1 != nil{
			log.Fatal("error")
		}
		jsonData := string(data)
		fmt.Println(jsonData)
		err2 := json.Unmarshal([]byte(jsonData),&user)
		if err2 != nil{
			log.Fatal("error")
		}
		fmt.Println(user)
		res := gin.H{
			"message": user,
		}
		c.JSON(http.StatusOK, res)
	}
	r.GET("/data",getDataFunc)
	r.Run()
}