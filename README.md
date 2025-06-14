# chpaper

A command-line utility for changing desktop wallpapers and automatically configuring color schemes in Linux.

## Description

`chpaper` is a tool that makes it easy to change your desktop wallpaper and automatically configure your system's color scheme. It integrates with [pywal](https://github.com/dylanaraps/pywal) to generate color schemes based on your wallpaper.

## Requirements

- Go 1.24.4 or higher
- pywal
- Python (for configuration scripts)

## Usage

```bash
chpaper --path /path/to/image.png
```

### Supported Image Formats

- PNG
- JPEG/JPG
- WebP

## Features

- Image conversion to PNG format
- Automatic color scheme generation using pywal
- Restart of necessary processes to apply changes
- Multi-threaded processing for fast operation

## License

MIT
