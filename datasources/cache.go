package datasources

import "time"

type Cache interface {
	Set(key string, value interface{}, expiration time.Duration) error
	KeyExists(key string) (bool, error)
}
