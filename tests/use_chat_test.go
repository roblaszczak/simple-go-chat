package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/roblaszczak/simple-go-chat/cmd/gochat"
	"github.com/roblaszczak/simple-go-chat/config"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
	"time"
	"log"
)

var _ = Describe("UserConnect", func() {
	var page *agouti.Page

	BeforeEach(func() {
		go func() {
			RunServer(config.SERVER_HOST, config.SERVER_PORT)
		}()

		var err error
		page, err = agoutiDriver.NewPage(agouti.Browser("chrome"))
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("should communicate via chat", func() {
		By("allowing the user to connect chat", func() {
			Expect(page.Navigate("http://localhost:8080")).To(Succeed())

			html, err := page.HTML()
			if err != nil {
				panic(err)
			}
			print("html:", html)


			time.Sleep(time.Second*2)

			logs, err := page.ReadAllLogs("browser")
			if err != nil {
				log.Fatal(err)
			}
			print("log:", logs)

			firstPostContent := getLastPost(page).Find(".content")
			Expect(firstPostContent).To(MatchText("hello, anonymus_[0-9]{3}"))
		})

		By("allowing the user to send message", func() {
			Expect(page.Find("[ng-model='message']").Fill("hello, biczez")).To(Succeed())
			Expect(page.FindByButton("Send").Submit()).To(Succeed())
			Expect(getLastPost(page).Find(".content")).To(HaveText("hello, biczez"))
		})
	})
})

func getLastPost(page *agouti.Page) *agouti.Selection {
	posts := page.All("#chat .message")
	postsCount, err := posts.Count()
	if err != nil {
		panic(err)
	}

	return posts.At(postsCount - 1)
}
