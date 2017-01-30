package pages

import "github.com/bunsenapp/go-selenium"

type DockerHomePage struct {
	Page
}

func NewDockerHomePage(driver goselenium.WebDriver) DockerHomePage {

	var p = Page{
		Driver: driver,
	}

	return DockerHomePage{
		Page: p,
	}
}

func (d *DockerHomePage)GoTo() {
	d.Driver.Go("https://www.docker.com/")
}