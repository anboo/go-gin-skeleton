package parser

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

func browser() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{})
	var ignoreHttpErrors = true
	browser.NewContext(playwright.BrowserNewContextOptions{
		IgnoreHttpsErrors: &ignoreHttpErrors,
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("http://pkk.rosreestr.ru/#/search/65.64951699999888,122.73014399999792/4/@bz87g7bj0", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
	}); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	page.Fill("div.search-input-container > div.container-type-ahead.without-border > input", "50:06:0090601:936")
	page.Click("div[title='Найти']")
	page.Click("div[title='Отдалить']")

	page.WaitForTimeout(1000)
	page.Click(".leaflet-control-zoom-out")
	page.Click(".leaflet-control-zoom-out")
	page.Click(".leaflet-control-zoom-out")

	page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("foo.png"),
	})
}
