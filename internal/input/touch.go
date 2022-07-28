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

package input

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/divVerent/aaaaxy/internal/flag"
	"github.com/divVerent/aaaaxy/internal/image"
	m "github.com/divVerent/aaaaxy/internal/math"
)

var (
	touch      = flag.Bool("touch", true, "enable touch input")
	touchForce = flag.Bool("touch_force", flag.SystemDefault(map[string]interface{}{
		"android/*": true,
		"ios/*":     true,
		"js/*":      false, // TODO(divVerent): Figure out why they fail, and once solved, enable.
		"*/*":       false,
	}).(bool), "always show touch controls")
)

var (
	leftTouch   = m.Rect{Origin: m.Pos{X: 0, Y: 296}, Size: m.Delta{DX: 32, DY: 32}}
	rightTouch  = m.Rect{Origin: m.Pos{X: 32, Y: 296}, Size: m.Delta{DX: 32, DY: 32}}
	downTouch   = m.Rect{Origin: m.Pos{X: 16, Y: 328}, Size: m.Delta{DX: 32, DY: 32}}
	upTouch     = m.Rect{Origin: m.Pos{X: 16, Y: 264}, Size: m.Delta{DX: 32, DY: 32}}
	jumpTouch   = m.Rect{Origin: m.Pos{X: 608, Y: 328}, Size: m.Delta{DX: 32, DY: 32}}
	actionTouch = m.Rect{Origin: m.Pos{X: 576, Y: 328}, Size: m.Delta{DX: 32, DY: 32}}
	exitTouch   = m.Rect{Origin: m.Pos{X: 0, Y: 0}, Size: m.Delta{DX: 64, DY: 32}}
)

const (
	touchClickMaxFrames = 30
	touchPadFrames      = 300
)

type touchInfo struct {
	frames int
	pos    m.Pos
	hit    bool
}

var (
	touchWantPad  bool
	touches       = map[ebiten.TouchID]*touchInfo{}
	touchIDs      []ebiten.TouchID
	touchHoverPos m.Pos
	touchPadFrame int
)

func touchUpdate(screenWidth, screenHeight, gameWidth, gameHeight int, crtK1, crtK2 float64) {
	if !*touch {
		return
	}
	for _, t := range touches {
		t.hit = false
	}
	touchIDs = ebiten.AppendTouchIDs(touchIDs[:0])
	if len(touchIDs) > 0 {
		// Either support touch OR mouse. This prevents duplicate click events.
		mouseCancel()
		touchPadFrame = touchPadFrames
	} else if touchPadFrame > 0 {
		touchPadFrame--
	}
	for _, id := range touchIDs {
		t, found := touches[id]
		if !found {
			t = &touchInfo{}
			touches[id] = t
		}
		t.hit = true
		t.frames++
	}
	hoverAcc := m.Pos{}
	hoverCnt := 0
	for id, t := range touches {
		if !t.hit {
			if t.frames < touchClickMaxFrames {
				clickPos = &t.pos
			}
			delete(touches, id)
			continue
		}
		x, y := ebiten.TouchPosition(id)
		t.pos = pointerCoords(screenWidth, screenHeight, gameWidth, gameHeight, crtK1, crtK2, x, y)
		if t.frames < touchClickMaxFrames {
			hoverAcc = hoverAcc.Add(t.pos.Delta(m.Pos{}))
			hoverCnt++
		}
	}
	if hoverCnt > 0 {
		touchHoverPos = hoverAcc.Add(m.Delta{DX: hoverCnt / 2, DY: hoverCnt / 2}).Div(hoverCnt)
		hoverPos = &touchHoverPos
	}
}

func touchSetWantPad(want bool) {
	touchWantPad = want
}

func (i *impulse) touchPressed() InputMap {
	if !touchWantPad {
		return 0
	}
	if i.touchRect.Size.IsZero() {
		return 0
	}
	for _, t := range touches {
		if i.touchRect.DeltaPos(t.pos).IsZero() {
			return Touchscreen
		}
	}
	return 0
}

func touchInit() error {
	var err error
	Left.touchImage, err = image.Load("sprites", "touch_left.png")
	if err != nil {
		return err
	}
	Right.touchImage, err = image.Load("sprites", "touch_right.png")
	if err != nil {
		return err
	}
	Up.touchImage, err = image.Load("sprites", "touch_up.png")
	if err != nil {
		return err
	}
	Down.touchImage, err = image.Load("sprites", "touch_down.png")
	if err != nil {
		return err
	}
	Jump.touchImage, err = image.Load("sprites", "touch_jump.png")
	if err != nil {
		return err
	}
	Action.touchImage, err = image.Load("sprites", "touch_action.png")
	if err != nil {
		return err
	}
	Exit.touchImage, err = image.Load("sprites", "touch_exit.png")
	if err != nil {
		return err
	}
	return nil
}

func touchDraw(screen *ebiten.Image) {
	if !touchWantPad {
		return
	}
	if !*touchForce && touchPadFrame <= 0 {
		return
	}
	for _, i := range impulses {
		if i.touchRect.Size.IsZero() {
			continue
		}
		img := i.touchImage
		if img == nil {
			continue
		}
		options := &ebiten.DrawImageOptions{
			CompositeMode: ebiten.CompositeModeSourceOver,
			Filter:        ebiten.FilterNearest,
		}
		w, h := img.Size()
		options.GeoM.Scale(
			float64(i.touchRect.Size.DX)/float64(w),
			float64(i.touchRect.Size.DY)/float64(h))
		options.GeoM.Translate(float64(i.touchRect.Origin.X), float64(i.touchRect.Origin.Y))
		if i.Held {
			options.ColorM.Scale(-1, -1, -1, 1)
			options.ColorM.Translate(1, 1, 1, 0)
		}
		screen.DrawImage(img, options)
	}
}
