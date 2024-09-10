package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GetPublicTokenFromSecretToken(secretToken, siteSecret string) string {
	generator := hmac.New(sha256.New, []byte(siteSecret))
	generator.Write([]byte(secretToken))
	result := generator.Sum(nil)
	return hex.EncodeToString(result)
}
