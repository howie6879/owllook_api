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

// SearchAuthors returns all novels resource that you serached
func SearchAuthors(c *gin.Context) {
	var currentRule config.NovelRule
	var ok bool
	novelName := c.Param("name")
	novelSource := c.Param("source")
	if novelName != "" {
		currentRule, ok = NovelsRulesMap[novelSource+"_1"]
		log.Println(ok)
		if ok == false {
			currentRule, ok = NovelsRulesMap[novelSource]
		}
		log.Println(ok)
		if ok {
			resultData, err := common.FetchHtml(novelName, currentRule)
			if err != nil {
				log.Println("Request URL error", err)
				c.JSON(http.StatusOK, gin.H{"statue": 0, "msg": "Request error"})
			} else {
				c.JSON(http.StatusOK, gin.H{"status": 1, "info": resultData})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"statue": 0, "msg": "Parameter error"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"statue": 0, "msg": "Parameter name can't be empty"})
	}
}

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
				c.JSON(http.StatusOK, gin.H{"statue": 0, "msg": "Request error"})
			} else {
				c.JSON(http.StatusOK, gin.H{"status": 1, "info": resultData})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"statue": 0, "msg": "Parameter error"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"statue": 0, "msg": "Parameter name can't be empty"})
	}
}
