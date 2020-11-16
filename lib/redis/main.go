package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v7"
)

type Car struct {
	Name string
}

func (c Car) String() string {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func main() {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	set(client)
	get(client)
}

func set(client *redis.Client) {
	pipe := client.Pipeline()
	car1 := Car{Name: "SUV"}
	car2 := Car{Name: "Toyota"}
	pipe.Set("suv", car1.String(), time.Hour)
	pipe.Set("toyota", car2.String(), time.Hour)
	res, err := pipe.Exec()
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range res {
		log.Print(r)
	}
}

func get(client *redis.Client) {
	pipe := client.Pipeline()
	pipe.Get("suv")
	pipe.Get("toyota")
	res, err := pipe.Exec()
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range res {
		log.Print(r.String())
	}
}
