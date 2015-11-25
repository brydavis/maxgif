package main

// import (
// 	"fmt"
// 	"image"
// 	"image/gif"
// 	"image/jpeg"
// 	"io/ioutil"
// 	"os"
// 	"path"
// )

// func main() {
// 	dir, _ := ioutil.ReadDir(".")

// 	for _, f := range dir {
// 		if path.Ext(f.Name()) == ".jpeg" {
// 			// fmt.Println(f.Name())
// 			file, _ := os.Open(f.Name())

// 			j, _ := jpeg.Decode(file)
// 			// imgs = append(imgs, j)

// 			out, err := os.Create(file.Name() + ".gif")

// 			if err != nil {
// 				fmt.Println(err)
// 				os.Exit(1)
// 			}

// 			defer out.Close()

// 			// ok, write out the data into the new GIF file

// 			var opt gif.Options
// 			opt.NumColors = 256

// 			err = gif.Encode(out, j, &opt)
// 			if err != nil {
// 				fmt.Println(err)
// 				os.Exit(1)
// 			}

// 			fmt.Println("Generated image \n")

// 		}

// 	}

// 	var gif_imgs []*image.Paletted

// 	for _, f := range dir {
// 		if path.Ext(f.Name()) == ".gif" {
// 			fmt.Println(f.Name())

// 			file, _ := os.Open(f.Name())

// 			g, _ := gif.DecodeAll(file)

// 			g_img := new(image.Paletted)

// 			if len(g.Image) > 0 {
// 				g_img = g.Image[0]
// 			}

// 			gif_imgs = append(gif_imgs, g_img)

// 		}

// 	}

// 	g := new(gif.GIF)

// 	g.Delay = []int{100, 100}
// 	g.LoopCount = 1000
// 	g.Image = gif_imgs

// 	out, err := os.Create("out.gif")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer out.Close()

// 	e := gif.EncodeAll(out, g)
// 	if e != nil {
// 		fmt.Print(e)
// 	}
// }
