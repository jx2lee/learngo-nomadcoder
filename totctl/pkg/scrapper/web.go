package scrapper

import (
	"os"
	"strings"

	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("pkg/scrapper/home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove("fileName")
	term := strings.ToLower(CleanString(c.FormValue("term")))
	Scrape(term)
	return c.Attachment("fileName", "fileName")
	return nil
}

func ScrapperEcho() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
