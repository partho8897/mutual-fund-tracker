package request

import "github.com/mutual-fund-tracker/src/domain/model/request/enum"

type BackTrackRequest struct {
	FundInfos           []FundInfoRequest        `json:"fundInfos"`
	From                string                   `json:"from"`
	InvestmentType      enum.InvestmentType      `json:"investmentType"`
	InvestmentFrequency enum.InvestmentFrequency `json:"investmentFrequency"`
}

type FundInfoRequest struct {
	SchemeCode string `json:"schemeCode"`
	Amount     int64  `json:"amount"`
}
