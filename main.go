// Copyright 2017 The Ebiten Authors
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

// +build example jsgo

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	raudio "github.com/hajimehoshi/ebiten/examples/resources/audio"
)

const (
	screenWidth  = 320
	screenHeight = 240
	sampleRate   = 22050

	introLengthInSecond = 5
	loopLengthInSecond  = 4
)

var audioContext *audio.Context

func init() {
	var err error
	audioContext, err = audio.NewContext(sampleRate)
	if err != nil {
		log.Fatal(err)
	}
}

var player *audio.Player

func update(screen *ebiten.Image) error {
	if player == nil {
		// Decode the wav file.
		// wavS is a decoded io.ReadCloser and io.Seeker.
		oggS, err := mp3.Decode(audioContext, audio.BytesReadSeekCloser(raudio.Ragtime_ogg))
		if err != nil {
			return err
		}

		// Create an infinite loop stream from the decoded bytes.
		// s is still an io.ReadCloser and io.Seeker.
		s := audio.NewInfiniteLoopWithIntro(oggS, introLengthInSecond*4*sampleRate, loopLengthInSecond*4*sampleRate)

		player, err = audio.NewPlayer(audioContext, s)
		if err != nil {
			return err
		}

		// Play the infinite-length stream. This never ends.
		player.Play()
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	pos := player.Current()
	if pos > 5*time.Second {
		pos = (player.Current()-5*time.Second)%(4*time.Second) + 5*time.Second
	}
	msg := fmt.Sprintf(`TPS: %0.2f
This is an example using
audio.NewInfiniteLoopWithIntro.

Intro:   0[s] - %[2]d[s]
Loop:    %[2]d[s] - %[3]d[s]
Current: %0.2[4]f[s]`, ebiten.CurrentTPS(), introLengthInSecond, introLengthInSecond+loopLengthInSecond, float64(pos)/float64(time.Second))
	ebitenutil.DebugPrint(screen, msg)
	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Audio Infinite Loop (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
