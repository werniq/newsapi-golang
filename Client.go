package newsapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/werniq/newsapi-golang/logger"
	"github.com/werniq/newsapi-golang/models"
	"log"
	"net/http"
	"time"
)

type Client struct {
	ApiKey string
}

var (
	newsApiUri        = "https://newsapi.org/v2/everything"
	possibleLanguages = []string{"ar", "de", "en", "es", "fr", "he", "it", "nl", "no", "pt", "ru", "sv", "ud", "zh"}
	possibleCountries = []string{"us", "ca", "gb", "de", "fr", "it", "es", "ar", "cl", "co", "mx"}
)

// NewClient creates a news api client
func NewClient(apiKey string) *Client {
	request, err := http.NewRequest("GET", newsApiUri, nil)
	if err != nil {
		logger.NewLogger().Printf("error creating new request")
		return nil
	}

	request.Header.Add("X-Api-Key", apiKey)

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		logger.NewLogger().Printf("error sending request")
		return nil
	}

	if res.StatusCode == http.StatusOK {
		return &Client{
			ApiKey: apiKey,
		}
	} else {
		log.Fatalf("error creating client: %s", res.Status)
		return nil
	}
}

// getRequestWithApiKey sends a post request to the news api with the api key
func (c *Client) getRequestWithApiKey(uri string, body []byte) ([]*models.NewsApiResponse, error) {
	request, err := http.NewRequest("GET", uri, bytes.NewBuffer(body))
	if err != nil {
		logger.NewLogger().Printf("error creating new request")
		return nil, err
	}

	request.Header.Add("X-Api-Key", c.ApiKey)

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		logger.NewLogger().Printf("error sending request")
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		var newsApiResponse []*models.NewsApiResponse
		err = json.NewDecoder(res.Body).Decode(&newsApiResponse)
		if err != nil {
			logger.NewLogger().Printf("error decoding response body")
			return nil, err
		}
		return newsApiResponse, nil
	} else {
		log.Fatalf("error creating client: %s", res.Status)
		return nil, nil
	}
}

// GetNews function creates a new request and sends it to the news api
func (c *Client) GetNews() ([]*models.NewsApiResponse, error) {
	req, err := http.NewRequest("GET", newsApiUri, nil)
	if err != nil {
		logger.NewLogger().Printf("error creating new request")
		return nil, err
	}

	req.Header.Add("X-Api-Key", c.ApiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.NewLogger().Printf("error sending request")
		return nil, err
	}

	var newsApiResponse []*models.NewsApiResponse

	err = json.NewDecoder(res.Body).Decode(&newsApiResponse)
	if err != nil {
		logger.NewLogger().Printf("error decoding response body")
		return nil, err
	}

	return newsApiResponse, nil
}

// GetTopHeadlines function gets the top headlines from the news api
func (c *Client) GetTopHeadlines() ([]*models.NewsApiResponse, error) {
	newsApiResponse, err := c.getRequestWithApiKey("https://newsapi.org/v2/top-headline", nil)
	if err != nil {
		logger.NewLogger().Printf("error getting top headlines")
		return nil, err
	}
	return newsApiResponse, nil
}

// GetEverything function gets everything from the news api
func (c *Client) GetEverything() ([]*models.NewsApiResponse, error) {
	newsApiResponse, err := c.getRequestWithApiKey("https://newsapi.org/v2/everything", nil)
	if err != nil {
		logger.NewLogger().Printf("error getting everything")
		return nil, err
	}
	return newsApiResponse, nil
}

// GetSources function gets the sources from the news api
func (c *Client) GetSources() ([]*models.NewsApiResponse, error) {
	newsApiResponse, err := c.getRequestWithApiKey("https://newsapi.org/v2/sources", nil)
	if err != nil {
		logger.NewLogger().Printf("error getting sources")
		return nil, err
	}
	return newsApiResponse, nil
}

// GetLatestNewsBySource function gets the latest news by source from the news api
func (c *Client) GetLatestNewsBySource(source string) ([]*models.NewsApiResponse, error) {
	today := time.Now()
	lastHour := today.Add(-time.Hour)
	lastHourStr := lastHour.Format("2006-01-02")
	uri := fmt.Sprintf("https://newsapi.org/v2/everything?sources=%s?from=%s", source, lastHourStr)

	newsApiResponse, err := c.getRequestWithApiKey(uri, nil)
	if err != nil {
		logger.NewLogger().Printf("error getting latest news by source")
		return nil, err
	}

	return newsApiResponse, nil
}

// SearchNewsByQuery function searches for news from the news api by provided query
func (c *Client) SearchNewsByQuery(query string) ([]*models.NewsApiResponse, error) {
	newsApiResponse, err := c.getRequestWithApiKey(newsApiUri+"?q="+query, nil)
	if err != nil {
		logger.NewLogger().Printf("error searching news by query")
		return nil, err
	}

	return newsApiResponse, nil
}

// SearchNewsByLanguage function searches for news from the news api by provided language
func (c *Client) SearchNewsByLanguage(language string) ([]*models.NewsApiResponse, error) {
	ok := false
	for _, v := range possibleLanguages {
		if language == v {
			ok = true
		}
	}

	if !ok {
		return nil, errors.New("invalid language")
	}

	newsApiResponse, err := c.getRequestWithApiKey(newsApiUri+"?language="+language, nil)
	if err != nil {
		return nil, err
	}

	return newsApiResponse, nil
}

// SearchNewsByCountry function searches for news from the news api by provided country
func (c *Client) SearchNewsByCountry(country string) ([]*models.NewsApiResponse, error) {
	ok := false
	for _, v := range possibleCountries {
		if country == v {
			ok = true
		}
	}

	if !ok {
		return nil, errors.New("invalid country")
	}

	newsApiResponse, err := c.getRequestWithApiKey(newsApiUri+"?country="+country, nil)
	if err != nil {
		return nil, err
	}

	return newsApiResponse, nil
}

// GetNewsByCategory is shorthand for SearchNewsByQuery
func (c *Client) GetNewsByCategory(category string) ([]*models.NewsApiResponse, error) {
	return c.SearchNewsByQuery(category)
}

// GetNewsByLanguage is shorthand for SearchNewsByLanguage
func (c *Client) GetNewsByLanguage(language string) ([]*models.NewsApiResponse, error) {
	return c.SearchNewsByLanguage(language)
}

// GetNewsByCountry is shorthand for SearchNewsByCountry
func (c *Client) GetNewsByCountry(country string) ([]*models.NewsApiResponse, error) {
	return c.SearchNewsByCountry(country)
}

// SetPageSize sets the page size for the news api
func (c *Client) SetPageSize(pageSize int, query string) string {
	return fmt.Sprintf(query+"?pageSize=%d", pageSize)
}

// SetLanguage sets the language for the query to news api
func (c *Client) SetLanguage(query, language string) string {
	return query + fmt.Sprintf("?language=%s", language)
}

// SetSources sets the sources for the query to news api
func (c *Client) SetSources(query, source string) string {
	return fmt.Sprintf(query+"?source=%s", source)
}

// SetDateRange sets the date range for the query to news api
func (c *Client) SetDateRange(query, from, to string) string {
	return fmt.Sprintf(query+"?from=%s&to=%s", from, to)
}

// SetSorting sets the sorting for the query to news api
func (c *Client) SetSorting(query, sortingType string) string {
	return fmt.Sprintf(query+"?sortBy=%s", sortingType)
}

// SetFromDate sets the from date for the query to news api
func (c *Client) SetFromDate(query, from string) string {
	return fmt.Sprintf(query+"?from=%s", from)
}

// SetToDate sets the to date for the query to news api
func (c *Client) SetToDate(query, to string) string {
	return fmt.Sprintf(query+"?to=%s", to)
}

// SetDomain sets the domain for the query to news api
func (c *Client) SetDomain(query, domain string) string {
	return fmt.Sprintf(query+"?domains=%s", domain)
}
