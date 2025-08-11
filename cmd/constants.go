package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/chai2010/webp"
)

var mimetypes = []string{"jpeg", "jpg", "png", "webp"}

var decoders = map[string]func(i io.Reader) (image.Image, error){
	"jpeg": jpeg.Decode,
	"jpg":  jpeg.Decode,
	"png":  png.Decode,
	"webp": webp.Decode,
}

const outputPath = "~/.config/hypr/wallpaper/wallpaper.png"

var scripts = []string{
	"~/.config/hypr/set_wal_colors.py",
	"~/.config/kitty/set_wal_colors.py",
}

var processes = []func(){
	restartHyprpaper,
	restartSwaync,
}
