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

package drivers

import (
	"encoding/json"
	"fmt"
	"github.com/contiv/netplugin/core"
	"github.com/jainvipin/bitset"
	"state"
)

// implements the State interface for a network implemented using
// vlans with ovs. The st is stored as Json objects.

type OvsCfgNetworkState struct {
	Id         string        `json:"id"`
	Tenant     string        `json:"tenant"`
	PktTagType string        `json:"pktTagType"`
	PktTag     int           `json:"pktTag"`
	ExtPktTag  int           `json:"extPktTag"`
	SubnetIp   string        `json:"subnetIp"`
	SubnetLen  uint          `json:"subnetLen"`
	DefaultGw  string        `json:"defaultGw"`
	EpCount    int           `json:"epCount"`
	IpAllocMap bitset.BitSet `json:"ipAllocMap"`
}

func (s OvsCfgNetworkState) Key() string {
	return NW_CFG_PATH_PREFIX + s.Id
}

func readNwCfg(id string) (st state.State, nwCfg *OvsCfgNetworkState, err error) {
	st := state.NewState{Data: OvsCfgNetworkState{Id: id}}
	err := st.Read()
	if err != nil {
		return err
	}
	nwCfg := &((st.Data()).(OvsCfgNetworkState))
	return
}

func newNwCfgFromId(id string) (st state.State, nwCfg *OvsCfgNetworkState, err error) {
	st := state.NewState{Data: OvsCfgNetworkState{Id: id}}
	nwCfg := &((st.Data()).(OvsCfgNetworkState))
	return
}

func newNwCfgFromData(data []byte) (st state.State, nwCfg *OvsCfgNetworkState, err error) {
	st := state.NewState{Data: OvsCfgNetworkState{}}
	if err := st.Set([]byte(value)); err != nil {
		return err
	}
	nwCfg := &((st.Data()).(OvsCfgNetworkState))
	return
}
