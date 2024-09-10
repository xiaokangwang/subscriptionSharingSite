package keyValueStorage

import "context"

type ScopedPersistentStorage interface {
	ScopedPersistentStorageEngine()
	Put(ctx context.Context, key string, value []byte) error      // OK
	Get(ctx context.Context, key string) ([]byte, error)          // OK
	List(ctx context.Context, keyPrefix []byte) ([]string, error) // OK
	Clear(ctx context.Context)
	NarrowScope(ctx context.Context, key string) (ScopedPersistentStorage, error)
	DropScope(ctx context.Context, key string) error
}
