package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fogleman/gg"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"github.com/tylerb/graceful"
)

type mapxy struct {
	x      float64
	y      float64
	imageX int
	imageY int
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
	r.POST("/generateTemplate", generateTemplate)

	graceful.ListenAndServe(e.Server, 5*time.Second)
}

//Metodo que gera a imagem e retorna o caminho onde ela foi salca
func generateTemplate(c echo.Context) (err error) {

	groups := new(Groups)

	if err := c.Bind(groups); err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	img := generateImage(*groups)

	return c.JSON(http.StatusCreated, img)
}

func generateImage(groups Groups) string {
	var basenameOpts []mapxy
	basenameOpts = []mapxy{
		mapxy{
			x:      125,
			y:      230,
			imageX: 88,
			imageY: 215,
		},
		mapxy{
			x:      410,
			y:      230,
			imageX: 375,
			imageY: 215,
		},
		mapxy{
			x:      700,
			y:      230,
			imageX: 665,
			imageY: 215,
		},
		mapxy{
			x:      990,
			y:      230,
			imageX: 955,
			imageY: 215,
		},
		mapxy{
			x:      125,
			y:      400,
			imageX: 88,
			imageY: 386,
		},
		mapxy{
			x:      410,
			y:      400,
			imageX: 375,
			imageY: 386,
		},
		mapxy{
			x:      700,
			y:      400,
			imageX: 665,
			imageY: 386,
		},
		mapxy{
			x:      990,
			y:      400,
			imageX: 955,
			imageY: 386,
		},
	}

	const X = 1200
	const Y = 630
	//carrega a imagem de background
	im, err := gg.LoadImage(`C:\Users\vggarcia\go\src\generate-image\cup-simulator\Esporte_Simulador_Copa_compartilhamento.png`)
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(X, Y)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	//seta o a fonte que vai ser usada nos nomes dos times
	if err := dc.LoadFontFace(`C:\Users\vggarcia\go\src\generate-image\cup-simulator\FiraSans-BlackItalic.ttf`, 16); err != nil {
		panic(err)
	}

	dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
	dc.DrawImage(im, 0, 0)
	//seta a cor do texto
	dc.SetHexColor("#FFFFFF")

	//percorre o objeto com os times seta as imagens e os nomes no template
	for i, group := range groups.Groups {
		for _, time := range group {
			slug, _ := gg.LoadImage("C:/Users/vggarcia/go/src/generate-image/cup-simulator/Bandeiras/" + "franca" + ".png")
			dc.DrawImage(slug, int(basenameOpts[i].imageX), int(basenameOpts[i].imageY))
			dc.DrawString(time.Name, basenameOpts[i].x, basenameOpts[i].y)
			basenameOpts[i].imageY = basenameOpts[i].imageY + 40
			basenameOpts[i].y = basenameOpts[i].y + 40
		}
	}

	u1 := uuid.NewV4()
	fmt.Println(u1)

	dc.Clip()
	dc.SavePNG("//bandeirantes.com.br/webcontent/Portal_Band/S_Apiportal/Images/simulador-copa-do-mundo/copadomundo2018-" + u1.String() + ".png")
	return "//bandeirantes.com.br/webcontent/Portal_Band/S_Apiportal/Images/simulador-copa-do-mundo/copadomundo2018-" + u1.String() + ".png"
}
