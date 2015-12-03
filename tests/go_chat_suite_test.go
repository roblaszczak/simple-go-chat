package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"

	"testing"
)

func TestGoChat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoChat Suite")
}

var agoutiDriver *agouti.WebDriver

var _ = BeforeSuite(func() {
	agoutiDriver = agouti.NewWebDriver(
		"http://username:access_key@ondemand.saucelabs.com/wd/hub",
		[]string{"echo", "do nothing"})

	Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
	Expect(agoutiDriver.Stop()).To(Succeed())
})
