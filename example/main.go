package main

import (
	"flag"
	"fmt"
	"image/color"

	"github.com/ShiningRush/avatarbuilder"
	"github.com/ShiningRush/avatarbuilder/calc"
)

var colors = []uint32{
	0xff6200, 0x42c58e, 0x5a8de1, 0x785fe0,
}

func main() {
	flag.Parse()

	// init avatarbuilder, you need to tell builder ttf file and how to alignment text
	ab := avatarbuilder.NewAvatarBuilder("./SourceHanSansSC-Medium.ttf", &calc.SourceHansSansSCMedium{})
	ab.SetBackgroundColorHex(colors[1])
	ab.SetFrontgroundColor(color.White)
	ab.SetFontSize(80)
	ab.SetAvatarSize(200, 200)
	if err := ab.GenerateImageAndSave("12", "./out.png"); err != nil {
		fmt.Println(err)
		return
	}
}
