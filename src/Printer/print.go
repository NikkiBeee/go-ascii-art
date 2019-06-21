package main

import (
    "fmt"
    "os"
	"log"
	// "io/ioutil"
	// "strconv"
	// "reflect"
	"image/jpeg"
	"image/color"
	"./resize"
)



func main() {
	i, err := os.Open("../AsciiArt/dotLarge.jpg")
	if err != nil {
		log.Fatal("Err in open: ", err)
	}
	imageData, err := jpeg.Decode(i)
	if err != nil {
		log.Fatal("Err in image: ", err)
	}

	ratio := imageData.Bounds().Max.Y / imageData.Bounds().Max.X

	h := 15 * ratio

	imageResize := resize.Resize(30, uint(h), imageData, resize.MitchellNetravali)

	// fmt.Print(imageResize)

	// var size int64 = len([]imageData)
	// fmt.Print("imageData: ", imageData.Bounds())
	// fmt.Print("imageData: ", imageData.At(0,0))
	// fmt.Print("imageType: ", imageType)
	// fileInfo, _ := i.Stat()
  	// var size int64 = fileInfo.Size()
	// bytes := make([]byte, size)


	// fmt.Print(imageData)

	
	
	for y := 0; y < h; y++ {
		for x := 0; x < 30; x++ {
			c := color.GrayModel.Convert(imageResize.At(x,y)).(color.Gray)
			if  c.Y >  50 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Print("\n")
	}
	
}
