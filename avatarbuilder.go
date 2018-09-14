package avatarbuilder

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

type FontCenterCalculator interface {
	CalculateCenterLocation(string, *AvatarBuilder) (int, int)
}

type AvatarBuilder struct {
	W        int
	H        int
	fontfile string
	fontsize float64
	bg       color.Color
	fg       color.Color
	ctx      *freetype.Context
	calc     FontCenterCalculator
}

func NewAvatarBuilder(fontfile string, calc FontCenterCalculator) *AvatarBuilder {
	ab := &AvatarBuilder{}
	ab.fontfile = fontfile
	ab.bg, ab.fg = color.White, color.Black
	ab.W, ab.H = 200, 200
	ab.fontsize = 95
	ab.calc = calc

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

func (ab *AvatarBuilder) SetFontSize(size float64) {
	ab.fontsize = size
}

func (ab *AvatarBuilder) SetAvatarSize(w int, h int) {
	ab.W = w
	ab.H = h
}

func (ab *AvatarBuilder) GenerateImage(s string, outname string) error {
	rgba := ab.buildColorImage()
	if ab.ctx == nil {
		if err := ab.buildDrawContext(rgba); err != nil {
			return err
		}
	}

	x, y := ab.calc.CalculateCenterLocation(s, ab)
	pt := freetype.Pt(x, y)
	if _, err := ab.ctx.DrawString(s, pt); err != nil {
		return errors.New("draw string: " + err.Error())
	}

	return ab.save(outname, rgba)
}

func (ab *AvatarBuilder) buildColorImage() *image.RGBA {
	bg := image.NewUniform(ab.bg)
	rgba := image.NewRGBA(image.Rect(0, 0, ab.W, ab.H))
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

func (ab *AvatarBuilder) buildDrawContext(rgba *image.RGBA) error {
	// Read the font data.
	fontBytes, err := ioutil.ReadFile(ab.fontfile)
	if err != nil {
		return errors.New("error when open font file:" + err.Error())
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return errors.New("error when parse font file:" + err.Error())
	}

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(ab.fontsize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(image.NewUniform(ab.fg))
	c.SetHinting(font.HintingNone)

	ab.ctx = c
	return nil
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

func (ab *AvatarBuilder) GetFontWidth() int {
	return int(ab.ctx.PointToFixed(ab.fontsize) >> 6)
}
