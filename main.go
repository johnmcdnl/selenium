package main

import (
	"github.com/johnmcdnl/selenium/driver"
	"github.com/Sirupsen/logrus"
	"time"
	"github.com/johnmcdnl/selenium/pages"
	"github.com/bunsenapp/go-selenium"
	"io/ioutil"
	"os"
	"fmt"
	"sync"
)

func init() {
	for i := 0; i <= 5; i++ {
		logrus.Infof("Waiting for startup %d", i)
		time.Sleep(1 * time.Second)
	}
}

var wg sync.WaitGroup

func main() {

	drivers := make(map[string]goselenium.WebDriver)
	drivers["firefox"] = driver.FirefoxDriver()
	drivers["chrome"] = driver.ChromeDriver()

	for browser, webDriver := range drivers {
		wg.Add(1)
		go MyTest(browser, webDriver)
	}
	wg.Wait()
}

func MyTest(browser string, driver goselenium.WebDriver) {
	dockerHomePage := pages.NewDockerHomePage(driver)
	dockerHomePage.GoTo()
	TakeScreenShot(browser, dockerHomePage.Driver)
	url, err := dockerHomePage.Driver.CurrentURL()
	logrus.Info(url, err)
	getStarted, err := dockerHomePage.Driver.FindElement(goselenium.ByLinkText("Get Started"))
	getStarted.Click()
	TakeScreenShot(browser, dockerHomePage.Driver)
	url, err = dockerHomePage.Driver.CurrentURL()
	logrus.Info(url, err)
	wg.Done()
}

func TakeScreenShot(browser string, driver goselenium.WebDriver) {
	screenShot, err := driver.Screenshot()
	if err != nil {
		panic(err)
	}
	bytes, err := screenShot.ImageBytes()
	if err != nil {
		panic(err)
	}
	if err = ioutil.WriteFile(fmt.Sprintf("/screenshots/%s_%d_sceenshot.jpg", browser, time.Now().Nanosecond()), bytes, os.ModePerm); err != nil {
		panic(err)
	}
}



