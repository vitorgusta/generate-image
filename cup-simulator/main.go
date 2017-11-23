package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fogleman/gg"
	"github.com/labstack/echo"
	"github.com/tylerb/graceful"
)

type mapxy struct {
	x float64
	y float64
}

type Groups struct {
	Groups [][]struct {
		Image string `json:"image"`
		Name  string `json:"name"`
	} `json:"groups"`
}

func main() {

	var port = flag.String("p", "8001", "server port number")

	flag.Parse()

	log.Printf("Starting Rest API service on port %s\n", *port)

	e := echo.New()
	e.Server.Addr = ":" + *port

	r := e.Group("/api/v1")
	r.POST("/something", generateTemplate)

	graceful.ListenAndServe(e.Server, 5*time.Second)
}

func generateTemplate(c echo.Context) (err error) {

	groups1 := new(Groups)

	if err := c.Bind(groups1); err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	img := generateImage(*groups1)
	fmt.Println("JSON", groups1)

	return c.JSON(http.StatusCreated, img)
}

func generateImage(groups Groups) string {
	var basenameOpts []mapxy
	basenameOpts = []mapxy{
		mapxy{
			x: 150,
			y: 230,
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
	if err := dc.LoadFontFace(`/Users/vitorlevy/Documents/generate-image/cup-simulator/FiraSans-Medium.ttf`, 30); err != nil {
		panic(err)
	}

	dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
	dc.DrawImage(im, 0, 0)

	fmt.Println(basenameOpts)
	for i, group := range groups.Groups {
		for _, time := range group {
			fmt.Println("vitor", i, time.Image, time.Image)
			dc.DrawStringAnchored(time.Name, 250, 330, 0.5, 0.5)
		}
	}
	dc.Clip()
	dc.SavePNG("out1.png")
	return "umg"
}
