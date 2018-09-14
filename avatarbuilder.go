package main

import (
	"bufio"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

type AvatarBuilder struct {
	fontfile string
	fontsize float64
	bg       color.Color
	fg       color.Color
	w        int
	h        int
	x        int
	y        int
}

func NewAvatarBuilder(fontfile string) *AvatarBuilder {
	ab := &AvatarBuilder{}
	ab.fontfile = fontfile
	ab.bg, ab.fg = color.White, color.Black
	ab.w, ab.h = 200, 200
	ab.x, ab.y = 10, 40
	ab.fontsize = 95

	return ab
}

func (ab *AvatarBuilder) SetFrontgroundColor(c color.Color) {
	ab.fg = c
}

func (ab *AvatarBuilder) SetBackgroundColor(c color.Color) {
	ab.bg = c
}

func (ab *AvatarBuilder) SetFrontgroundColorHex(hex uint32) {
	ab.fg = ab.hexToRGBA(hex)
}

func (ab *AvatarBuilder) SetBackgroundColorHex(hex uint32) {
	ab.bg = ab.hexToRGBA(hex)
}

func (ab *AvatarBuilder) SetFontStyle(x int, y int, size float64) {
	ab.x = x
	ab.y = y
	ab.fontsize = size
}

func (ab *AvatarBuilder) SetAvatarSize(w int, h int) {
	ab.w = w
	ab.h = h
}

func (ab *AvatarBuilder) GenerateImage(s string, outname string) error {
	rgba := ab.buildColorImage()
	c, err := ab.buildDrawContext(rgba)
	if err != nil {
		return err
	}

	pt := freetype.Pt(ab.x, ab.y+int(c.PointToFixed(ab.fontsize)>>6))
	if _, err := c.DrawString(s, pt); err != nil {
		return errors.New("draw string: " + err.Error())
	}

	return ab.save(outname, rgba)
}

func (ab *AvatarBuilder) buildColorImage() *image.RGBA {
	bg := image.NewUniform(ab.bg)
	rgba := image.NewRGBA(image.Rect(0, 0, ab.w, ab.h))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	return rgba
}

func (ab *AvatarBuilder) hexToRGBA(h uint32) *color.RGBA {
	rgba := &color.RGBA{
		R: uint8(h >> 16),
		G: uint8((h & 0x00ff00) >> 8),
		B: uint8(h & 0x0000ff),
		A: 255,
	}

	return rgba
}

func (ab *AvatarBuilder) buildDrawContext(rgba *image.RGBA) (*freetype.Context, error) {
	// Read the font data.
	fontBytes, err := ioutil.ReadFile(ab.fontfile)
	if err != nil {
		return nil, errors.New("error when open font file:" + err.Error())
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, errors.New("error when parse font file:" + err.Error())
	}

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(ab.fontsize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(image.NewUniform(ab.fg))
	c.SetHinting(font.HintingNone)

	return c, nil
}

func (ab *AvatarBuilder) save(filePath string, rgba *image.RGBA) error {
	// Save that RGBA image to disk.
	outFile, err := os.Create(filePath)
	if err != nil {
		return errors.New("error when create file: " + err.Error())
	}
	defer outFile.Close()

	b := bufio.NewWriter(outFile)

	if err = png.Encode(b, rgba); err != nil {
		return errors.New("error when encode image: " + err.Error())
	}

	if err = b.Flush(); err != nil {
		return errors.New("error when flush image: " + err.Error())
	}

	return nil
}
