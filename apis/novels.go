package apis

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/howie6879/owllook_api/common"
	"github.com/howie6879/owllook_api/config"
)

var (
	NovelsRulesMap = config.NovelsRulesMap
)

// SearchNovels returns all novels resource that you serached
func SearchNovels(c *gin.Context) {
	novelName := c.Param("name")
	novelSource := c.Param("source")
	if novelName != "" {
		currentRule, ok := NovelsRulesMap[novelSource]
		if ok {
			resultData, err := common.FetchHtml(novelName, currentRule)
			if err != nil {
				log.Println("Request URL error", err)
				c.JSON(http.StatusOK, gin.H{"statue": 0, "info": "Request error"})
			}
			c.JSON(http.StatusOK, gin.H{"status": 1, "info": resultData})
		} else {
			c.JSON(http.StatusOK, gin.H{"statue": 0, "info": "Parameter error"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"statue": 0, "info": "Parameter name can't be empty"})
	}
}
