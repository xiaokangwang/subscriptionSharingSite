package rediskv

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/xiaokangwang/subscriptionSharingSite/keyValueStorage"
)

func NewKVFromURL(url string) keyValueStorage.ScopedPersistentStorage {
	opt, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)
	return NewKV(client)

}

func NewKV(client *redis.Client) keyValueStorage.ScopedPersistentStorage {
	return &KV{client: client}
}

type KV struct {
	client *redis.Client
}

func (k *KV) ScopedPersistentStorageEngine() {
}

func (k *KV) Put(ctx context.Context, key string, value []byte) error {
	result := k.client.Set(ctx, key, value, 0)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (k *KV) Get(ctx context.Context, key string) ([]byte, error) {
	result := k.client.Get(ctx, key)
	if result.Err() != nil {
		return nil, result.Err()
	}
	return []byte(result.Val()), nil
}

func (k *KV) List(ctx context.Context, keyPrefix []byte) ([]string, error) {
	result := k.client.Keys(ctx, string(keyPrefix)+"*")
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result.Val(), nil
}

func (k *KV) Clear(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (k *KV) NarrowScope(ctx context.Context, key string) (keyValueStorage.ScopedPersistentStorage, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KV) DropScope(ctx context.Context, key string) error {
	//TODO implement me
	panic("implement me")
}
