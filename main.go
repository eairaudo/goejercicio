package main

import (
	"./controllers/miapi"
	"github.com/gin-gonic/gin"
	"sync"
)

const (
	port = ":8090"
)


var	(
	router = gin.Default()
)

var Wg sync.WaitGroup

func main()  {

	router.GET("/user/:userID" , miapi.GetUser)
	router.GET("/site/:siteID" , miapi.GetSite)
	router.GET("/countries/:countryID" , miapi.GetCountry)
	router.GET("/result/:userID" , miapi.GetResult)

	router.Run(port)

}