package model

import (
	"context"
	"github.com/xiaokangwang/subscriptionSharingSite/keyValueStorage"
)

type ProxyServer struct {
	Group       string
	PublicToken string
	EntryName   string
}

func (s *ProxyServer) GetDatabaseKey() string {
	return s.Group + "/" + s.PublicToken + "/" + s.EntryName
}

func (s *ProxyServer) GetContentFromKV(
	storage keyValueStorage.ScopedPersistentStorage) ([]byte, error) {
	return storage.Get(context.TODO(), s.GetDatabaseKey())
}

func (s *ProxyServer) PutContentToKV(
	storage keyValueStorage.ScopedPersistentStorage, content []byte) error {
	return storage.Put(context.TODO(), s.GetDatabaseKey(), content)
}
