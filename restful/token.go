package restful

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiaokangwang/subscriptionSharingSite/common"
)

func SecureGenerateRandomString() string {
	var entropy [32]byte
	_, err := rand.Read(entropy[:])
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(entropy[:])
}

func (s *Server) GenerateToken(c *gin.Context) {
	secretToken := SecureGenerateRandomString()
	publicToken := common.GetPublicTokenFromSecretToken(secretToken, s.SiteSecret)

	host := c.GetHeader("Host")
	shareURL := fmt.Sprintf("http://%s/%s/proxy/%s/%s/%s", host, s.prefix, "Sharing", secretToken, "SharedProxy")
	receiveURL := fmt.Sprintf("http://%s/%s/proxy/%s/%s/%s", host, s.prefix, "Sharing", publicToken, "SharedProxy")

	c.JSON(200, gin.H{
		"secretToken": secretToken,
		"publicToken": publicToken,
		"shareURL":    shareURL,
		"receiveURL":  receiveURL,
	})
}
