package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/howie6879/owllook_api/apis"
)

// Index returns a home page
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"projetc_name": "owllook api dmeo", "version": "v1", "url": "/v1"})
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", Index)
	v1 := router.Group("v1")
	v1.GET("/novels/:name/:source", apis.SearchNovels)
	router.Run()
}
