package restful

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaokangwang/subscriptionSharingSite/common"
	"github.com/xiaokangwang/subscriptionSharingSite/model"
	"github.com/xiaokangwang/subscriptionSharingSite/subscription/containers"
	"github.com/xiaokangwang/subscriptionSharingSite/subscription/containers/base64urlline"
)

func (s *Server) PutProxyConfiguration(c *gin.Context) {
	var proxyServer model.ProxyServer
	var err error

	group := c.Param("group")
	privateToken := c.Param("privateToken")
	entryName := c.Param("entryName")

	publicToken := common.GetPublicTokenFromSecretToken(privateToken, s.SiteSecret)

	proxyServer.Group = group
	proxyServer.PublicToken = publicToken
	proxyServer.EntryName = entryName

	data, err := c.GetRawData()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = proxyServer.PutContentToKV(s.kv, data)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
	})
}

func (s *Server) GetProxyConfiguration(c *gin.Context) {
	var proxyServer model.ProxyServer
	var err error

	group := c.Param("group")
	publicToken := c.Param("publicToken")
	entryName := c.Param("entryName")

	proxyServer.Group = group
	proxyServer.PublicToken = publicToken
	proxyServer.EntryName = entryName

	data, err := proxyServer.GetContentFromKV(s.kv)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Data(200, "application/vnd.v2ray.subscription-singular", data)
}

func (s *Server) GetProxyConfigurationByGroup(c *gin.Context) {
	var err error

	group := c.Param("group")

	entries, err := model.ListAllEntriesByGroup(s.kv, group)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	var container containers.Container
	container.ServerSpecs = make([]containers.UnparsedServerConf, 0)
	for _, v := range entries {
		container.ServerSpecs = append(container.ServerSpecs, containers.UnparsedServerConf{
			Content:  []byte(v),
			KindHint: "URL",
		})
	}
	data, err := base64urlline.NewBase64URLLineWrapper().WrapSubscriptionContainerDocument(&container)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Data(200, "application/vnd.v2ray.subscription-base64urlline", data)
}
