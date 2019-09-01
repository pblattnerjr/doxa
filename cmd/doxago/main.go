// Copyright © 2019 The Orthodox Christian Mission Center (OCMC)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/liturgiko/doxa/cmd/doxago/cmd"
	"log"
	"os"
)

var DOXAHOME = os.Getenv("DOXAHOME")
var LIMLHOME = os.Getenv("LIMLHOME")
var DOXAPORT = os.Getenv("DOXAPORT")

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	checkVar(DOXAHOME)
	checkVar(DOXAPORT)
	checkVar(LIMLHOME)
	cmd.Execute()
}
func checkVar(v string) {
	if v == "" {
		fmt.Println(tmpl)
		os.Exit(1)
	}
}
const tmpl = `
doxa cannot run unless you first set 
the environmental variables shown below.
How you do this depends on your operating system.

On a Mac, edit .bash_profile in your home directory and
add the following:

export DOXAPORT=8080
export DOXAHOME=$HOME/doxa
export PATH=$PATH:$DOXAHOME
export LIMLHOME=$HOME/liml

On Windows, TBD...`