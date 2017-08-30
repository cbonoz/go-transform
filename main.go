package main

import (
	"fmt"
	"image"
	"os"
	"image/gif"
	"github.com/disintegration/imaging"
	"image/color"
	"log"
	"image/jpeg"
	"bytes"
)

/*
 * Go-transform takes your image and swirls it into a neat output gif.
 * Remember to open your gif in chrome or other live gif viewer.
 */

/*
 * Library basket #6:
 * unicode
 * net
 * image
 * bytes
 * encoding
 */

const APP_NAME = "go-transform"
const INPUT_IMG_NAME = "./static/input.jpg"
const OUTPUT_FILE = "out.gif"

const FRAME_IMG_NAME = "frame"
const FRAME_ROTATION = 6 // degrees
const IMG_SIZE = 256

func playGif(files []string) {
	// load static image and construct outGif
	outGif := &gif.GIF{}
	for _, name := range files {
		simage, _ := imaging.Open(name)

		buf := bytes.Buffer{}
		if err := gif.Encode(&buf, simage, nil); err != nil {
			log.Printf("Skipping file %s due to error in gif encoding:%s", name, err)
			continue
		}

		tmpimg, err := gif.Decode(&buf)
		if err != nil {
			log.Printf("Skipping file %s due to error reading the temporary gif :%s", name, err)
			continue
		}

		outGif.Image = append(outGif.Image, tmpimg.(*image.Paletted))
		outGif.Delay = append(outGif.Delay, 0)
	}

	// save to out.gif
	f, _ := os.OpenFile(OUTPUT_FILE, os.O_WRONLY | os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, outGif)
}

// Write the image.
func saveImage(img image.Image, dest string) {
	toimg, _ := os.Create(dest)
	defer toimg.Close()
	jpeg.Encode(toimg, img, &jpeg.Options{jpeg.DefaultQuality})
}

func generateNextImage(img1 image.Image) *image.NRGBA {
	nextImg := imaging.Rotate(img1, FRAME_ROTATION, color.White)
	nextImg = imaging.Resize(nextImg, IMG_SIZE, IMG_SIZE, imaging.Lanczos)
	return nextImg
}

func generateImageFrames(inputFile string) []string {
	img, err := imaging.Open(inputFile)
	if err != nil {
		log.Fatalf("Open failed: %v", err)
	}
	img = imaging.Resize(img, IMG_SIZE, IMG_SIZE, imaging.Lanczos)
	dest := fmt.Sprintf("static/%s%d.jpg", FRAME_IMG_NAME, 0)
	saveImage(img, dest)
	files := []string{dest}

	for i := 1; i <= 25; i++ {
		dest = fmt.Sprintf("static/%s%d.jpg", FRAME_IMG_NAME, i)
		img = generateNextImage(img)
		saveImage(img, dest)
		files = append(files, dest)
	}

	return files
}

func main() {
	fmt.Printf("%s is transforming your image...", APP_NAME)

	inputFile := INPUT_IMG_NAME

	fmt.Printf("%s is transforming your image: %s...", APP_NAME, inputFile)

	// Generate the gif frames from the input file.
	frames := generateImageFrames(inputFile)
	fmt.Printf("\nframes: %s", frames)

	playGif(frames)
	fmt.Print("\ndone")
}
