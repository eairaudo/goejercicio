package miapi

import (
	"../../services/miapi"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	paramCountryID = "siteID"
)

func GetCountry(c *gin.Context)  {

	countryID := c.Param(paramCountryID)

	country , apiError := miapi.GetCountryFromApi(countryID)

	if apiError != nil{
		c.JSON(apiError.Status, apiError)
		return
	}


	rafaga := make(chan time.Time,4)

	go func() {
		for t := range time.Tick(3000 * time.Millisecond){
			for i := 0; i<3 ; i++ {
				rafaga <- t
			}
		}
	}()

	rafagaRequest := make(chan int, 15)

	for i:= 1; i <=15 ; i++{
		rafagaRequest <- i
	}

	close(rafagaRequest)

	for req := range rafagaRequest{
		<- rafaga
		c.JSON(http.StatusOK, country)
		fmt.Println(req)
	}
}
