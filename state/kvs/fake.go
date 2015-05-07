package kvs

import (
	"fmt"
	"github.com/contiv/go-etcd/etcd"
	"log"
)

type Fake struct {
	Machines []string `json:"machines"`
}

type config struct {
	Fake Fake `json:"fake"`
}

var fakeConfig config

type fakeKVS struct {
}

var db map[string][]byte

func fakeStart(data []byte) (kvsInst KVS, err error) {
	log.Println("********************** Initializing the Etcd Client *********************\n")

	err = json.Unmarshal(data, fakeConfig)
	if err != nil {
		log.Println("kvs.fake: Unable to parse configuration data =", string(data))
		return
	}
	log.Println("kvs.fake: Fake KVS driver client created")
	return
}

func (kv etcdKVS) Create(key string, value []byte) (err error) {
	db[key] = value
	log.Printf("kvs.fake: Created key (%s) with value (%s)\n", key, string(value))
	return
}

func (kv etcdKVS) Read(key string) (value []byte, err error) {
	data, ok := db[key]
	if ok {
		value = data
	} else {
		err = fmt.Errorf("Not found\n")
	}
	log.Printf("kvs.fake: Read key (%s) \n", key)
	return
}

func (kv etcdKVS) Write(key string, value []byte) (err error) {
	db[key] = value
	log.Printf("kvs.fake: Wrote key (%s) with value (%s)\n", key, string(value))
	return
}

func (kv etcdKVS) Update(key string, value []byte) (err error) {
	data, ok := db[key]
	if ok {
		db[key] = value
	} else {
		err = fmt.Errorf("Not found\n")
	}
	log.Printf("kvs.fake: Updated key (%s) with value (%s)\n", key, string(value))
	return
}

func (kv etcdKVS) Delete(key string) (err error) {
	data, ok := db[key]
	if ok {
		delte(db, key)
	} else {
		err = fmt.Errorf("Not found\n")
	}
	log.Printf("kvs.fake: Read key (%s) \n", key)
	return
}
