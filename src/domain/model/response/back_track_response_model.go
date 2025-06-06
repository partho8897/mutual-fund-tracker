package response

import "time"

type BackTrackResponse struct {
	TotalInvestment int64                     `json:"totalInvestment"`
	TotalReturn     float32                   `json:"totalReturn"`
	CAGR            float32                   `json:"cagr"`
	XIRR            float32                   `json:"xirr"`
	InvestmentInfos []*InvestmentInfoResponse `json:"investmentInfos"`
}

type InvestmentInfoResponse struct {
	SchemeCode      string      `json:"schemeCode"`
	Amount          int64       `json:"amount"`
	Returns         float32     `json:"returns"`
	Cagr            float32     `json:"cagr"`
	Xirr            float32     `json:"xirr"`
	InvestmentDates []time.Time `json:"-"`
}
