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

// implements the State interface for a network implemented using
// vlans with ovs. The state is stored as Json objects.
const (
	BASE_PATH           = "/contiv/"
	CFG_PATH            = BASE_PATH + "config/"
	OPER_PATH           = BASE_PATH + "oper/"
	EP_CFG_PATH_PREFIX  = CFG_PATH + "eps/"
	EP_OPER_PATH_PREFIX = OPER_PATH + "eps/"
	NW_CFG_PATH_PREFIX  = CFG_PATH + "nets/"
	NW_OPER_PATH_PREFIX = OPER_PATH + "nets/"
)
