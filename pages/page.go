package pages

import "github.com/bunsenapp/go-selenium"

type Page struct {
	Driver goselenium.WebDriver
	GoTo   func()
}
