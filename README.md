# avatarbuilder
Using go freetype to build default avatar with string

## Install

```
go get -u github.com/ShiningRush/avatarbuilder
```

## Usage

You can referrence `./example`

Some snipet is as blow

```
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
```
