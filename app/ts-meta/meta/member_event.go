/*
Copyright 2022 Huawei Cloud Computing Technologies Co., Ltd.

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

package meta

import "github.com/openGemini/openGemini/lib/util/lifted/hashicorp/serf/serf"

type memberEvent struct {
	event   serf.MemberEvent
	handler memberEventHandler
}

func initMemberEvent(e serf.MemberEvent, handler memberEventHandler) *memberEvent {
	return &memberEvent{
		event:   e,
		handler: handler,
	}
}

func (m *memberEvent) handle() error {
	return m.handler.handle(m)
}
