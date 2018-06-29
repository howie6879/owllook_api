package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchChapters returns all chapters that you serached
func SearchChapters(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"title": "chapters demo"})
}
