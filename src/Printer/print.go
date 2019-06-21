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
)



func main() {
	i, err := os.Open("../AsciiArt/dot.jpg")
	if err != nil {
		log.Fatal("Err in open: ", err)
	}
	imageData, err := jpeg.Decode(i)
	if err != nil {
		log.Fatal("Err in image: ", err)
	}

	// var size int64 = len([]imageData)
	// fmt.Print("imageData: ", imageData.Bounds())
	// fmt.Print("imageData: ", imageData.At(0,0))
	// fmt.Print("imageType: ", imageType)
	// fileInfo, _ := i.Stat()
  	// var size int64 = fileInfo.Size()
	// bytes := make([]byte, size)


	// fmt.Print(imageData)

	
	
	for y := imageData.Bounds().Min.Y; y < imageData.Bounds().Max.Y; y++ {
		for x := imageData.Bounds().Min.X; x < imageData.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(imageData.At(x,y)).(color.Gray)
			if  c.Y >  50 {
				fmt.Print("  ")
			} else {
				fmt.Print("**")
			}
		}
		fmt.Print("\n")
	}
	
}
