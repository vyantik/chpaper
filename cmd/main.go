package main

import (
	"flag"
	"image/png"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

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

	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(imgPath), "."))
	if ext == "" {
		log.Fatalln("is not image")
	}
	if isValidImage := isImage(ext); !isValidImage {
		log.Fatalln("is not image")
	}

	convertImageToPNG(imgPath, outputPath)
	generateNewColors(imgPath)
	var wg sync.WaitGroup
	for _, scriptPath := range scripts {
		wg.Add(1)
		go func() {
			defer wg.Done()
			extendedPath, err := expandTilde(scriptPath)
			if err != nil {
				log.Fatalln("can't get path from tilda contained path")
			}
			execPython(extendedPath)
		}()
	}
	wg.Wait()

	for _, processToRestart := range processes {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processToRestart()
		}()
	}

	wg.Wait()

	log.Println("All passed good... chill bro :)")
}

func convertImageToPNG(inputPath, outputPath string) {
	log.Println("outputPath:", outputPath)
	expandedPath, err := expandTilde(outputPath)
	if err != nil {
		log.Fatalln("can't get path from tilda contained path")
	}

	inFile, err := os.Open(inputPath)
	if err != nil {
		log.Fatalln("failed to open input file")
	}
	defer inFile.Close()

	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(inputPath), "."))

	decodeFunc, ok := decoders[ext]
	if !ok {
		log.Fatalln("unsupported input image format.")
	}

	img, err := decodeFunc(inFile)
	if err != nil {
		log.Fatalln("failed to decode image from")
	}

	outFile, err := os.Create(expandedPath)
	if err != nil {
		log.Fatalln("failed to create output file")
	}
	defer outFile.Close()

	err = png.Encode(outFile, img)
	if err != nil {
		log.Fatalln("failed to encode image to PNG")
	}
}

func generateNewColors(wallpaperPath string) {
	if !fileExist(wallpaperPath) {
		log.Fatalln("image doesn't exist")
	}
	expandedPath, err := expandTilde(wallpaperPath)
	if err != nil {
		log.Fatalln("can't get path from tilda contained path")
	}
	cmd := exec.Command("wal", "-i", expandedPath)
	if output, err := cmd.CombinedOutput(); err != nil {
		log.Fatalf("error when exec 'wal': %v\nOutput: %s\n", err, output)
	}
	log.Println("command 'wal' success")
}
