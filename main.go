package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	cwd, _ := os.Getwd()
	log.Printf("Starting Redcount from %s ...", cwd)
	app := &Redcount{}
	app.conf = &Conf{}
	app.conf.readFromFile("redcount.yaml")
	app.redis = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", app.conf.Redis.Host, app.conf.Redis.Port),
		DB: 0,
	})

	val, err := app.redis.Exists(RedisKey).Result()
	if err != nil {
		log.Printf("error checking key [%s]: %s", RedisKey, err)
		val = 0
	}
	if val != 1 {
		log.Printf("[%s] does not exist, create it with value = 0", RedisKey)
		_, err = app.redis.Set(RedisKey, 0, 0).Result()
		if err != nil {
			log.Printf("could not set default value to key [%s]: %s", RedisKey, err)
		}
	}

	app.srv = &http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/", app.HandleSlash)
	log.Fatal(app.srv.ListenAndServe())
}
