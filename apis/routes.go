package apis

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/obsidianwallet/incognito-tools/decoder"
)

func StartService() error {
	log.Println("Starting service...") //
	r := gin.Default()

	r.Use(cors.Default())
	store := persistence.NewInMemoryStore(time.Second)

	decoderG := r.Group("/decoder")
	decoderG.POST("/decodekey", cache.CachePage(store, 30*time.Second, decodeWalletKey))

	generatorG := r.Group("/generator")
	generatorG.POST("/genota", cache.CachePage(store, 30*time.Second, genOTA))

	return r.Run(":8080")
}

func decodeWalletKey(c *gin.Context) {
	var req DecodedWalletRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Error": err.Error()})
		return
	}

	result, err := decoder.DecodeWalletKey(req.Key, req.ShardsNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Result": result})
}

func genOTA(c *gin.Context) {

}
