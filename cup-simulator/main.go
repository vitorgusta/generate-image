package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fogleman/gg"
	fb "github.com/huandu/facebook"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Server.Addr = ":" + *port

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.HEAD},
	}))
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
	im, err := gg.LoadImage("./Esporte_Simulador_Copa_compartilhamento.png")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(X, Y)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	//seta o a fonte que vai ser usada nos nomes dos times
	if err := dc.LoadFontFace("./FiraSans-BlackItalic.ttf", 16); err != nil {
		panic(err)
	}

	dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
	dc.DrawImage(im, 0, 0)
	//seta a cor do texto
	dc.SetHexColor("#FFFFFF")

	//percorre o objeto com os times seta as imagens e os nomes no template
	for i, group := range groups.Groups {
		for _, time := range group {
			slug, _ := gg.LoadImage("//bandeirantes.com.br/webcontent/Sites_Old/S_Esporte/futebol/copa2018/simulador/images/Bandeiras/" + time.Image)
			dc.DrawImage(slug, int(basenameOpts[i].imageX), int(basenameOpts[i].imageY))
			dc.DrawString(time.Name, basenameOpts[i].x, basenameOpts[i].y)
			basenameOpts[i].imageY = basenameOpts[i].imageY + 40
			basenameOpts[i].y = basenameOpts[i].y + 40
		}
	}

	u1 := uuid.NewV4()

	dc.Clip()
	dc.SavePNG("//bandeirantes.com.br/webcontent/Sites_Old/S_Esporte/futebol/copa2018/simulador/images/timesheets/" + u1.String() + ".png")
	sharedImageFacebook(u1.String())

	return u1.String() + ".png"
}

func sharedImageFacebook(u1 string) {
	imageFacebook := "http://esporte.band.uol.com.br/futebol/copa2018/simulador/images/timesheets/" + u1 + ".png"
	res, e := fb.Post("/me/feed", fb.Params{
		"type":         "link",
		"name":         "test news is here",
		"caption":      "The caption of a link in the post ",
		"picture":      imageFacebook,
		"link":         imageFacebook,
		"description":  "Simulador Copa 2018",
		"access_token": "EAAAATkB4Tx4BANfmFZC2Ti3bNjaZA5TbGrE9hL4vOExkCA6KExcKwTgUZBPZCL9E6sCvGpdlZAfQI0u5ZAixDqxa15feDzRoxZAE0CfEbk0hfDbyY2A6JijP758Pr0ANDZC7QqNQJexGcXzhOkDFhRDHg3mhix2NoE3pHiaCzggleQZDZD",
	})

	if e != nil {
		fmt.Println("Erro ao postar no Facebook.", res)
	}
}
