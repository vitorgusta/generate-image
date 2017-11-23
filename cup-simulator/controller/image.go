package controller

import (
	"fmt"
	"log"

	"github.com/fogleman/gg"
)

type mapxy struct {
	x float64
	y float64
}

func generateImage() {

	type opt struct {
		shortnm      byte
		longnm, help string
		needArg      bool
	}

	var basenameOpts []mapxy

	basenameOpts = []mapxy{
		mapxy{
			x: 0,
			y: 0,
		},
		mapxy{
			x: 0,
			y: 100,
		},
		mapxy{
			x: 0,
			y: 200,
		},
		mapxy{
			x: 0,
			y: 300,
		},
		mapxy{
			x: 100,
			y: 0,
		},
		mapxy{
			x: 100,
			y: 100,
		},
		mapxy{
			x: 100,
			y: 200,
		},
		mapxy{
			x: 100,
			y: 300,
		},
	}
	fmt.Println("basenameOpts", basenameOpts)
	const X = 1200
	const Y = 630
	im, err := gg.LoadImage(`/Users/vitorlevy/Documents/generate-image/cup-simulator/Esporte_Simulador_Copa_compartilhamento.png`)
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(X, Y)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace(`/Users/vitorlevy/Documents/generate-image/cup-simulator/Roboto-Black.ttf`, 30); err != nil {
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
