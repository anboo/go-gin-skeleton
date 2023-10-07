package torgi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"torgi.gov.ru/model/lot"
)

type Parser struct {
	client *http.Client
}

func (p *Parser) SearchLots(ctx context.Context, search string) (lot.ListResponse, error) {
	req, err := http.NewRequest("GET", "https://torgi.gov.ru/new/api/public/lotcards/search?text="+url.QueryEscape(search)+"&byFirstVersion=true&withFacets=true&size=100&sort=firstVersionPublicationDate,desc", nil)
	if err != nil {
		return lot.ListResponse{}, fmt.Errorf("new request: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return lot.ListResponse{}, fmt.Errorf("request failed: %w", err)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return lot.ListResponse{}, fmt.Errorf("read response: %w", err)
	}

	var response lot.ListResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		return lot.ListResponse{}, fmt.Errorf("unmarshal response: %w", err)
	}

	return response, nil
}
