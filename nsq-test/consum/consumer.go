package main

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
)

type user struct {
	Name string `json:"name"`
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("test", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		var user1 user
		log.Print(message.Body)
		err := json.Unmarshal(message.Body, &user1)
		if err != nil {
			log.Panic(err)
		}
		log.Printf("Got a message: %s", user1.Name)
		wg.Done()
		return nil
	}))
	err := q.ConnectToNSQD("10.0.0.99:4150")
	if err != nil {
		log.Panic(err)
	}
	wg.Wait()
}
