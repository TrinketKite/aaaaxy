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

package ending

import (
	"github.com/divVerent/aaaaxy/internal/engine"
	"github.com/divVerent/aaaaxy/internal/fun"
	"github.com/divVerent/aaaaxy/internal/level"
	"github.com/divVerent/aaaaxy/internal/log"
)

// CreditsTarget shows the credits.
type CreditsTarget struct {
	World *engine.World
	State bool
}

func (c *CreditsTarget) Spawn(w *engine.World, sp *level.SpawnableProps, e *engine.Entity) error {
	c.World = w
	return nil
}

func (c *CreditsTarget) Despawn() {}

func (c *CreditsTarget) Update() {}

func (c *CreditsTarget) SetState(originator, predecessor *engine.Entity, state bool) {
	if state == c.State {
		return
	}
	c.State = state
	if !state {
		return
	}
	c.World.ForceCredits = true
	c.World.PlayerState.SetWon()
	err := c.World.Save()
	if err != nil {
		log.Errorf("could not save game: %v", err)
	}

	log.Infof("%v", fun.FormatText(&c.World.PlayerState,
		"your time: {{GameTime}}; your speedrun categories: {{SpeedrunCategories}}; try next: {{SpeedrunTryNext}}."))
}

func (c *CreditsTarget) Touch(other *engine.Entity) {}

func init() {
	engine.RegisterEntityType(&CreditsTarget{})
}
