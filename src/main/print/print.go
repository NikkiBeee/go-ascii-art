package print

import (
	"fmt"
	"log"
	"os"

	// "io/ioutil"
	// "strconv"
	// "reflect"
	// "image"
	"image/jpeg"
	// "image/png"
	"image/color"

	"../resize"
)

func Printer(image string) string {
	// fmt.Println("Print is running")
	i, err := os.Open(image)
	if err != nil {
		log.Fatal("Err in open: ", err)
	}
	imageData, err := jpeg.Decode(i)
	if err != nil {
		log.Fatal("Err in image: ", err)
	}

	var ratio, w, h float64
	ratio = (float64(imageData.Bounds().Max.Y) / float64(imageData.Bounds().Max.X))

	w = 300
	h = (w) * ratio

	imageResize := resize.Resize(uint(w), uint(h), imageData, resize.MitchellNetravali)

	var newString string = ""

	for y := 0; y < int(h); y++ {
		for x := 0; x < int(w); x++ {
			c := color.GrayModel.Convert(imageResize.At(x, y)).(color.Gray)
			if c.Y > 200 {
				newString += "  "
			} else if c.Y > 180 {
				newString += " *"
			} else if c.Y > 100 {
				newString += "**"
			} else if c.Y > 50 {
				newString += " ."
			} else {
				newString += ".."
			}
		}
		newString += "\n"
	}

	fmt.Print(newString)
	return newString
}
