package restful

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaokangwang/subscriptionSharingSite/keyValueStorage"
)

type Server struct {
	SiteSecret string

	kv keyValueStorage.ScopedPersistentStorage

	prefix string
}

func NewServer(kv keyValueStorage.ScopedPersistentStorage, SiteSecret string) *Server {
	return &Server{
		SiteSecret: SiteSecret,
		kv:         kv,
	}
}

func (s *Server) RegisterHandlers(engine *gin.Engine, apiPrefix string) {
	s.prefix = apiPrefix

	api := engine.Group(apiPrefix)
	{
		api.GET("/token", s.GenerateToken)
		api.POST("/proxy/:group/:privateToken/:entryName", s.PutProxyConfiguration)
		api.GET("/proxy/:group/:publicToken/:entryName", s.GetProxyConfiguration)
		api.GET("/GetProxyConfigurationByGroup/:group", s.GetProxyConfigurationByGroup)
	}

}
