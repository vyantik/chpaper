package main

import (
	"flag"
	"fmt"
)

func main() {
	imgPath := flag.String("path", "", "path to the image you want to change the wallpaper insted of current one")
	flag.Parse()

	if *imgPath == "" {
		fmt.Println("pass the image path like 'chpaper --path path/to/image.png'")
		return
	}

	fmt.Println(*imgPath)
}
