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
	"state"
)

// implements the State interface for an endpoint implemented using
// vlans with ovs. The st is stored as Json objects.

type OvsCfgEndpointState struct {
	state.CommonState
	NetId      string `json:"netId"`
	ContName   string `json:"contName"`
	AttachUUID string `json:"attachUUID"`
	IpAddress  string `json:"ipAddress"`
	HomingHost string `json:"homingHost"`
	IntfName   string `json:"intfName"`
	VtepIp     string `json:'vtepIP"`
}

func (s OvsCfgEndpointState) Key() string {
	return EP_CFG_PATH_PREFIX + s.Id
}

func readEpCfg(id string) (st core.State, cfgNw *OvsCfgEndpointState, err error) {
	st := state.NewState{Data: OvsCfgEndpointState{Id: id}}
	err := st.Read()
	if err != nil {
		return err
	}
	cfgNw := &((st.Data()).(OvsCfgEndpointState))
	return
}

func newEpCfgFromId(Id string) (st core.State, cfgNw *OvsCfgEndpointState, err error) {
	st := state.NewState{Data: OvsCfgEndpointState{Id: id}}
	cfgNw := &((st.Data()).(OvsCfgEndpointState))
	return
}

func newEpCfgFromData(data []byte) (st core.State, cfgNw *OvsCfgEndpointState, err error) {
	st := state.NewState{Data: OvsCfgEndpointState{}}
	if err := st.Set([]byte(value)); err != nil {
		return err
	}
	cfgNw := &((st.Data()).(OvsCfgEndpointState))
	return
}

type OvsOperEndpointState struct {
	core.CommonState
	NetId      string `json:"netId"`
	ContName   string `json:"contName"`
	AttachUUID string `json:"attachUUID"`
	IpAddress  string `json:"ipAddress"`
	PortName   string `json:"portName"`
	HomingHost string `json:"homingHost"`
	IntfName   string `json:"intfName"`
	VtepIp     string `json:'vtepIP"`
}

func (s OvsOperEndpointState) Key() string {
	return EP_OPER_PATH_PREFIX + s.Id
}

func readEpOper(id string) (st core.State, cfgNw *OvsOperEndpointState, err error) {
	st := state.NewState{Data: OvsOperEndpointState{Id: id}}
	err := st.Read()
	if err != nil {
		return err
	}
	cfgNw := &((st.Data()).(OvsOperEndpointState))
	return
}

func newEpOperFromId(id string) (st core.State, cfgNw *OvsOperEndpointState, err error) {
	st := state.NewState{Data: OvsOperEndpointState{Id: id}}
	cfgNw := &((st.Data()).(OvsOperEndpointState))
	return
}

func newEpOperFromData(data []byte) (st core.State, cfgNw *OvsOperEndpointState, err error) {
	st := state.NewState{Data: OvsOperEndpointState{}}
	if err := st.Set([]byte(value)); err != nil {
		return err
	}
	cfgNw := &((st.Data()).(OvsOperEndpointState))
	return
}
