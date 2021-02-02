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

package menu

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/divVerent/aaaaaa/internal/engine"
	_ "github.com/divVerent/aaaaaa/internal/game" // Load entities.
	"github.com/divVerent/aaaaaa/internal/input"
	"github.com/divVerent/aaaaaa/internal/timing"
)

var (
	resetSave = flag.Bool("reset_save", false, "reset the savegame on startup")
	showFps   = flag.Bool("show_fps", false, "show fps counter")
)

type MenuScreen interface {
	Init(m *Menu) error
	Update() error
	Draw(screen *ebiten.Image)
}

type Menu struct {
	World  engine.World
	Screen MenuScreen
}

func (m *Menu) Update() error {
	defer timing.Group()()

	timing.Section("once")
	if !m.World.Initialized() {
		m.World.Init()
		if !*resetSave {
			err := m.World.Load()
			if err != nil {
				return err
			}
		}
	}

	timing.Section("global_hotkeys")
	if input.Exit.JustHit && m.Screen == nil {
		return m.SwitchToScreen(&MainScreen{})
	}

	timing.Section("screen")
	if m.Screen != nil {
		err := m.Screen.Update()
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Menu) UpdateWorld() error {
	if m.Screen != nil {
		// Game is paused while in menu.
		return nil
	}
	return m.World.Update()
}

func (m *Menu) Draw(screen *ebiten.Image) {
	defer timing.Group()()

	timing.Section("screen")
	if m.Screen != nil {
		m.Screen.Draw(screen)
	}

	timing.Section("global_overlays")
	if *showFps {
		timing.Section("fps")
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%.1f fps, %.1f tps", ebiten.CurrentFPS(), ebiten.CurrentTPS()), 0, engine.GameHeight-16)
	}
}

func (m *Menu) DrawWorld(screen *ebiten.Image) {
	// TODO: If a menu screen is active, just draw the previous saved bitmap, but blur it.
	m.World.Draw(screen)
}

// SwitchToGame is called by menu screens to view the game.
func (m *Menu) SwitchToGame() error {
	m.Screen = nil
	return nil
}

// SwitchToCheckpoitn switches to a specific checkpoint.
func (m *Menu) SwitchToCheckpoint(cp string) error {
	m.World.RespawnPlayer(cp)
	return m.SwitchToGame()
}

// SwitchToScreen is called by menu screens to go to a different menu screen.
func (m *Menu) SwitchToScreen(screen MenuScreen) error {
	m.Screen = screen
	return m.Screen.Init(m)
}

// QuitGame is called by menu screens to end the game.
func (m *Menu) QuitGame() error {
	err := m.World.Save()
	if err != nil {
		log.Panicf("could not save game: %v", err)
	}
	return errors.New("game exited normally")
}
