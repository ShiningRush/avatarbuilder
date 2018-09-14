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

	ab := avatarbuilder.NewAvatarBuilder("./SourceHanSansSC-Medium.ttf", &calc.SourceHansSansSCMedium{})
	ab.SetBackgroundColorHex(colors[1])
	ab.SetFrontgroundColor(color.White)
	ab.SetFontSize(80)
	ab.SetAvatarSize(200, 200)
	if err := ab.GenerateImage("缄默", "./out.png"); err != nil {
		fmt.Println(err)
		return
	}
}
