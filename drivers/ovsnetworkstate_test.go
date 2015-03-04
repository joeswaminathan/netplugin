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
	"fmt"
	"testing"

	"github.com/contiv/netplugin/core"
)

const (
	testNwId  = "testNw"
	nwCfgKey  = NW_CFG_PATH_PREFIX + testNwId
	nwOperKey = NW_OPER_PATH_PREFIX + testNwId
)

var nwStateDriver *testNwStateDriver = &testNwStateDriver{}

type testNwStateDriver struct {
}

func (d *testNwStateDriver) Init(config *core.Config) error {
	return &core.Error{Desc: "Shouldn't be called!"}
}

func (d *testNwStateDriver) Deinit() {
}

func (d *testNwStateDriver) Write(key string, value []byte) error {
	return &core.Error{Desc: "Shouldn't be called!"}
}

func (d *testNwStateDriver) Read(key string) ([]byte, error) {
	return []byte{}, &core.Error{Desc: "Shouldn't be called!"}
}

func (d *testNwStateDriver) ReadRecursive(baseKey string) ([]string, error) {
	return []string{}, &core.Error{Desc: "Shouldn't be called!"}
}

func (d *testNwStateDriver) validateKey(key string) error {
	if key != nwCfgKey && key != nwOperKey {
		return &core.Error{Desc: fmt.Sprintf("Unexpected key. recvd: %s expected: %s or %s ",
			key, nwCfgKey, nwOperKey)}
	} else {
		return nil
	}
}

func (d *testNwStateDriver) ClearState(key string) error {
	return d.validateKey(key)
}

func (d *testNwStateDriver) SafeClearState(key string, prevVal core.State,
	marshal func(interface{}) ([]byte, error)) error {
	return d.validateKey(key)
}

func (d *testNwStateDriver) ReadState(key string, value core.State,
	unmarshal func([]byte, interface{}) error) error {
	return d.validateKey(key)
}

func (d *testNwStateDriver) WriteState(key string, value core.State,
	marshal func(interface{}) ([]byte, error)) error {
	return d.validateKey(key)
}

func (d *testNwStateDriver) SafeWriteState(key string, value core.State,
	marshal func(interface{}) ([]byte, error),
	prevVal func(core.State) core.State,
	nextVal func(core.State) core.State) error {
	return d.validateKey(key)
}

func TestOvsCfgNetworkStateRead(t *testing.T) {
	epCfg := &OvsCfgNetworkState{StateDriver: nwStateDriver}

	err := epCfg.Read(testNwId)
	if err != nil {
		t.Fatalf("read config state failed. Error: %s", err)
	}
}

func TestOvsCfgNetworkStateWrite(t *testing.T) {
	epCfg := &OvsCfgNetworkState{StateDriver: nwStateDriver, Id: testNwId}

	err := epCfg.Write()
	if err != nil {
		t.Fatalf("write config state failed. Error: %s", err)
	}
}

func TestOvsCfgNetworkStateClear(t *testing.T) {
	epCfg := &OvsCfgNetworkState{StateDriver: nwStateDriver, Id: testNwId}

	err := epCfg.Clear()
	if err != nil {
		t.Fatalf("clear config state failed. Error: %s", err)
	}
}

func TestOvsOperNetworkStateRead(t *testing.T) {
	epOper := &OvsOperNetworkState{StateDriver: nwStateDriver}

	err := epOper.Read(testNwId)
	if err != nil {
		t.Fatalf("read oper state failed. Error: %s", err)
	}
}

func TestOvsOperNetworkStateWrite(t *testing.T) {
	epOper := &OvsOperNetworkState{StateDriver: nwStateDriver, Id: testNwId}

	err := epOper.Write()
	if err != nil {
		t.Fatalf("write oper state failed. Error: %s", err)
	}
}

func TestOvsOperNetworkStateSafeWrite(t *testing.T) {
	epOper := &OvsOperNetworkState{StateDriver: nwStateDriver, Id: testNwId}

	err := epOper.SafeWrite(nil, nil)
	if err != nil {
		t.Fatalf("write oper state failed. Error: %s", err)
	}
}

func TestOvsOperNetworkStateClear(t *testing.T) {
	epOper := &OvsOperNetworkState{StateDriver: nwStateDriver, Id: testNwId}

	err := epOper.Clear()
	if err != nil {
		t.Fatalf("clear oper state failed. Error: %s", err)
	}
}

func TestOvsOperNetworkStateSafeClear(t *testing.T) {
	epOper := &OvsOperNetworkState{StateDriver: nwStateDriver, Id: testNwId}

	err := epOper.SafeClear()
	if err != nil {
		t.Fatalf("clear oper state failed. Error: %s", err)
	}
}
