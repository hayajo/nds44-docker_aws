package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/garyburd/redigo/redis"
)

func main() {
	port := os.Getenv("COUNTER_PORT")
	if port == "" {
		port = "8080"
	}
	redisKey := os.Getenv("COUNTER_REDIS_KEY")
	if redisKey == "" {
		redisKey = "counter"
	}
	redisAddr := os.Getenv("COUNTER_REDIS_ADDRESS")
	if redisAddr == "" {
		redisAddr = "redis:6379"
	}
	c, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		r, err := c.Do("INCR", redisKey)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		rw.Write([]byte(fmt.Sprint(r)))
	})
	http.ListenAndServe(":"+port, nil)
}
