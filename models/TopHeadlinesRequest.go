package models

var (
	RelevancySort           = "relevancy"
	PopularitySort          = "popularity"
	PublishedAtSort         = "publishedAt"
	USACountry              = "us"
	GBRCountry              = "gb"
	DECountry               = "de"
	FRCountry               = "fr"
	ITCountry               = "it"
	ESCountry               = "es"
	ARCountry               = "ar"
	CLCountry               = "cl"
	COCountry               = "co"
	MXCountry               = "mx"
	GBRLanguage             = "en"
	DELanguage              = "de"
	FRLanguage              = "fr"
	ITLanguage              = "it"
	MXLanguage              = "es"
	CNNSource               = "cnn"
	ABCNewsSource           = "abc-news"
	TheNewYorkTimesSource   = "the-new-york-times"
	TheWashingtonPostSource = "the-washington-post"
	ReutersSource           = "reuters"
	TheGuardianSource       = "the-guardian-uk"
	NBCNewsSource           = "nbc-news"
)

// TopHeadlinesRequest is the request for the top headlines
type TopHeadlinesRequest struct {
	Query    string
	Sources  string
	Language string
	Country  string
	Page     int
	PageSize int
	SortBy   string
}

// ChangeSortOptions changes the sort options of the request
func (t *TopHeadlinesRequest) ChangeSortOptions(s string) (*TopHeadlinesRequest, bool) {
	if s != RelevancySort && s != PopularitySort && s != PublishedAtSort {
		return t, false
	}
	t.SortBy = s
	return t, true
}

// ChangeCountryOptions changes the country of the request
func (t *TopHeadlinesRequest) ChangeCountryOptions(c string) (*TopHeadlinesRequest, bool) {
	if c != USACountry && c != GBRCountry && c != DECountry && c != FRCountry && c != ITCountry && c != ESCountry && c != ARCountry && c != CLCountry && c != COCountry && c != MXCountry {
		return t, false
	}
	t.Country = c
	return t, true
}

// ChangeLanguage changes the language of the request
func (t *TopHeadlinesRequest) ChangeLanguage(c string) (*TopHeadlinesRequest, bool) {
	if c != GBRLanguage && c != DELanguage && c != FRLanguage && c != ITLanguage && c != MXLanguage {
		return t, false
	}
	t.Language = c
	return t, true
}

// ChangeSource changes the source of the request
func (t *TopHeadlinesRequest) ChangeSource(s string) (*TopHeadlinesRequest, bool) {
	if s != CNNSource && s != ABCNewsSource && s != TheNewYorkTimesSource && s != TheWashingtonPostSource && s != ReutersSource && s != TheGuardianSource && s != NBCNewsSource {
		return t, false
	}
	t.Sources = s
	return t, true
}

// ChangeQuery changes the query of the request
func (t *TopHeadlinesRequest) ChangeQuery(q string) *TopHeadlinesRequest {
	t.Query = q
	return t
}

// ChangePage changes the page of the request
func (t *TopHeadlinesRequest) ChangePage(p int) *TopHeadlinesRequest {
	t.Page = p
	return t
}

// ChangePageSize changes the page size of the request
func (t *TopHeadlinesRequest) ChangePageSize(ps int) *TopHeadlinesRequest {
	t.PageSize = ps
	return t
}
