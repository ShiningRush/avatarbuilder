# avatarbuilder
Using [go freetype](https://github.com/golang/freetype) to build default avatar with string


![number text](https://github.com/ShiningRush/avatarbuilder/blob/master/example/out.png "number text")
![english text](https://github.com/ShiningRush/avatarbuilder/blob/master/example/outEn.png "english text")
![chinese text](https://github.com/ShiningRush/avatarbuilder/blob/master/example/outCh.png "chinese text")

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

## Extend Other Font

Because element of width of each font is different, so you need tell avatar how to align the content.
Avatar already implement a free font(made by google and adobe)'s center algorithm in `./calc`,
If you need other font, feel free to PR or issue.
