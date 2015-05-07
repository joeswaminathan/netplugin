/********************************************************************
 *
 *      File:
 *
 *      Description:
 *       This file contains generic Key Value Store functionality
 *
 *      Copyright (c) 2015 by Cisco Systems, Inc.
 *
 *       ALL RIGHTS RESERVED. THESE SOURCE FILES ARE THE SOLE PROPERTY
 *       OF CISCO SYTEMS, Inc. AND CONTAIN CONFIDENTIAL  AND PROPRIETARY
 *       INFORMATION.  REPRODUCTION OR DUPLICATION BY ANY MEANS OF ANY
 *       PORTION OF THIS SOFTWARE WITHOUT PRIOR WRITTEN CONSENT OF
 *       CISCO SYSTEMS, Inc. IS STRICTLY PROHIBITED.
 *
 *      Author: Joseph Swaminathan
 *
 ********************************************************************/

package kvs

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

/*
 * KVS package can use Etcd or Consul or any other key value store in the future
 * KVS assumes each of these tools to provide the following abstraction
 *
 * etcd.go provides the abstraction for etcd.
 *
 */
type KVS interface {
	Exist(key string) bool
	Create(key string, data []byte) (err error)
	Read(key string) (data []byte, err error)
	Write(key string, data []byte) (err error)
	Update(key string, data []byte) (err error)
	Delete(key string) (err error)
}

var kvsInst KVS

var (
	ErrInvalidConfig  = errors.New("Invalid Configuration")
	ErrNotInitialized = errors.New("Not intialized")
)

const (
	ETCD = "etcd"
	FAKE = "fake"
)

type Driver struct {
	State string `json:"state"`
}

type config struct {
	Driver Driver `json:"drivers"`
}

var kvsConfig config

func Init(data []byte) (err error) {

	log.Println("********************** Initializing the KVS Library *********************\n")

	err = json.Unmarshal(data, kvsConfig)
	if err != nil {
		return
	}

	if kvsConfig.Driver.State == ETCD {
		if kvsInst, err = etcdStart(data); err != nil {
			log.Panicln("Init : Unable to create a etcd Client. Exiting")
			return fmt.Errorf("Init : Unable to create a etcd Client. Exiting")
		}
	} else if kvsConfig.Driver.State == FAKE {
		if kvsInst, err = fakeStart(data); err != nil {
			log.Panicln("Init : Unable to create a etcd Client. Exiting")
			return fmt.Errorf("Init : Unable to create a etcd Client. Exiting")
		}
	} else {
		return ErrInvalidConfig
	}

	return nil
}

func Stat(key string) bool {
	if kvsInst == nil {
		return false
	}
	return kvsInst.Exist(key)
}

func Read(key string) (data []byte, err error) {
	if kvsInst == nil {
		err = ErrNotInitialized
		return
	}

	if data, err = kvsInst.Read(key); err != nil {
		log.Panicln("KVS:Read Error Reading key = ", key)
		return
	}
	return
}

func Write(key string, data []byte) (err error) {
	if kvsInst == nil {
		return ErrNotInitialized
	}
	log.Println("KVS:Write : Path =", key, " data =", string(data))
	return kvsInst.Write(key, data)
}

func Update(key string, data []byte) (err error) {
	if kvsInst == nil {
		return ErrNotInitialized
	}
	log.Println("KVS:Update : Path =", key, " data =", string(data))
	return kvsInst.Update(key, data)
}

func Create(key string, data []byte) (err error) {
	if kvsInst == nil {
		return ErrNotInitialized
	}
	log.Println("KVS:Create : Path =", key, " data =", string(data))
	return kvsInst.Create(key, data)
}

func Delete(key string) (err error) {
	if kvsInst == nil {
		return ErrNotInitialized
	}
	log.Println("KVS:Delete : Path =", key)
	return kvsInst.Delete(key)
}
