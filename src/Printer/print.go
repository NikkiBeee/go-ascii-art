package main

import (
    "fmt"
    "os"
	"log"
	// "io/ioutil"
	// "strconv"
	// "reflect"
	// "image"
	"image/jpeg"
	// "image/png"
	"image/color"
	"./resize"
)



func main() {
	i, err := os.Open("../AsciiArt/BlackWhiteRubberDuckJPEG.jpg")
	if err != nil {
		log.Fatal("Err in open: ", err)
	}
	imageData, err := jpeg.Decode(i)
	if err != nil {
		log.Fatal("Err in image: ", err)
	}

	var ratio, w, h float64
	ratio = (float64(imageData.Bounds().Max.Y) / float64(imageData.Bounds().Max.X))

	w = 70
	h = (w / 2) * ratio

	imageResize := resize.Resize(uint(w), uint(h), imageData, resize.MitchellNetravali)

	fmt.Print("H", h)
	fmt.Print("ratio", ratio)

	// var size int64 = len([]imageData)
	// fmt.Print("imageData: ", imageData.Bounds())
	// fmt.Print("imageData: ", imageData.At(0,0))
	// fmt.Print("imageType: ", imageType)
	// fileInfo, _ := i.Stat()
  	// var size int64 = fileInfo.Size()
	// bytes := make([]byte, size)


	// fmt.Print(imageData)

	
	
	for y := 0; y < int(h); y++ {
		for x := 0; x < int(w); x++ {
			c := color.GrayModel.Convert(imageResize.At(x,y)).(color.Gray)
			if  c.Y >  200 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Print("\n")
	}
	
}
