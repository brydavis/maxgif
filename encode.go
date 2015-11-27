package main

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"io/ioutil"
	"os"
	"path"

	rs "github.com/brydavis/resize" // "../resize"
)

func main() {
	dir, _ := ioutil.ReadDir(".")

	var imgs []*image.Paletted
	for _, file := range dir {
		f, _ := os.Open(file.Name())
		defer f.Close()

		if ext := path.Ext(file.Name()); ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			fmt.Printf("File not an image: %s\n", file.Name())
		} else {
			img, err := rs.ResizePixels(f, 400, 400)
			if err != nil {
				fmt.Println(err)
			}

			var opt gif.Options
			opt.NumColors = 256

			buf := new(bytes.Buffer)
			gif.Encode(buf, img, &opt)

			g, _ := gif.DecodeAll(buf)
			imgs = append(imgs, g.Image[0])
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
