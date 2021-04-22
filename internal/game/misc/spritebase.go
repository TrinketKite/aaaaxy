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

package misc

import (
	"fmt"
	"strconv"

	"github.com/divVerent/aaaaaa/internal/engine"
	"github.com/divVerent/aaaaaa/internal/game/constants"
	"github.com/divVerent/aaaaaa/internal/level"
	m "github.com/divVerent/aaaaaa/internal/math"
)

// SpriteBase is a base class for sprites.
// To instantiate it, just set the entity image, then forward to this.
type SpriteBase struct{}

func (s *SpriteBase) Spawn(w *engine.World, sp *level.Spawnable, e *engine.Entity) error {
	w.SetSolid(e, sp.Properties["solid"] == "true")
	w.SetOpaque(e, sp.Properties["opaque"] == "true")
	if s := sp.Properties["player_solid"]; s != "" {
		w.MutateContentsBool(e, level.PlayerSolidContents, s == "true")
	}
	if s := sp.Properties["object_solid"]; s != "" {
		w.MutateContentsBool(e, level.ObjectSolidContents, s == "true")
	}
	if sp.Properties["alpha"] != "" {
		var err error
		e.Alpha, err = strconv.ParseFloat(sp.Properties["alpha"], 64)
		if err != nil {
			return fmt.Errorf("could not decode alpha %q: %v", sp.Properties["alpha"], err)
		}
	}
	if mapBlackToString := sp.Properties["map_black_to"]; mapBlackToString != "" {
		var r, g, b, a int
		if _, err := fmt.Sscanf(mapBlackToString, "#%02x%02x%02x%02x", &a, &r, &g, &b); err != nil {
			return fmt.Errorf("could not decode color %q: %v", mapBlackToString, err)
		}
		e.ColorAdd[0] = float64(r) / 255.0
		e.ColorAdd[1] = float64(g) / 255.0
		e.ColorAdd[2] = float64(b) / 255.0
		e.ColorAdd[3] = float64(a) / 255.0
	}
	if mapWhiteToString := sp.Properties["map_white_to"]; mapWhiteToString != "" {
		var r, g, b, a int
		if _, err := fmt.Sscanf(mapWhiteToString, "#%02x%02x%02x%02x", &a, &r, &g, &b); err != nil {
			return fmt.Errorf("could not decode color %q: %v", mapWhiteToString, err)
		}
		e.ColorMod[0] = float64(r)/255.0 - e.ColorAdd[0]
		e.ColorMod[1] = float64(g)/255.0 - e.ColorAdd[1]
		e.ColorMod[2] = float64(b)/255.0 - e.ColorAdd[2]
		e.ColorMod[3] = float64(a)/255.0 - e.ColorAdd[3]
	}
	if sp.Properties["z_index"] != "" {
		zIndex, err := strconv.Atoi(sp.Properties["z_index"])
		if err != nil {
			return fmt.Errorf("could not decode z index %q: %v", sp.Properties["z_index"], err)
		}
		if zIndex < constants.MinSpriteZ || zIndex > constants.MaxSpriteZ {
			return fmt.Errorf("z index out of range: got %v, want %v..%v", zIndex, constants.MinSpriteZ, constants.MaxSpriteZ)
		}
		w.SetZIndex(e, zIndex)
	}
	if sp.Properties["no_transform"] == "true" {
		// Undo transform of orientation by tile.
		e.Orientation = sp.Orientation
	}
	if e.Transform.Determinant() < 0 {
		// e.Orientation: in-editor transform. Applied first.
		// Normally the formula is e.Transform.Inverse().Concat(e.Orientation).
		// Add an FlipX() between the two to "undo" any sense difference in the editor.
		// This flips the view on the _level editor_ X axis.
		switch sp.Properties["no_flip"] {
		case "x":
			e.Orientation = e.Transform.Inverse().Concat(m.FlipX()).Concat(sp.Orientation)
		case "y":
			e.Orientation = e.Transform.Inverse().Concat(m.FlipY()).Concat(sp.Orientation)
		}
	}
	return nil
}

// The other methods to reduce code duplication in implementors.

func (s *SpriteBase) Despawn() {}

func (s *SpriteBase) Update() {}

func (s *SpriteBase) Touch(other *engine.Entity) {}
