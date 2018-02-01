/*
 *  Copyright (c) 2018, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/nebulaim/telegramd/baselib/app"
	"github.com/nebulaim/telegramd/access/frontend2/server"
	"github.com/golang/glog"
	"net"
)

type ServerConfig struct {
	Network string
	Addr string
	ProtoName string
	ServiceName string
}

type AuthKeyInsance struct {
	server *server.FrontendServer
}

func (this *AuthKeyInsance) Initialize() error {
	listener, err := net.Listen("tcp", "0.0.0.0:22345")
	if err != nil {
		glog.Errorf("listen error: %v", err)
		return err
	}

	this.server = server.NewFrontendServer(listener, "ztproto")
	return nil
}

func (this *AuthKeyInsance) RunLoop() {
	this.server.Serve()
}

func (this *AuthKeyInsance) Destroy() {
	this.server.Stop()
}


func main() {
	instance := &AuthKeyInsance{}
	app.DoMainAppInsance(instance)
}
