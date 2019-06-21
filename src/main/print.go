package main

// import (
//     "fmt"
//     "os"
// 	"log"
// 	// "io/ioutil"
// 	// "strconv"
// 	// "reflect"
// 	// "image"
// 	"image/jpeg"
// 	// "image/png"
// 	"image/color"
// 	"./resize"
// )



// func Print(image string) {
// 	image = "../AsciiArt/goGopher2.jpg"
// 	i, err := os.Open(image)
// 	if err != nil {
// 		log.Fatal("Err in open: ", err)
// 	}
// 	imageData, err := jpeg.Decode(i)
// 	if err != nil {
// 		log.Fatal("Err in image: ", err)
// 	}

// 	var ratio, w, h float64
// 	ratio = (float64(imageData.Bounds().Max.Y) / float64(imageData.Bounds().Max.X))

// 	w = 70
// 	h = (w / 2) * ratio

// 	imageResize := resize.Resize(uint(w), uint(h), imageData, resize.MitchellNetravali)

	
// 	for y := 0; y < int(h); y++ {
// 		for x := 0; x < int(w); x++ {
// 			c := color.GrayModel.Convert(imageResize.At(x,y)).(color.Gray)
// 			if  c.Y >  200 {
// 				fmt.Print(" ")
// 			} else if c.Y >  100 {
// 				fmt.Print("*")
// 			} else if c.Y >  50 {
// 				fmt.Print(".")
// 			} else {
// 				fmt.Print("'")
// 			}
// 		}
// 		fmt.Print("\n")
// 	}
	
// }
