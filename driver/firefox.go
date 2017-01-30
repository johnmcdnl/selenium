package driver

import (
	"github.com/bunsenapp/go-selenium"
)

func FirefoxDriver() goselenium.WebDriver {
	capabilities := goselenium.Capabilities{}
	capabilities.SetBrowser(goselenium.FirefoxBrowser())
	return newDriver(capabilities)
}