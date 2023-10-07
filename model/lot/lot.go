package lot

import (
	"time"
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
