package main

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	Redis struct {
		Host string
		Port uint16
		Db uint8
	}
}

func (c *Conf) readFromFile(filename string) (*Conf) {

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	log.Printf("Redis host: %s\n", c.Redis.Host)
	log.Printf("Redis port: %d\n", c.Redis.Port)
	log.Printf("Redis db: %d\n", c.Redis.Db)

	return c
}

