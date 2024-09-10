package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaokangwang/subscriptionSharingSite/keyValueStorage/rediskv"
	"github.com/xiaokangwang/subscriptionSharingSite/restful"
	"os"
)

func mustGetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic("missing required environment variable " + key)
	}
	return value
}

func main() {
	redisURL := mustGetEnv("REDIS_URL")
	redisKV := rediskv.NewKVFromURL(redisURL)

	serverSecret := mustGetEnv("SERVER_SECRET")
	apiPrefix := mustGetEnv("API_PREFIX")
	port := mustGetEnv("PORT")

	r := gin.Default()

	server := restful.NewServer(redisKV, serverSecret)

	server.RegisterHandlers(r, apiPrefix)

	r.Run(":" + port)

}
