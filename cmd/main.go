package main

import (
	"flag"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/chai2010/webp"
)

var mimetypes = []string{"jpeg", "jpg", "png", "webp"}

var decoders = map[string]func(i io.Reader) (image.Image, error){
	"jpeg": jpeg.Decode,
	"jpg":  jpeg.Decode,
	"png":  png.Decode,
	"webp": webp.Decode,
}

const outputPath = "/home/vyantik/.config/hypr/wallpaper/wallpaper.png"

func main() {
	imgPathPtr := flag.String("path", "", "path to the image you want to change the wallpaper insted of current one")
	flag.Parse()

	if *imgPathPtr == "" {
		log.Fatalln("pass the image path like 'chpaper --path path/to/image.png'")
		return
	}

	imgPath := *imgPathPtr
	if !fileExist(imgPath) {
		log.Fatalln("pass the image, not dir")
	}
	mimetypeFromSplit := strings.Split(imgPath, ".")[1:]
	if len(mimetypeFromSplit) < 1 {
		log.Fatalln("is not image")
	}
	if isValidImage := isImage(mimetypeFromSplit[0]); isValidImage == false {
		log.Fatalln("is not image")
	}

	convertImageToPNG(imgPath, outputPath)
}

func fileExist(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatalln("image doesn't exist through this path")
	}
	if err != nil {
		log.Fatalln("internal error")
	}
	return !info.IsDir()
}

func isImage(filetype string) bool {
	return slices.Contains(mimetypes, filetype)
}

func convertImageToPNG(inputPath, outputPath string) {

	log.Println("outputPath:", outputPath)

	inFile, err := os.Open(inputPath)
	if err != nil {
		log.Fatalln("failed to open input file")
	}
	defer inFile.Close()

	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(inputPath), "."))

	decodeFunc, ok := decoders[ext]
	if !ok {
		log.Fatalf("unsupported input image format.")
	}

	img, err := decodeFunc(inFile)
	if err != nil {
		log.Fatalln("failed to decode image from")
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatalln("failed to create output file")
	}
	defer outFile.Close()

	err = png.Encode(outFile, img)
	if err != nil {
		log.Fatalln("failed to encode image to PNG")
	}
}
