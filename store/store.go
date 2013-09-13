// Package urls creates and stores URLs in a Redis database.
package urls

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	maxIdleConnections    = 3
	idleConnectionTimeout = 45 * time.Second
	keyFmt                = "gourls:key:%s"
	urlKeyFmt             = "gourls:url:%s"
)

var (
	ErrConflict = errors.New("Conflict: Key exists")
)

// Type URLStore defines a store for URLs.
type URLStore struct {
	pool   *redis.Pool
	random *rand.Rand
}

// NewURLStore creates a new URL store.
func NewURLStore(redisAddr string) *URLStore {
	pool := &redis.Pool{
		MaxIdle:     maxIdleConnections,
		IdleTimeout: idleConnectionTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisAddr)
		},
	}
	return &URLStore{
		pool:   pool,
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// New stores a URL with a new, randomly-generated key.
func (s *URLStore) New(url string) (string, error) {
	c := s.pool.Get()
	defer c.Close()

	// Check if the URL already exists
	if keys, err := redis.Strings(c.Do("LRANGE", fullKeyForUrl(url), 0, -1)); check(err) && len(keys) > 0 {
		return strings.Join(keys, ","), nil
	}

	// URL doesn't exist
	key := s.createKey()
	c.Send("MULTI")
	c.Send("SET", fullKey(key), url)
	c.Send("LPUSH", fullKeyForUrl(url), key)
	if _, err := c.Do("EXEC"); err != nil {
		return "", err
	}

	return key, nil
}

// Set stores a URL with the given key.
func (s *URLStore) Set(key, url string) error {
	c := s.pool.Get()
	defer c.Close()

	// Check if key already exists
	if exists, err := redis.Bool(c.Do("EXISTS", fullKey(key))); err != nil {
		return err
	} else if exists {
		return ErrConflict
	}

	// Key doesn't exist. Set it now.
	c.Send("MULTI")
	c.Send("SET", fullKey(key), url)
	c.Send("LPUSH", fullKeyForUrl(url), key)
	if _, err := c.Do("EXEC"); err != nil {
		return err
	}

	return nil
}

// Get retrieves a URL for the given key.
func (s *URLStore) Get(key string) (string, error) {
	c := s.pool.Get()
	defer c.Close()

	return redis.String(c.Do("GET", fullKey(key)))
}

func (s *URLStore) createKey() string {
	return strconv.FormatInt(s.random.Int63(), 36)
}

// Utilities

func fullKey(key string) string {
	return fmt.Sprintf(keyFmt, key)
}

func fullKeyForUrl(url string) string {
	return fmt.Sprintf(urlKeyFmt, url)
}

func check(err error) bool {
	return err == nil || err == redis.ErrNil
}
