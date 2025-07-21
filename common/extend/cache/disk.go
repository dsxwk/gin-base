package cache

import (
	"errors"
	"fmt"
	"github.com/dgraph-io/badger/v4"
	"log"
	"time"
)

// DiskCache 磁盘缓存
type DiskCache struct {
	db *badger.DB
}

func NewDiskCache(dir string) (*DiskCache, error) {
	opts := badger.DefaultOptions(dir)
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(fmt.Sprintf("init disk cache failed: %s", err.Error()))
		return nil, err
	}
	return &DiskCache{db: db}, nil
}

func (b *DiskCache) SetCache(key string, value interface{}, expire time.Duration) error {
	valBytes, ok := value.([]byte)
	if !ok {
		valBytes = []byte(fmt.Sprintf("%v", value))
	}
	err := b.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(key), valBytes)
		if expire > 0 {
			e = e.WithTTL(expire)
		}
		return txn.SetEntry(e)
	})
	return err
}

func (b *DiskCache) GetCache(key string) (interface{}, bool) {
	var val []byte
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		val, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		return nil, false
	}
	return val, true
}

func (b *DiskCache) DeleteCache(key string) error {
	return b.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

func (b *DiskCache) GetCacheExpire(key string) (interface{}, time.Time, bool, error) {
	var (
		val        []byte
		expireTime time.Time
	)

	// Badger 不直接支持获取剩余 TTL，只能判断是否存在
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		val, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}
		ttl := item.ExpiresAt()
		if ttl > 0 {
			expireTime = time.Unix(int64(ttl), 0)
		}
		return nil
	})
	if err != nil {
		return nil, time.Time{}, false, errors.New("cache key not found")
	}
	return val, expireTime, true, nil
}
