package main

/*
 * Go-transform takes your image and swirls it into a neat output gif.
 * Remember to open your gif in chrome or other live gif viewer.
 */

/*
 * unicode
  net
  image
  bytes
  encoding
 */


import (
	"fmt"
	"image"
	"os"
	"image/gif"
)

const APP_NAME = "go-transform"
const IMG_DIR = "static/"
const OUT_DIR = "out/"

func playGif(imageFrames []string) {
	// load static image and construct outGif
	outGif := &gif.GIF{}
	for _, name := range imageFrames {
		f, _ := os.Open(name)
		inGif, _ := gif.Decode(f)
		f.Close()

		outGif.Image = append(outGif.Image, inGif.(*image.Paletted))
		outGif.Delay = append(outGif.Delay, 0)
	}

	for _, name := range imageFrames {
		f, _ := os.Open(name)
		inGif, _ := gif.Decode(f)
		f.Close()

		outGif.Image = append(outGif.Image, inGif.(*image.Paletted))
		outGif.Delay = append(outGif.Delay, 0)
	}

	// save to out.gif
	f, _ := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, outGif)
}

func main() {
	fmt.Printf("%s is transforming your image...", APP_NAME)
	// Decode the JPEG data. If reading from file, create a reader with
	files := []string{"static/g1.gif", "static/g2.gif","static/g3.gif", "static/g2.gif"}
	playGif(files)
	fmt.Print("done")
}
