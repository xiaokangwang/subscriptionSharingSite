package model

import (
	"context"
	"github.com/xiaokangwang/subscriptionSharingSite/keyValueStorage"
	"log"
)

func ListAllEntriesByGroup(kv keyValueStorage.ScopedPersistentStorage, group string) ([]string, error) {
	keys, err := kv.List(context.TODO(), []byte(group+"/"))
	if err != nil {
		return nil, err
	}
	entries := make([]string, 0)
	for i := 0; i < len(keys); i++ {
		result, err := kv.Get(context.TODO(), keys[i])
		if err != nil {
			log.Println(err)
			continue
		}
		entries = append(entries, string(result))
	}
	return entries, nil
}
