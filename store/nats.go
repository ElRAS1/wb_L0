package store

import (
	"encoding/json"
	"log"

	"github.com/nats-io/stan.go"
)

func (s *Store) NatsSubscribe(ns stan.Conn) {
	_, err := ns.Subscribe("wb", func(msg *stan.Msg) {
		data := Order{}
		err := json.Unmarshal(msg.Data, &data)

		if err != nil {
			log.Fatalln(err)
		}
		if err != nil {
			log.Println(err)
			return
		}
	}, stan.DurableName("wb"))

	if err != nil {
		log.Fatalln("Failed to subscribe:", err)
	}
}

func (s *Store) NatsPublish(ns stan.Conn) {
	data, err := s.jsonData()

	if err != nil {
		log.Fatalln(err)
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error serializing Order to JSON: %v", err)
	}

	ok := s.checkData(data)

	if ok {
		return
	}

	err = ns.Publish("wb", []byte(jsonData))
	if err != nil {
		log.Fatalf("Error while trying to send msg: %v", err)
	}

}
