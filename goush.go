package main

import "github.com/go-martini/martini"
import "github.com/garyburd/redigo/redis"

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	m := martini.Classic()
	m.Get("/", func() string {
		url, _ := redis.String(c.Do("GET", "u:1"))
		return "Hello rafał! " + url
	})
	m.Run()
}
