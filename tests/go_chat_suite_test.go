package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"

	"testing"
	"fmt"
	"os"
)

func TestGoChat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoChat Suite")
}

var agoutiDriver *agouti.WebDriver

var _ = BeforeSuite(func() {
	driverUrl := fmt.Sprintf(
		"http://%s:%s@ondemand.saucelabs.com/wd/hub",
		os.Getenv("SAUCE_USERNAME"),
		os.Getenv("SAUCE_ACCESS_KEY"))

	capabilities := agouti.NewCapabilities()
	capabilities["tunnel-identifier"] = os.Getenv("TRAVIS_JOB_NUMBER")

	option := agouti.Desired(capabilities)

	agoutiDriver = agouti.NewWebDriver(driverUrl, []string{"echo", "do nothing"}, option)

	Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
	Expect(agoutiDriver.Stop()).To(Succeed())
})
