package tests

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"testing"

	"github.com/golang/freetype/truetype"
	"github.com/lukhaiminh/go-captcha/v2/base/option"
	"github.com/lukhaiminh/go-captcha/v2/click"
)

var textCapt click.Captcha

func init() {
	builder := click.NewBuilder(
		click.WithRangeLen(option.RangeVal{Min: 4, Max: 6}),
		click.WithRangeVerifyLen(option.RangeVal{Min: 2, Max: 4}),
		click.WithDisabledRangeVerifyLen(true),
	)

	fontN, err := loadFont("../.cache/fzshengsksjw_cu.ttf")
	if err != nil {
		log.Fatalln(err)
	}

	bgImage, err := loadPng("../.cache/bg.png")
	if err != nil {
		log.Fatalln(err)
	}

	builder.SetResources(
		//click.WithChars([]string{"这", "是", "随", "机", "的", "文", "本", "种", "子", "呀"}),
		click.WithChars([]string{"A1", "B2", "C3", "D4", "E5", "F6", "G7", "H8", "I9", "J0"}),
		click.WithFonts([]*truetype.Font{
			fontN,
		}),
		click.WithBackgrounds([]image.Image{
			bgImage,
		}),
		//click.WithThumbBackgrounds([]image.Image{
		//	thumbImage,
		//}),
	)

	textCapt = builder.Make()
}

func TestClickTextCaptcha(t *testing.T) {
	captData, err := textCapt.Generate()
	if err != nil {
		log.Fatalln(err)
	}

	dotData := captData.GetData()
	if dotData == nil {
		log.Fatalln(">>>>> generate err")
	}

	dots, _ := json.Marshal(dotData)
	fmt.Println(string(dots))
	fmt.Println(captData.GetMasterImage().ToBase64())
	fmt.Println(captData.GetThumbImage().ToBase64())

	err = captData.GetMasterImage().SaveToFile("../.cache/master.jpg", option.QualityNone)
	if err != nil {
		fmt.Println(err)
	}
	err = captData.GetThumbImage().SaveToFile("../.cache/thumb.png")
	if err != nil {
		fmt.Println(err)
	}
}
