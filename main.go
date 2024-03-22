package main

import (
	"convertSvg/src/converter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http"
)

var svgConverter *converter.Converter

func upload(c echo.Context) error {
	// Source
	r := c.Request()
	method := c.QueryParam("method")
	body := r.Body
	data, err := io.ReadAll(body)
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {

		}
	}(body)

	if err != nil {
		return c.NoContent(400)
	}

	var output []byte

	if method == "inkscape" || method == "1" {
		output, err = svgConverter.Convert(data)
		if err != nil {
			return c.NoContent(400)
		}
	} else {
		return c.NoContent(400)
	}
	return c.Blob(http.StatusOK, "image/png", output)
}

func main() {
	svgConverter = converter.New()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server OK!")
	})
	e.POST("/upload", upload)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.IPExtractor = echo.ExtractIPFromXFFHeader()
	e.Logger.Fatal(e.Start(":80"))
}
