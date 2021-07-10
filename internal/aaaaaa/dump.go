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

package aaaaaa

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/divVerent/aaaaaa/internal/audiowrap"
	"github.com/divVerent/aaaaaa/internal/engine"
	"github.com/divVerent/aaaaaa/internal/flag"
	m "github.com/divVerent/aaaaaa/internal/math"
)

var (
	dumpVideo = flag.String("dump_video", "", "filename prefix to dump game frames to")
	dumpAudio = flag.String("dump_audio", "", "filename to dump game audio to")
)

var (
	dumpFrameCount = 0
	dumpVideoFile  *os.File
	dumpAudioFile  *os.File
	pngEncoder     = png.Encoder{
		CompressionLevel: png.BestSpeed,
	}
)

func initDumping() error {
	if *dumpAudio != "" {
		var err error
		dumpAudioFile, err = os.Create(*dumpAudio)
		if err != nil {
			return fmt.Errorf("could not initialize audio dump: %v", err)
		}
		audiowrap.InitDumping()
	}

	if *dumpVideo != "" {
		var err error
		dumpVideoFile, err = os.Create(*dumpVideo)
		if err != nil {
			return fmt.Errorf("could not initialize video dump: %v", err)
		}
	}

	return nil
}

func dumping() bool {
	return dumpAudioFile != nil || dumpVideoFile != nil
}

func dumpFrame(screen *ebiten.Image) {
	if !dumping() {
		return
	}
	dumpFrameCount++
	if dumpVideoFile != nil {
		err := pngEncoder.Encode(dumpVideoFile, screen)
		if err != nil {
			log.Printf("Failed to encode video - expect corruption: %v", err)
			dumpVideoFile.Close()
			dumpVideoFile = nil
		}
	}
	if dumpAudioFile != nil {
		err := audiowrap.DumpFrame(dumpAudioFile, time.Duration(dumpFrameCount)*time.Second/engine.GameTPS)
		if err != nil {
			log.Printf("Failed to encode audio - expect corruption: %v", err)
			dumpAudioFile.Close()
			dumpAudioFile = nil
		}
	}
}

func ffmpegCommand(audio, video, output string) string {
	var pre string
	inputs := []string{}
	settings := []string{}
	// Video first, so we can refer to the video stream as [0:v] for sure.
	if video != "" {
		inputs = append(inputs, fmt.Sprintf("-f png_pipe -r %d -i '%s'", engine.GameTPS, strings.ReplaceAll(video, "'", "'\\''")))
		// Note: the two step upscale simulates the effect of the normal2x shader.
		// Note: using high quality, fast settings and many keyframes
		// as the assumption is that the output file will be further edited.
		// Note: disabling 8x8 DCT here as some older FFmpeg versions -
		// or even newer versions with decoding options changed for compatibility,
		// if the video file has also been losslessly cut -
		// have trouble decoding that.
		var filterComplex string
		switch *screenFilter {
		case "linear":
			filterComplex = "[0:v]premultiply=inplace=1,scale=1920:1080"
		case "linear2x":
			filterComplex = "[0:v]premultiply=inplace=1,scale=1280:720:flags=neighbor,scale=1920:1080"
		case "linear2xcrt":
			darkened := m.Rint(255 * (1.0 - 2.0/3.0**screenFilterScanLines))
			pre = fmt.Sprintf("echo 'P2 1 3 255 %d 255 %d' | convert -size 1920x1080 TILE:PNM:- scanlines.png; ", darkened, darkened)
			filterComplex = fmt.Sprintf("[0:v]premultiply=inplace=1,scale=1280:720:flags=neighbor,scale=1920:1080,format=gbrp[scaled]; movie=scanlines.png,format=gbrp[scanlines]; [scaled][scanlines]blend=all_mode=multiply,lenscorrection=i=bilinear:k1=%f:k2=%f", crtK1(), crtK2())
		case "simple", "nearest":
			filterComplex = "[0:v]premultiply=inplace=1,scale=1920:1080:flags=neighbor"
		}
		settings = append(settings, "-codec:v libx264 -profile:v high444 -preset:v fast -crf:v 10 -8x8dct:v 0 -keyint_min 10 -g 60 -filter_complex '"+filterComplex+"'")
	}
	if audio != "" {
		inputs = append(inputs, fmt.Sprintf("-f s16le -ac 2 -ar %d  -i '%s'", audiowrap.Rate(), strings.ReplaceAll(audio, "'", "'\\''")))
		settings = append(settings, "-codec:a aac -b:a 128k")
	}
	return fmt.Sprintf("%sffmpeg %s %s -vsync vfr %s", pre, strings.Join(inputs, " "), strings.Join(settings, " "), strings.ReplaceAll(output, "'", "'\\''"))
}

func finishDumping() {
	if !dumping() {
		return
	}
	if dumpAudioFile != nil {
		err := dumpAudioFile.Close()
		if err != nil {
			log.Printf("Failed to close audio - expect corruption: %v", err)
		}
		dumpAudioFile = nil
	}
	if dumpVideoFile != nil {
		err := dumpVideoFile.Close()
		if err != nil {
			log.Printf("Failed to close video - expect corruption: %v", err)
		}
		dumpVideoFile = nil
	}
	log.Print("Media has been dumped.")
	log.Print("To convert to something uploadable, run:")
	log.Print(ffmpegCommand(*dumpAudio, *dumpVideo, "video.mp4"))
}
