/***
Copyright 2014 Cisco Systems Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package netmaster

import (
	"encoding/json"
	"fmt"
	"github.com/contiv/netplugin/core"
	"state"
)

const (
	HOST_CFG_PATH_PREFIX = CFG_PATH + "hosts/"
	HOST_CFG_PATH        = HOST_CFG_PATH_PREFIX + "%s"
)

type MasterHostConfig struct {
	state.CommonState
	Name   string `json:"name"`
	Intf   string `json:"intf"`
	VtepIp string `json:"vtepIp"`
	NetId  string `json:"netId"`
}

func (s MasterHostConfig) Key() string {
	return NW_CFG_PATH_PREFIX + s.Id
}

func readHostCfg(id string) (st core.State, mstrHostCfg *MasterHostConfig, err error) {
	st := state.NewState{Data: MasterHostConfig{Id: id}}
	err := st.Read()
	if err != nil {
		return err
	}
	mstrHostCfg := &((st.Data()).(MasterHostConfig))
	return
}

func newHostCfgFromId(id string) (st core.State, mstrHostCfg *MasterHostConfig, err error) {
	st := state.NewState{Data: MasterHostConfig{Id: id}}
	mstrHostCfg := &((st.Data()).(MasterHostConfig))
	return
}

func newHostCfgFromData(data []byte) (st core.State, mstrHostCfg *MasterHostConfig, err error) {
	st := state.NewState{Data: MasterHostConfig{}}
	if err := st.Set([]byte(value)); err != nil {
		return err
	}
	mstrHostCfg := &((st.Data()).(MasterHostConfig))
	return
}
