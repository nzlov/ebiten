// Copyright 2014 Hajime Hoshi
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

package blocks

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	"image/color"
	"math"
)

func init() {
	imagePaths["font"] = "images/blocks/font.png"
}

const charWidth = 8
const charHeight = 8

func textWidth(str string) int {
	return charWidth * len(str)
}

func drawText(rt *ebiten.Image, images *Images, str string, ox, oy, scale int, c color.Color) {
	fontImageId := images.GetImage("font")
	dsts, srcs := []image.Rectangle{}, []image.Rectangle{}

	locationX, locationY := 0, 0
	for _, c := range str {
		if c == '\n' {
			locationX = 0
			locationY += charHeight
			continue
		}
		code := int(c)
		x := (code % 16) * charWidth
		y := ((code - 32) / 16) * charHeight
		dsts = append(dsts, image.Rect(locationX, locationY, locationX+charWidth, locationY+charHeight))
		srcs = append(srcs, image.Rect(x, y, x+charWidth, y+charHeight))
		locationX += charWidth
	}

	geo := ebiten.ScaleGeo(float64(scale), float64(scale))
	geo.Concat(ebiten.TranslateGeo(float64(ox), float64(oy)))
	c2 := color.NRGBA64Model.Convert(c).(color.NRGBA64)
	const max = math.MaxUint16
	r := float64(c2.R) / max
	g := float64(c2.G) / max
	b := float64(c2.B) / max
	a := float64(c2.A) / max
	clr := ebiten.ScaleColor(r, g, b, a)
	rt.DrawImage(fontImageId, &ebiten.DrawImageOptions{
		DstParts: dsts,
		SrcParts: srcs,
		GeoM:     geo,
		ColorM:   clr,
	})
}

func drawTextWithShadow(rt *ebiten.Image, images *Images, str string, x, y, scale int, clr color.Color) {
	drawText(rt, images, str, x+1, y+1, scale, color.NRGBA{0, 0, 0, 0x80})
	drawText(rt, images, str, x, y, scale, clr)
}
