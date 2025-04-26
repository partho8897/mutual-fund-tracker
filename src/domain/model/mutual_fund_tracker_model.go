package model

type FundDetails struct {
	SchemeCode int32  `json:"schemeCode"`
	SchemeName string `json:"schemeName"`
}

type FundMetaData struct {
	FundHouse      string `json:"fund_house"`
	SchemeType     string `json:"scheme_type"`
	SchemeCategory string `json:"scheme_category"`
	SchemeCode     int32  `json:"scheme_code"`
	SchemeName     string `json:"scheme_name"`
}

type FundData struct {
	Date string `json:"date"`
	NAV  string `json:"nav"`
}

type FundHistoricalData struct {
	FundMetaData FundMetaData `json:"meta"`
	FundData     []FundData   `json:"data"`
}
