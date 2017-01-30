package driver

import (
	"github.com/bunsenapp/go-selenium"
)

func ChromeDriver() goselenium.WebDriver {
	capabilities := goselenium.Capabilities{}
	capabilities.SetBrowser(goselenium.ChromeBrowser())
	return newDriver(capabilities)
}