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

// netmaster  - implements the network intent translation to plugin
// events; uses state distribution to achieve intent realization
// netmaster runs as a logically centralized unit on in the cluster

package netmaster

import (
	"encoding/json"
	"fmt"
	"github.com/contiv/netplugin/core"
	"github.com/contiv/netplugin/drivers"
	"state"
)

const (
	BASE_PATH          = drivers.BASE_PATH + "master/"
	CFG_PATH           = BASE_PATH + "config/"
	NW_CFG_PATH_PREFIX = CFG_PATH + "nets/"
	NW_CFG_PATH        = NW_CFG_PATH_PREFIX + "%s"
)

type MasterNwConfig struct {
	state.CommonState
	Tenant     string `json:"tenant"`
	PktTagType string `json:"pktTagType"`
	PktTag     string `json:"pktTag"`
	SubnetIp   string `json:"subnetIp"`
	SubnetLen  uint   `json:"subnetLen"`
	DefaultGw  string `json:"defaultGw"`
}

func (s MasterNwConfig) Key() string {
	return NW_CFG_PATH_PREFIX + s.Id
}

func readNwCfg(id string) (st core.State, mstrNwCfg *MasterNwConfig, err error) {
	st := state.NewState{Data: MasterNwConfig{Id: id}}
	err := st.Read()
	if err != nil {
		return err
	}
	mstrNwCfg := &((st.Data()).(MasterNwConfig))
	return
}

func newNwCfgFromId(id string) (st core.State, mstrNwCfg *MasterNwConfig, err error) {
	st := state.NewState{Data: MasterNwConfig{Id: id}}
	mstrNwCfg := &((st.Data()).(MasterNwConfig))
	return
}

func newNwCfgFromData(data []byte) (st core.State, mstrNwCfg *MasterNwConfig, err error) {
	st := state.NewState{Data: MasterNwConfig{}}
	if err := st.Set([]byte(value)); err != nil {
		return err
	}
	mstrNwCfg := &((st.Data()).(MasterNwConfig))
	return
}
