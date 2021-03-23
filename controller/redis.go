package controller

import "time"

type Redis interface {
	Set(key, value string, expire time.Duration) error
	Get(key string) (string, error)
	Inc(key string) error
}
