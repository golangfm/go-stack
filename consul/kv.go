package main

import (
	"log"

	"github.com/hashicorp/consul/api"
)

func main() {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	kv := client.KV()
	// pair := &api.KVPair{Key: "adn/key1", Value: []byte("val1")}
	// if _, err := kv.Put(pair, nil); err != nil {
	// 	log.Fatal(err)
	// }

	key := "adn/config"
	if p, _, err := kv.Get(key, nil); err != nil {
		log.Fatal(err)
	} else {
		log.Println(p.Key, "==>", string(p.Value))
	}
}
