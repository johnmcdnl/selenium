package driver

import (
	"github.com/bunsenapp/go-selenium"
	"fmt"
)

var hubEndPoint = "http://selenium_hub.selenium_selenium_grid_internal:4444/wd/hub"

func newDriver(capabilities goselenium.Capabilities) goselenium.WebDriver {
	driver, err := goselenium.NewSeleniumWebDriver(hubEndPoint, capabilities)
	if err != nil {
		panic(fmt.Sprintf("Failed to open session: %s\n", err))
	}

	_, err = driver.CreateSession()
	if err != nil {
		panic(err)
	}
	return driver
}