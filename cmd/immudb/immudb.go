/*
Copyright 2019-2020 vChain, Inc.

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

package main

import (
	c "github.com/codenotary/immudb/cmd/helper"
	immudb "github.com/codenotary/immudb/cmd/immudb/command"
	"github.com/codenotary/immudb/cmd/version"
	"github.com/codenotary/immudb/pkg/server"
)

func main() {
	execute(server.DefaultServer())
}

func execute(immudbServer server.ImmuServerIf) error {
	version.App = "immudb"
	cl := immudb.Commandline{P: c.NewPlauncher()}
	cmd := cl.NewCmd(immudbServer)
	return cmd.Execute()
}
