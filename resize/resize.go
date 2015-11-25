package main

import (
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	dir, _ := ioutil.ReadDir(".")
	for _, item := range dir {
		switch ext := path.Ext(item.Name()); ext {
		case ".jpg", ".jpeg":
			file, _ := os.Open(item.Name())
			img, _ := jpeg.Decode(file)

			mx := img.Bounds().Max.X
			my := img.Bounds().Max.Y

			rect := image.Rect(0, 0, (mx/2)+1, (my/2)+1)
			// rect := image.Rect(0, 0, mx, my)

			i := image.NewRGBA(rect)

			// x, y := 0, 0
			for ix := 0; ix <= rect.Max.X; ix += 1 {
				for iy := 0; iy <= rect.Max.Y; iy += 1 {
					i.Set(ix, iy, img.At(ix, iy))
					// i.Set(ix, iy, img.At(x, y))
					// y += 4 // x += 2
				}
				// y += 2
				// x += 4

			}

			out, _ := os.Create("half_" + file.Name())

			// var opt gif.Options
			// opt.NumColors = 256

			// g := new(gif.GIF)

			// var palets []*image.Paletted
			// palets = append(palets, i)

			// g.Image = palets
			// g.Delay = []int{100}
			// g.LoopCount = 1000

			// gif.EncodeAll(out, g)
			var opt jpeg.Options
			opt.Quality = 1000

			jpeg.Encode(out, i, &opt)

		}
	}
}
