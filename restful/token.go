package restful

import (
	"crypto/rand"
	"encoding/hex"
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
	c.JSON(200, gin.H{
		"secretToken": secretToken,
		"publicToken": publicToken,
	})
}
