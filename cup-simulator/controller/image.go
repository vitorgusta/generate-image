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
// 	img := image.NewRGBA(image.Rect(0, 0, 300, 100))
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

// import (
// 	"log"

// 	"github.com/fogleman/gg"
// )

// func main() {
// 	const S = 1024
// 	im, err := gg.LoadImage(`C:/Users/vggarcia/go/src/cup-simulator/src.jpg`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	dc := gg.NewContext(S, S)
// 	dc.SetRGB(1, 1, 1)
// 	dc.Clear()
// 	dc.SetRGB(0, 0, 0)
// 	if err := dc.LoadFontFace(`C:/Users/vggarcia/go/src/cup-simulator/Roboto-Black.ttf`, 30); err != nil {
// 		panic(err)
// 	}
// 	dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5)

// 	dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
// 	dc.DrawImage(im, 0, 0)
// 	dc.DrawStringAnchored("Grupo A", S/8, S/3, 0.5, 0.5)
// 	//dc.DrawStringAnchored("BRASIL", S/7, S/3, 0.5, 0.5)

// 	dc.DrawStringAnchored("Grupo B", S/2, S/3, 0.5, 0.5)
// 	//dc.DrawStringAnchored("ALEMANHA", S/1, S/3, 0.5, 0.5)

// 	dc.DrawStringAnchored("Grupo C", S/2, S/3, 0.5, 0.5)
// 	//dc.DrawStringAnchored("RUSSIA", S/2, S/3, 0.5, 0.5)

// 	//dc.DrawStringAnchored("Grupo D", S/7, S/3, 0.5, 0.5)

// 	// dc.DrawStringAnchored("Grupo E", S/7, S/3, 0.5, 0.5)

// 	// dc.DrawStringAnchored("Grupo F", S/7, S/3, 0.5, 0.5)

// 	// dc.DrawStringAnchored("Grupo G", S/7, S/3, 0.5, 0.5)

// 	// dc.DrawStringAnchored("Grupo H", S/7, S/3, 0.5, 0.5)

// 	dc.Clip()
// 	dc.SavePNG("out.png")
// }

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {

	fBg, err := os.Open(`C:/Users/vggarcia/go/src/cup-simulator/src.jpg`)
	defer fBg.Close()
	bg, _, err := image.Decode(fBg)

	fSrc, err := os.Open("arrow1.jpg")
	defer fSrc.Close()
	src, _, err := image.Decode(fSrc)

	fMaskImg, err := os.Open("mask.jpg")
	defer fMaskImg.Close()
	maskImg, _, err := image.Decode(fMaskImg)

	bounds := src.Bounds() //you have defined that both src and mask are same size, and maskImg is a grayscale of the src image. So we'll use that common size.
	mask := image.NewAlpha(bounds)
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			//get one of r, g, b on the mask image ...
			r, _, _, _ := maskImg.At(x, y).RGBA()
			//... and set it as the alpha value on the mask.
			mask.SetAlpha(x, y, color.Alpha{uint8(255 - r)}) //Assuming that white is your transparency, subtract it from 255
		}
	}

	m := image.NewRGBA(bounds)
	draw.Draw(m, m.Bounds(), bg, image.ZP, draw.Src)

	draw.DrawMask(m, bounds, src, image.ZP, mask, image.ZP, draw.Over)

	toimg, _ := os.Create("new.jpeg")
	defer toimg.Close()

	err = jpeg.Encode(toimg, m, nil)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
}
