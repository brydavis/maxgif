package main

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path"
)

func main() {

	dir, _ := ioutil.ReadDir(".")

	var imgs []*image.Paletted

	for _, file := range dir {
		switch ext := path.Ext(file.Name()); ext {
		case ".jpg", ".jpeg":
			f, _ := os.Open(file.Name())
			defer f.Close()

			img, _ := jpeg.Decode(f)

			var opt gif.Options
			opt.NumColors = 256

			buf := new(bytes.Buffer)
			gif.Encode(buf, img, &opt)

			g, _ := gif.DecodeAll(buf)

			pic := g.Image[0]
			size := pic.Rect.Size()
			fmt.Printf("%s %v\n", file.Name(), size)

			// rect:= image.Rect(0, 0, 500, 500)
			rx, ry := 0, 0

			for x := 0; x < pic.Rect.Max.X; x++ {

				for y := 0; y < pic.Rect.Max.X; y++ {

					if y%2 != 0.0 && x%2 != 0.0 {

						pic.Set(rx, ry, pic.At(x, y))
						rx++
						ry++
					}

				}
			}
			imgs = append(imgs, pic)

		default:
			fmt.Println(ext)
		}
	}

	g := new(gif.GIF)

	g.Image = imgs
	g.Delay = evenDelay(100, len(g.Image))
	g.LoopCount = 1000

	var opt gif.Options
	opt.NumColors = 256

	out, err := os.Create("./output.gif")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = gif.EncodeAll(out, g)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Generated image to output.gif \n")
}

func evenDelay(delay, imgs int) (res []int) {
	for i := 0; i < imgs; i++ {
		res = append(res, delay)
	}
	return
}
