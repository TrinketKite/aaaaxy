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

package demo

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/divVerent/aaaaxy/internal/flag"
	"github.com/divVerent/aaaaxy/internal/font"
	"github.com/divVerent/aaaaxy/internal/log"
	m "github.com/divVerent/aaaaxy/internal/math"
)

var (
	demoPlayRegressionPrefix = flag.String("demo_play_regression_prefix", "", "dump screenshots of regressions to files with this prefix")
)

var (
	regressionCount           int
	regressionScreenshotCount int
	regressionsPrevFrame      bool
	regressionsThisFrame      []string
	regressionsToDraw         []string
)

func regression(format string, args ...interface{}) {
	regression := fmt.Sprintf(format, args...)
	log.Errorf("REGRESSION: %s", regression)
	regressionsThisFrame = append(regressionsThisFrame, regression)
	regressionCount++
}

func regressionPostPlayFrame() {
	// Update state.
	regressions := regressionsThisFrame
	haveRegressions := len(regressions) > 0
	hadRegressions := regressionsPrevFrame
	regressionsPrevFrame = haveRegressions
	regressionsThisFrame = nil

	// Report this regression?
	if haveRegressions && !hadRegressions {
		// Worth reporting, not a dupe from last frame.
		regressionsToDraw = append(regressionsToDraw, fmt.Sprintf("Frame %d:", demoPlayerFrameIdx))
		regressionsToDraw = append(regressionsToDraw, regressions...)
	}
}

func regressionPostDrawFrame(screen *ebiten.Image) {
	// Update state.
	regressions := regressionsToDraw
	regressionsToDraw = nil

	// Only if we have regressions.
	if len(regressions) == 0 {
		return
	}

	// Only if actually active.
	if *demoPlayRegressionPrefix == "" {
		return
	}

	// Duplicate screen.
	// This isn't just NewImageFromImage as we want to remove the alpha channel to get a proper screenshot.
	w, h := screen.Size()
	dup := ebiten.NewImage(w, h)
	dup.Fill(color.Gray{0})
	dup.DrawImage(screen, &ebiten.DrawImageOptions{
		CompositeMode: ebiten.CompositeModeSourceOver,
	})

	// Draw text on it.
	text := strings.Join(regressions, "\n")
	bounds := font.DebugSmall.BoundString(text)
	font.DebugSmall.Draw(dup, text, m.Pos{
		X: w / 2,
		Y: -bounds.Origin.Y,
	}, true, color.Gray{0}, color.Gray{255})

	// Remove alpha.
	// Actually we could just draw a black rectangle on it with GL_SRC_ALPHA GL_ZERO.

	// Build a file name.
	name := fmt.Sprintf("%s_%04d.png", *demoPlayRegressionPrefix, regressionScreenshotCount)
	log.Errorf("dumping regression screenshot to %v", name)
	regressionScreenshotCount++
	file, err := os.Create(name)
	if err != nil {
		log.Fatalf("failed to open regression file %v: %v", name, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close regression file %v: %v", name, err)
		}
	}()
	err = png.Encode(file, dup)
	if err != nil {
		log.Fatalf("failed to write to regression file %v: %v", name, err)
	}
}

func regressionBeforeExit() {
	if regressionCount != 0 {
		log.Fatalf("detected %d regressions", regressionCount)
	}
}
