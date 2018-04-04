package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/go-redis/redis"
)

const (
	RedisKey = "redcount"
)

type Redcount struct {
	conf *Conf
	srv *http.Server
	redis *redis.Client
}

func (app *Redcount) HandleSlash(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch verb := r.Method; verb {
	case "GET":
		val, err := app.redis.Get(RedisKey).Result()
		if err != nil {
			log.Printf("could not read the key [%s]: %s", RedisKey, err)
		}
		fmt.Fprintf(w, "%s", val)
	case "POST":
		val, err := app.redis.Incr(RedisKey).Result()
		if err != nil {
			log.Printf("could not incr the key [%s]: %s", RedisKey, err)
		}
		fmt.Fprintf(w, "%d", val)
	default:
		http.NotFound(w, r)
	}
}
