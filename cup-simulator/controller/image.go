// package main

// import "image"
// import "image/color"
// import "image/png"
// import "os"

// func main() {
// 	// Create an 100 x 50 image
// 	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

// 	// Draw a red dot at (2, 3)
// 	img.Set(2, 3, color.RGBA{255, 0, 0, 255})

// 	// Save to out.png
// 	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
// 	defer f.Close()
// 	png.Encode(f, img)
// }

//

// package main

// import (
// 	"image"
// 	"image/draw"
// 	"image/jpeg"
// 	"image/png"
// 	"os"
// )

// func main() {
// 	imgb, _ := os.Open("input.jpg")
// 	img, _ := jpeg.Decode(imgb)
// 	defer imgb.Close()

// 	wmb, _ := os.Open("watermark.png")
// 	watermark, _ := png.Decode(wmb)
// 	defer wmb.Close()

// 	offset := image.Pt(100, 100)
// 	b := img.Bounds()
// 	m := image.NewRGBA(b)
// 	draw.Draw(m, b, img, image.ZP, draw.Src)
// 	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

// 	imgw, _ := os.Create("output.jpg")
// 	jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
// 	defer imgw.Close()
// }

// package main

// import (
// 	"image"
// 	"image/color"
// 	"image/png"
// 	"os"

// 	"golang.org/x/image/font"
// 	"golang.org/x/image/font/basicfont"
// 	"golang.org/x/image/math/fixed"
// )

// func addLabel(img *image.RGBA, x, y int, label string) {
// 	col := color.RGBA{200, 100, 0, 255}
// 	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

// 	d := &font.Drawer{
// 		Dst:  img,
// 		Src:  image.NewUniform(col),
// 		Face: basicfont.Face7x13,
// 		Dot:  point,
// 	}
// 	d.DrawString(label)
// }

// func main() {
// 	img := image.NewRGBA(image.Rect(0, 0, 1200, 630))

// 	addLabel(img, 20, 30, "Hello Go")

// 	f, err := os.Create("hello-go.png")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
// 	if err := png.Encode(f, img); err != nil {
// 		panic(err)
// 	}
// }
package main

import (
	"log"

	"github.com/fogleman/gg"
)
type mapxy struct{
	x float64
	y float64
}
 

func main() {

	mapv := []mapxy{
		{x : 0, y: 0,},
		{
			x : 0,
			y : 100,
		},
		{
			x: 0,
			y: 200,
		},
		{
			x: 0,
			y: 300,
		},
		{
			x: 100,
			y: 0,
		},
		{
			x: 100,
			y: 100,
		},
		{
			x: 100,
			y: 200,
		},
		{
			x: 100,
			y: 300,
		}
		
	}

	const X = 1200
	const Y = 1200
	im, err := gg.LoadImage(`C:\Users\vggarcia\go\src\generate-image\cup-simulator\src.jpg`)
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(X, Y)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace(`C:\Users\vggarcia\go\src\generate-image\cup-simulator\Roboto-Black.ttf`, 30); err != nil {
		panic(err)
	}

	dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringAnchored("Hello, world!", 200, 200, 0.5, 0.5)
	dc.DrawStringAnchored("Grupo A", 80, 80, 0.5, 0.5)
	dc.DrawStringAnchored("BRASIL", 50, 50, 0.5, 0.5)

	// dc.DrawStringAnchored("Grupo B", S/2, S/3, 0.5, 0.5)
	// //dc.DrawStringAnchored("ALEMANHA", S/1, S/3, 0.5, 0.5)

	// dc.DrawStringAnchored("Grupo C", S/2, S/3, 0.5, 0.5)
	//dc.DrawStringAnchored("RUSSIA", S/2, S/3, 0.5, 0.5)

	//dc.DrawStringAnchored("Grupo D", S/7, S/3, 0.5, 0.5)

	// dc.DrawStringAnchored("Grupo E", S/7, S/3, 0.5, 0.5)

	// dc.DrawStringAnchored("Grupo F", S/7, S/3, 0.5, 0.5)

	// dc.DrawStringAnchored("Grupo G", S/7, S/3, 0.5, 0.5)

	// dc.DrawStringAnchored("Grupo H", S/7, S/3, 0.5, 0.5)

	dc.Clip()
	dc.SavePNG("out.png")
}

// package main

// import (
// 	"fmt"
// 	"image"
// )

// func main() {

// 	r := image.Rect(2, 1, 5, 5)
// 	// Dx and Dy return a rectangle's width and height.

// 	r.DrawStringAnchored
// 	fmt.Println(r.Dx(), r.Dy(), image.Pt(0, 0).In(r)) // prints 3 4 false
// }

// package main

// import (
// 	"fmt"
// 	"image"
// 	"image/color"
// 	"image/jpeg"
// 	"image/png"
// 	"os"

// 	"golang.org/x/image/font"
// 	"golang.org/x/image/font/basicfont"
// 	"golang.org/x/image/math/fixed"
// )

// func init() {
// 	// damn important or else At(), Bounds() functions will
// 	// caused memory pointer error!!
// 	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
// }

// func main() {
// 	imgfile, err := os.Open(`C:\Users\vggarcia\go\src\generate-image\cup-simulator\src.jpg`)

// 	if err != nil {
// 		fmt.Println("img.jpg file not found!")
// 		os.Exit(1)
// 	}

// 	defer imgfile.Close()

// 	// get image height and width with image/jpeg
// 	// change accordinly if file is png or gif

// 	imgCfg, _, err := image.DecodeConfig(imgfile)

// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	width := imgCfg.Width
// 	height := imgCfg.Height

// 	fmt.Println("Width : ", width)
// 	fmt.Println("Height : ", height)

// 	// we need to reset the io.Reader again for image.Decode() function below to work
// 	// otherwise we will  - panic: runtime error: invalid memory address or nil pointer dereference
// 	// there is no build in rewind for io.Reader, use Seek(0,0)
// 	imgfile.Seek(0, 0)

// 	// get the image
// 	img, _, err := image.Decode(imgfile)
// 	b := img.Bounds()
// 	imgSet := image.NewRGBA(b)
// 	addLabel(imgSet, 20, 30, "Hello Go")

// 	f, err := os.Create("hello-goss.png")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
// 	if err := png.Encode(f, img); err != nil {
// 		panic(err)
// 	}

// }
// func addLabel(img *image.RGBA, x, y int, label string) {

// 	col := color.RGBA{0xdd, 0x00, 0x93, 0xff}
// 	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

// 	d := &font.Drawer{
// 		Dst:  img,
// 		Src:  image.NewUniform(col),
// 		Face: basicfont.Face7x13,
// 		Dot:  point,
// 	}
// 	d.DrawString(label)
// }
