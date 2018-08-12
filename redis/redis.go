package redis

import (
	"fmt"

	re "github.com/go-redis/redis"
)

var rd *re.Client

func init() {
	rd = re.NewClient(&re.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func IsDuplicate(userID1, userID2 string) bool {
	return "true" == rd.Get(key(userID1, userID2)).Val()
}

func Duplicate(userID1, userID2 string) {
	rd.Set(key(userID1, userID2), "true", 0)
	rd.Set(key(userID2, userID1), "true", 0)
}

func key(userID1, userID2 string) string {
	return fmt.Sprintf("%s,%s", userID1, userID2)
}
