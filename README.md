<div align="center">
<img width="120" style="padding-top: 50px; margin: 0;" src="http://47.104.180.148/go-captcha/gocaptcha_logo.svg?v=1"/>
<h1 style="margin: 0; padding: 0">Go Captcha</h1>
<p>Behavior Security Captcha</p>
<a href="https://goreportcard.com/report/github.com/lukhaiminh/go-captcha"><img src="https://goreportcard.com/badge/github.com/lukhaiminh/go-captcha"/></a>
<a href="https://godoc.org/github.com/lukhaiminh/go-captcha"><img src="https://godoc.org/github.com/lukhaiminh/go-captcha?status.svg"/></a>
<a href="https://github.com/lukhaiminh/go-captcha/releases"><img src="https://img.shields.io/github/v/release/wenlng/go-captcha.svg"/></a>
<a href="https://github.com/lukhaiminh/go-captcha/blob/v2/LICENSE"><img src="https://img.shields.io/badge/License-Apache2.0-green.svg"/></a>
<a href="https://github.com/lukhaiminh/go-captcha"><img src="https://img.shields.io/github/stars/wenlng/go-captcha.svg"/></a>
<a href="https://github.com/lukhaiminh/go-captcha"><img src="https://img.shields.io/github/last-commit/wenlng/go-captcha.svg"/></a>
</div>

<br/>

> English | [中文](README_zh.md)

<p style="text-align: center"><a href="https://github.com/lukhaiminh/go-captcha">Go Captcha</a> is a behavior security CAPTCHA, which implements text click verification, slide verification and rotation verification.</p>

<p style="text-align: center"> ⭐️ If it helps you, please give a star.</p>

<div align="center"> 
    <img src="http://47.104.180.148/go-captcha/go-captcha-v2.jpg" alt="Poster">
</div>

<br/>

- GoCaptcha：[https://github.com/lukhaiminh/go-captcha](https://github.com/lukhaiminh/go-captcha)
- GoCaptcha Document：[http://gocaptcha.wencodes.com](http://gocaptcha.wencodes.com)
- Go Example：[https://github.com/lukhaiminh/go-captcha-example](https://github.com/lukhaiminh/go-captcha-example)
- Go Assets File：[https://github.com/lukhaiminh/go-captcha-assets](https://github.com/lukhaiminh/go-captcha-assets)
- Vue Package：[https://github.com/lukhaiminh/go-captcha-vue](https://github.com/lukhaiminh/go-captcha-vue)
- React Package：[https://github.com/lukhaiminh/go-captcha-react](https://github.com/lukhaiminh/go-captcha-react)
- Angular Package：[https://github.com/lukhaiminh/go-captcha-angular](https://github.com/lukhaiminh/go-captcha-angular)
- Svelte Package：[https://github.com/lukhaiminh/go-captcha-svelte](https://github.com/lukhaiminh/go-captcha-svelte)
- Solid Package：[https://github.com/lukhaiminh/go-captcha-solid](https://github.com/lukhaiminh/go-captcha-solid)
- Online Demo：[http://gocaptcha.wencodes.com/demo](http://gocaptcha.wencodes.com/demo)
- ...

<br/>

## Install Captcha Module
```shell
$ go get -u github.com/lukhaiminh/go-captcha/v2@latest
```

## Import Captcha Module
```go
package main

import "github.com/lukhaiminh/go-captcha/v2"

func main(){
   // ...
}
```

<br />

## 🖖 Click Mode Captcha
### Quick use
```go
package main

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/lukhaiminh/go-captcha/v2/base/option"
	"github.com/lukhaiminh/go-captcha/v2/click"
	"github.com/lukhaiminh/go-captcha/v2/base/codec"
)

var textCapt click.Captcha

func init() {
	builder := click.NewBuilder(
		click.WithRangeLen(option.RangeVal{Min: 4, Max: 6}),
		click.WithRangeVerifyLen(option.RangeVal{Min: 2, Max: 4}),
	)

	// You can use preset material resources：https://github.com/lukhaiminh/go-captcha-assets
	fontN, err := loadFont("../resources/fzshengsksjw_cu.ttf")
	if err != nil {
		log.Fatalln(err)
	}

	bgImage, err := loadPng("../resources/bg.png")
	if err != nil {
		log.Fatalln(err)
	}

	builder.SetResources(
		click.WithChars([]string{
			"1A",
			"5E",
			"3d",
			"0p",
			"78",
			"DL",
			"CB",
			"9M",
			// ...
		}),
		click.WithFonts([]*truetype.Font{
			fontN,
		}),
		click.WithBackgrounds([]image.Image{
			bgImage,
		}),
	)

	textCapt= builder.Make()
}

func loadPng(p string) (image.Image, error) {
	imgBytes, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return codec.DecodeByteToPng(imgBytes)
}

func loadFont(p string) (*truetype.Font, error) {
	fontBytes, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}
	return freetype.ParseFont(fontBytes)
}


func main() {
	captData, err := textCapt.Generate()
	if err != nil {
		log.Fatalln(err)
	}

	dotData := captData.GetData()
	if dotData == nil {
		log.Fatalln(">>>>> generate err")
	}

	dots, _ := json.Marshal(dotData)
	fmt.Println(">>>>> ", string(dots))

	err = captData.GetMasterImage().SaveToFile("../resources/master.jpg", option.QualityNone)
	if err != nil {
		fmt.Println(err)
	}
	err = captData.GetThumbImage().SaveToFile("../resources/thumb.png")
	if err != nil {
		fmt.Println(err)
	}
}
```

### Create instance method
- builder.Make()
- builder.MakeWithShape()

### Configuration options
> click.NewBuilder(click.WithXxx(), ...) OR builder.SetOptions()(click.WithXxx(), ...)
- click.WithImageSize(option.Size)
- click.WithRangeLen(option.RangeVal) 
- click.WithRangeAnglePos([]option.RangeVal) 
- click.WithRangeSize(option.RangeVal)
- click.WithRangeColors([]string) 
- click.WithDisplayShadow(bool) 
- click.WithShadowColor(string) 
- click.WithShadowPoint(option.Point)
- click.WithImageAlpha(float32) 
- click.WithUseShapeOriginalColor(bool)

- click.WithThumbImageSize(option.Size)
- click.WithRangeVerifyLen(option.RangeVal)
- click.WithDisabledRangeVerifyLen(bool)
- click.WithRangeThumbSize(option.RangeVal)
- click.WithRangeThumbColors([]string)
- click.WithRangeThumbBgColors([]string)
- click.WithIsThumbNonDeformAbility(bool)
- click.WithThumbBgDistort(int) 
- click.WithThumbBgCirclesNum(int) 
- click.WithThumbBgSlimLineNum(int) 


### Set resources
> builder.SetResources(click.WithXxx(), ...)
- click.WithChars([]string) 
- click.WithShapes(map[string]image.Image) 
- click.WithFonts([]*truetype.Font) 
- click.WithBackgrounds([]image.Image) 
- click.WithThumbBackgrounds([]image.Image) 

### Captcha Data
- GetData() map[int]*Dot
- GetMasterImage() imagedata.JPEGImageData
- GetThumbImage() imagedata.PNGImageData

<br />

## 🖖 Slide Mode Captcha
### Quick use
```go
package main

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"io/ioutil"

	"github.com/lukhaiminh/go-captcha/v2/base/option"
	"github.com/lukhaiminh/go-captcha/v2/slide"
	"github.com/lukhaiminh/go-captcha/v2/base/codec"
)

var slideTileCapt slide.Captcha

func init() {
	builder := slide.NewBuilder()

	// You can use preset material resources：https://github.com/lukhaiminh/go-captcha-assets
	bgImage, err := loadPng("../resources/bg.png")
	if err != nil {
		log.Fatalln(err)
	}

	bgImage1, err := loadPng("../resources/bg1.png")
	if err != nil {
		log.Fatalln(err)
	}

	graphs := getSlideTileGraphArr()

	builder.SetResources(
		slide.WithGraphImages(graphs),
		slide.WithBackgrounds([]image.Image{
			bgImage,
			bgImage1,
		}),
	)

	slideTileCapt = builder.Make()
}

func getSlideTileGraphArr() []*slide.GraphImage {
	tileImage1, err := loadPng("../resources/tile-1.png")
	if err != nil {
		log.Fatalln(err)
	}

	tileShadowImage1, err := loadPng("../resources/tile-shadow-1.png")
	if err != nil {
		log.Fatalln(err)
	}
	tileMaskImage1, err := loadPng("../resources/tile-mask-1.png")
	if err != nil {
		log.Fatalln(err)
	}

	return []*slide.GraphImage{
		{
			OverlayImage: tileImage1,
			ShadowImage:  tileShadowImage1,
			MaskImage:    tileMaskImage1,
		},
	}
}

func main() {
	captData, err := slideTileCapt.Generate()
	if err != nil {
		log.Fatalln(err)
	}

	blockData := captData.GetData()
	if blockData == nil {
		log.Fatalln(">>>>> generate err")
	}

	block, _ := json.Marshal(blockData)
	fmt.Println(">>>>>", string(block))

	err = captData.GetMasterImage().SaveToFile("../resources/master.jpg", option.QualityNone)
	if err != nil {
		fmt.Println(err)
	}
	err = captData.GetTileImage().SaveToFile("../resources/thumb.png")
	if err != nil {
		fmt.Println(err)
	}
}

func loadPng(p string) (image.Image, error) {
	imgBytes, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return codec.DecodeByteToPng(imgBytes)
}
```


### Create instance method
- builder.Make()
- builder.MakeWithRegion() 


### Configuration options
> slide.NewBuilder(slide.WithXxx(), ...) OR builder.SetOptions(slide.WithXxx(), ...)
- slide.WithImageSize(*option.Size)
- slide.WithImageAlpha(float32) 
- slide.WithRangeGraphSize(val option.RangeVal) 
- slide.WithRangeGraphAnglePos([]option.RangeVal) 
- slide.WithGenGraphNumber(val int)
- slide.WithEnableGraphVerticalRandom(val bool) 
- slide.WithRangeDeadZoneDirections(val []DeadZoneDirectionType) 


### Set resources
builder.SetResources(slide.WithXxx(), ...)
- slide.WithBackgrounds([]image.Image) 
- slide.WithGraphImages(images []*GraphImage)

### Captcha Data
- GetData() *Block
- GetMasterImage() imagedata.JPEGImageData
- GetTileImage() imagedata.PNGImageData


<br />

## 🖖 Rotate Mode Captcha
### Quick use
```go
package main

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"io/ioutil"

	"github.com/lukhaiminh/go-captcha/v2/rotate"
	"github.com/lukhaiminh/go-captcha/v2/base/codec"
)

var rotateCapt rotate.Captcha

func init() {
	builder := rotate.NewBuilder()

	// You can use preset material resources：https://github.com/lukhaiminh/go-captcha-assets
	bgImage, err := loadPng("../resources/bg.png")
	if err != nil {
		log.Fatalln(err)
	}

	bgImage1, err := loadPng("../resources/bg1.png")
	if err != nil {
		log.Fatalln(err)
	}

	builder.SetResources(
		rotate.WithImages([]image.Image{
			bgImage,
			bgImage1,
		}),
	)

	rotateCapt = builder.Make()
}

func main() {
	captData, err := rotateCapt.Generate()
	if err != nil {
		log.Fatalln(err)
	}

	blockData := captData.GetData()
	if blockData == nil {
		log.Fatalln(">>>>> generate err")
	}

	block, _ := json.Marshal(blockData)
	fmt.Println(">>>>>", string(block))

	err = captData.GetMasterImage().SaveToFile("../resources/master.png")
	if err != nil {
		fmt.Println(err)
	}
	err = captData.GetThumbImage().SaveToFile("../resources/thumb.png")
	if err != nil {
		fmt.Println(err)
	}
}

func loadPng(p string) (image.Image, error) {
	imgBytes, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return codec.DecodeByteToPng(imgBytes)
}
```


### Create instance method
- builder.Make()


### Configuration options
> rotate.NewBuilder(rotate.WithXxx(), ...) OR builder.SetOptions(rotate.WithXxx(), ...)
- rotate.WithImageSquareSize(val int) 
- rotate.WithRangeAnglePos(vals []option.RangeVal)
- rotate.WithRangeThumbImageSquareSize(val []int) 
- rotate.WithThumbImageAlpha(val float32)


### Set resources
builder.SetResources(rotate.WithXxx(), ...)
- rotate.WithBackgrounds([]image.Image)

### Captcha Data
- GetData() *Block
- GetMasterImage() imagedata.PNGImageData
- GetThumbImage() imagedata.PNGImageData

<br/>

## Captcha Image Data
### JPEGImageData object method
- Get() image.Image
- ToBytes() []byte
- ToBytesWithQuality(imageQuality int) []byte 
- ToBase64() string
- ToBase64WithQuality(imageQuality int) string
- SaveToFile(filepath string, quality int) error


### PNGImageData object method
- Get() image.Image 
- ToBytes() []byte
- ToBase64() string 
- SaveToFile(filepath string) error

<br/>

## Install package
- <p>Web Native ✔</p>
- <p>Vue ✔</p>
- <p>React ✔</p>
- <p>Angular ✔</p>
- <p>Svelte ✔</p>
- <p>Solid ✔</p>
- <p>MinProgram</p>
- <p>UniApp</p>
- <p>Android App</p>
- <p>IOS App</p>
- <p>Flutter App</p>
- <p>... </p>

<br/>

## LICENSE
Go Captcha source code is licensed under the Apache Licence, Version 2.0 [http://www.apache.org/licenses/LICENSE-2.0.html](http://www.apache.org/licenses/LICENSE-2.0.html)
