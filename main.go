package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-spatial/proj"
	"github.com/playwright-community/playwright-go"
)

type ItemCharacteristic struct {
	CharacteristicValue any    `json:"characteristicValue,omitempty"`
	Name                string `json:"name"`
	Code                string `json:"code"`
	Unit                struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"unit,omitempty"`
	Type string `json:"type"`
}

type ItemAttribute struct {
	Code          string `json:"code"`
	FullName      string `json:"fullName"`
	Value         any    `json:"value,omitempty"`
	AttributeType string `json:"attributeType"`
	Group         struct {
		Code             string `json:"code"`
		Name             string `json:"name"`
		DisplayGroupType string `json:"displayGroupType"`
	} `json:"group"`
	SortOrder int `json:"sortOrder"`
}

type NoticeAttribute struct {
	Code          string `json:"code"`
	FullName      string `json:"fullName"`
	Value         bool   `json:"value"`
	AttributeType string `json:"attributeType"`
	Group         struct {
		Code             string `json:"code"`
		Name             string `json:"name"`
		DisplayGroupType string `json:"displayGroupType"`
	} `json:"group"`
	SortOrder int `json:"sortOrder"`
}

type Item struct {
	ID           string `json:"id"`
	NoticeNumber string `json:"noticeNumber"`
	LotNumber    int    `json:"lotNumber"`
	LotStatus    string `json:"lotStatus"`
	BiddType     struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"biddType"`
	BiddForm struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"biddForm"`
	LotName         string               `json:"lotName"`
	LotDescription  string               `json:"lotDescription"`
	PriceMin        float64              `json:"priceMin"`
	BiddEndTime     time.Time            `json:"biddEndTime"`
	LotImages       []string             `json:"lotImages"`
	Characteristics []ItemCharacteristic `json:"characteristics"`
	CurrencyCode    string               `json:"currencyCode"`
	SubjectRFCode   string               `json:"subjectRFCode"`
	Category        struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"category"`
	CreateDate       time.Time         `json:"createDate"`
	TimeZoneName     string            `json:"timeZoneName"`
	TimezoneOffset   string            `json:"timezoneOffset"`
	HasAppeals       bool              `json:"hasAppeals"`
	IsStopped        bool              `json:"isStopped"`
	Attributes       []ItemAttribute   `json:"attributes"`
	NoticeAttributes []NoticeAttribute `json:"noticeAttributes"`
	IsAnnulled       bool              `json:"isAnnulled"`
}

type ListResponse struct {
	Content  []Item `json:"content"`
	Pageable struct {
		Sort struct {
			Unsorted bool `json:"unsorted"`
			Sorted   bool `json:"sorted"`
			Empty    bool `json:"empty"`
		} `json:"sort"`
		PageNumber int  `json:"pageNumber"`
		PageSize   int  `json:"pageSize"`
		Offset     int  `json:"offset"`
		Unpaged    bool `json:"unpaged"`
		Paged      bool `json:"paged"`
	} `json:"pageable"`
	CategoryFacet []struct {
		ID    string `json:"_id"`
		Count int    `json:"count"`
	} `json:"categoryFacet"`
	TotalPages       int  `json:"totalPages"`
	TotalElements    int  `json:"totalElements"`
	Last             bool `json:"last"`
	NumberOfElements int  `json:"numberOfElements"`
	First            bool `json:"first"`
	Size             int  `json:"size"`
	Number           int  `json:"number"`
	Sort             struct {
		Unsorted bool `json:"unsorted"`
		Sorted   bool `json:"sorted"`
		Empty    bool `json:"empty"`
	} `json:"sort"`
	Empty bool `json:"empty"`
}

func main() {
	var lonlat = []float64{3952358.010151192, 7536547.163115314}

	xy, err := proj.Inverse(proj.EPSG3857, lonlat)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v %v", xy[0], xy[1])

	ctx := context.Background()

	req, err := http.NewRequest("GET", "https://torgi.gov.ru/new/api/public/lotcards/search?text=%D0%A8%D0%B0%D1%85%D0%BE%D0%B2%D1%81%D0%BA%D0%B0%D1%8F&byFirstVersion=true&withFacets=true&size=100&sort=firstVersionPublicationDate,desc", nil)
	if err != nil {
		panic(err)
	}

	http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		panic("request failed" + err.Error())
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic("try read response " + err.Error())
	}

	var response ListResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		panic(err)
	}

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

	fmt.Printf("%v", response.Content[0].Characteristics[4])
}
