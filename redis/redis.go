package redis

import (
	"fmt"

	re "github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
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
func (p *Cache) IsDuplicate(userID1, userID2 int64) bool {
	val, err := p.rd.Get(key(userID1, userID2)).Result()
	if err != nil {
		log.WithFields(log.Fields{
			"userID1": userID1,
			"userID2": userID2,
		}).WithError(err).Error("read to cache")
	}
	return val == "true"
}

// Duplicate remebers to redis, that two users are actually duplicates.
func (p *Cache) Duplicate(userID1, userID2 int64) {
	if err := p.rd.Set(key(userID1, userID2), "true", 0).Err(); err != nil {
		log.WithFields(log.Fields{
			"userID1": userID1,
			"userID2": userID2,
		}).WithError(err).Error("write to cache")
	}
	if err := p.rd.Set(key(userID2, userID1), "true", 0).Err(); err != nil {
		log.WithFields(log.Fields{
			"userID1": userID1,
			"userID2": userID2,
		}).WithError(err).Error("write to cache")
	}
}

func key(userID1, userID2 int64) string {
	return fmt.Sprintf("%d,%d", userID1, userID2)
}
