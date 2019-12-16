package main

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"math/rand"
	"time"
)

type user struct {
	Name string `json:"name"`
}

func main() {
	config := nsq.NewConfig()
	w, err := nsq.NewProducer("10.0.0.99:4150", config)
	if err != nil {
		log.Panic(err)
	}
	chars := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for {
		buf := make([]byte, 4)
		for i := 0; i < 4; i++ {
			buf[i] = chars[rand.Intn(len(chars))]
		}
		user1 := user{Name: string(buf)}
		fmt.Println(user1)
		temp := make([]byte, 2)
		temp, err := json.Marshal(&user1)
		log.Println(string(temp))
		if err != nil {
			log.Panic(err)
		}
		err = w.Publish("test", temp)
		if err != nil {
			log.Panic(err)
		}
		time.Sleep(1 * time.Second)
	}
	w.Stop()
}
