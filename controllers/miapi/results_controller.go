package miapi

import (
	"../../services/miapi"
	"../../utils/apierrors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetResult(c *gin.Context)  {

	userID := c.Param(paramUserID)
	id, err := strconv.ParseInt(userID , 10 , 64)

	if err != nil {
		apiError := &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}

		c.JSON(apiError.Status, apiError)
		return
	}

	result , apiError := miapi.GetResultFromApi(id)

	if apiError != nil{
		c.JSON(apiError.Status, apiError)
		return
	}

	rafaga := make(chan time.Time,3)

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
		c.JSON(http.StatusOK, result)
		fmt.Println(req)
	}

}