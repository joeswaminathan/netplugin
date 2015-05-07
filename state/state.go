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

package state

import (
	"encoding/json"
	"errors"
	"fmt"
	//	"github.com/contiv/netplugin/state/kvs"
	//	"state/kvs"
)

type State interface {
	Read() error
	Write() error
	Clear() error
	Set(data []byte) error
	Data() StateData
}

type States interface {
	ReadAll() error
	Add(state *State) error
	Del(state *State) error
}

type StateData interface {
	Key() string
}

type PathFn func() string

type MarshalOper interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

type CommonMarshalOper struct {
}

func (state CommonMarshalOper) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (state CommonMarshalOper) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

type CommonStateModel struct {
	data  StateData
	mOper MarshalOper
}

func New(dataType reflect.Type) (st *State, err error) {
	if data := reflect.New(dataType); data == nil {
		return fmt.Errorf("Not enough memory")
	}
	st = &CommonStateModel{data: data, mOper: CommonMarshalOper{}}
	return
}

func NewState(data StateData, mOper MarshalOper) (state State) {
	if data == nil {
		return nil
	}
	if mOper == nil {
		state = &CommonStateModel{data: data, mOper: CommonMarshalOper{}}
	} else {
		state = &CommonStateModel{data: data, mOper: mOper}
	}
	return state
}

func (state CommonStateModel) Read() error {
	key := state.data.Key()
	fmt.Println("Key = ", key)
	/*
		data, err := kvs.Read(key)
		if err != nil {
			return errors.New("Unable to read key (%s)", key)
		}
		if err := state.mOper.Unmarshal(data, state.data); err != nil {
			return erros.New("Unable to parse data for key (%s)", key)
		}
	*/
	return nil
}

func (state CommonStateModel) Write() error {
	fmt.Println("CommonStateModel =", state)
	data, err := state.mOper.Marshal(state.data)
	if err != nil {
		fmt.Println("Error =", err)
		return errors.New("Unable to parse data")
	}
	fmt.Println("Data = ", string(data))
	key := state.data.Key()
	fmt.Println("Key = ", key)
	/*
		if err := kvs.Write(key, data); err != nil {
			return erros.New("Unable to write data for key (%s)", key)
		}
	*/
	return nil
}

func (state CommonStateModel) Clear() error {
	key := state.data.Key()
	fmt.Println("Key = ", key)
	/*
		if err := kvs.Delete(key); err != nil {
			return erros.New("Unable to write data for key (%s)", key)
		}
	*/
	return nil
}

func (state CommonStateModel) Set(data []byte) error {
	if err := state.mOper.Unmarshal(data, state.data); err != nil {
		return fmt.Errorf("Unable to parse data (%s)", string(data))
	}
	return nil
}

func (state CommonStateModel) Data() StateData {
	return state.data
}
