// package main

// import (
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/labstack/echo"
// 	"github.com/tylerb/graceful"
// )

// type Groups struct {
// 	group []Group `json:"groups"`
// }

// type Team struct {
// 	Name  string `json:"name"`
// 	Image string `json:"image"`
// }

// type Group struct {
// 	team []Team
// }

// type MyJsonName struct {
// 	Groups [][]struct {
// 		Image string `json:"image"`
// 		Name  string `json:"name"`
// 	} `json:"groups"`
// }

// func main() {

// 	var port = flag.String("p", "8001", "server port number")

// 	flag.Parse()

// 	log.Printf("Starting Rest API service on port %s\n", *port)

// 	e := echo.New()
// 	e.Server.Addr = ":" + *port

// 	r := e.Group("/api/v1")
// 	r.POST("/something", generateTemplate)

// 	graceful.ListenAndServe(e.Server, 5*time.Second)
// }

// func generateTemplate(c echo.Context) (err error) {

// 	groups := new(MyJsonName)

// 	if err := c.Bind(groups); err != nil {
// 		return c.JSON(http.StatusForbidden, err)
// 	}
// 	fmt.Println("JSON", groups)

// 	return c.JSON(http.StatusCreated, groups)
// }
