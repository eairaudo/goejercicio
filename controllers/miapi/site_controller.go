package miapi

import (
	"../../services/miapi"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	paramSiteID = "siteID"
)

func GetSite(c *gin.Context)  {

	siteID := c.Param(paramSiteID)

	sites , apiError := miapi.GetSiteFromApi(siteID)

	if apiError != nil{
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, sites)
}
