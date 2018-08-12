package redis

import (
	"fmt"

	re "github.com/go-redis/redis"
)

// Opts contains settings for redis.
type Opts struct {
	Addr string
}

// NewCache returns struct redis connection.
func NewCache(opt *Opts) *Cache {
	c := &Cache{re.NewClient(&re.Options{
		Addr:     opt.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})}
	c.rd.Ping() //establish connection

	return c
}

// Cache contains methods on top of redis cache.
type Cache struct {
	rd *re.Client
}

// IsDuplicate checks if we remembered previously in cache
// that two users are duplicate.
func (p *Cache) IsDuplicate(userID1, userID2 string) bool {
	return "true" == p.rd.Get(key(userID1, userID2)).Val()
}

// Duplicate remebers to redis, that two users are actually duplicates.
func (p *Cache) Duplicate(userID1, userID2 string) {
	p.rd.Set(key(userID1, userID2), "true", 0)
	p.rd.Set(key(userID2, userID1), "true", 0)
}

func key(userID1, userID2 string) string {
	return fmt.Sprintf("%s,%s", userID1, userID2)
}
