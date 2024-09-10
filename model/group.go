package model

import (
	"github.com/xiaokangwang/subscriptionSharingSite/keyValueStorage"
	"log"
)

func ListAllEntriesByGroup(kv keyValueStorage.ScopedPersistentStorage, group string) ([]string, error) {
	keys, err := kv.List(nil, []byte(group+"/"))
	if err != nil {
		return nil, err
	}
	entries := make([]string, 0)
	for i := 0; i < len(keys); i++ {
		result, err := kv.Get(nil, keys[i])
		if err != nil {
			log.Println(err)
			continue
		}
		entries = append(entries, string(result))
	}
	return entries, nil
}
