package main

import "github.com/garyburd/redigo/redis"
import "fmt"

var (
	pool *redis.Pool
)

func InitDb() {
	pool = redis.NewPool(func() (redis.Conn, error) { return redis.Dial("tcp", ":6379") }, 10000)
}

func GetPool() *redis.Pool { return pool }

func OpenConnection() redis.Conn {
	return pool.Get()
}

func GetUrl(uid string) (string, error) {
	c := OpenConnection()
	defer c.Close()

	return redis.String(c.Do("GET", fmt.Sprintf("u:%v", uid)))
}

func VisitUrl(uid string, visitData []byte) string {
	c := OpenConnection()
	defer c.Close()

	url, err := GetUrl(uid)
	if err != nil {
		panic(err)
	}
	c.Do("RPUSH", fmt.Sprintf("v:%v", uid), visitData)
	c.Do("INCR", "stats:visit_count")
	return url
}
