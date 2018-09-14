package example

import (
	"flag"
	"fmt"
	"image/color"

	"github.com/ShiningRush/avatarbuilder"
)

var colors = []uint32{
	0xff6200, 0x42c58e, 0x5a8de1, 0x785fe0,
}

func main() {
	flag.Parse()

	ab := avatarbuilder.NewAvatarBuilder("./RuanMengTi-2.ttf")
	ab.SetBackgroundColorHex(colors[0])
	ab.SetFrontgroundColor(color.White)
	ab.SetFontStyle(40, 40, 95)
	if err := ab.GenerateImage("04", "./out.png"); err != nil {
		fmt.Println(err)
		return
	}
}
