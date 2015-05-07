package kvs

import (
	"fmt"
	"github.com/contiv/go-etcd/etcd"
	"log"
)

var etcdClient *etcd.Client

type Etcd struct {
	Machines []string `json:"machines"`
}

type config struct {
	Etcd Etcd `json:"etcd"`
}

var etcdConfig config

type etcdKVS struct {
}

func etcdStart(data []byte) (kvsInst KVS, err error) {
	log.Println("********************** Initializing the Etcd Client *********************\n")

	err = json.Unmarshal(data, etcdConfig)
	if err != nil {
		log.Println("kvs.etcd: Unable to parse configuration data =", string(data))
		return
	}

	if etcdClient = etcd.NewClient(etcdConfi.Etcd.Machines); etcdClient == nil {
		log.Println("kvs.etcd: Unable to Create Etcd Client")
		err = fmt.Errorf("Unable to create a etcd Client. Exiting")
		return
	}
	return
}

func (kv etcdKVS) Create(key string, value []byte) (err error) {
	if _, err = etcdClient.Create(key, string(value), 0); err != nil {
		log.Panicln("Error =", err)
		return
	}
	return
}
