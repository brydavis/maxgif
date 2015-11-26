package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	dir, _ := ioutil.ReadDir(".")
	for _, item := range dir {
		// img := ResizePercent(item.Name(), 0.25)
		img := ResizePixels(item.Name(), 100, 800)
		out, _ := os.Create(fmt.Sprintf("resized_%s", item.Name()))

		var opt jpeg.Options
		opt.Quality = 1000

		jpeg.Encode(out, img, &opt)
	}
}

func ResizePercent(filename string, percent float64) (i *image.RGBA) {
	switch ext := path.Ext(filename); ext {
	case ".jpg", ".jpeg":
		file, _ := os.Open(filename)
		img, _ := jpeg.Decode(file)

		rect := image.Rect(0, 0, int(float64(img.Bounds().Max.X)*percent), int(float64(img.Bounds().Max.Y)*percent))
		i = image.NewRGBA(rect)

		for iy := 0; iy <= rect.Max.Y; iy += 1 {
			for ix := 0; ix <= rect.Max.X; ix += 1 {
				i.Set(ix, iy, img.At(int(float64(ix)/percent), int(float64(iy)/percent)))
			}
		}
	}

	return
}

func ResizePixels(filename string, x, y int) (i *image.RGBA) {

	switch ext := path.Ext(filename); ext {
	case ".jpg", ".jpeg":
		file, _ := os.Open(filename)
		img, _ := jpeg.Decode(file)

		dx := float64(x) / float64(img.Bounds().Max.X)
		dy := float64(y) / float64(img.Bounds().Max.Y)

		rect := image.Rect(0, 0, x, y)
		i = image.NewRGBA(rect)

		for iy := 0; iy <= rect.Max.Y; iy += 1 {
			for ix := 0; ix <= rect.Max.X; ix += 1 {
				i.Set(ix, iy, img.At(int(float64(ix)/dx), int(float64(iy)/dy)))
			}
		}
	}
	return
}
