// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"os"
	"runtime/pprof"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/divVerent/aaaaaa/internal/aaaaaa"
	"github.com/divVerent/aaaaaa/internal/flag"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write CPU profile to file")
)

func main() {
	flag.Parse(aaaaaa.LoadConfig)
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	aaaaaa.InitEbiten()
	game := &aaaaaa.Game{}
	err := ebiten.RunGame(game)
	aaaaaa.BeforeExit()
	if err != nil {
		log.Print(err)
	}
}
