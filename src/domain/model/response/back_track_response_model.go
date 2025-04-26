package response

type BackTrackResponse struct {
	TotalInvestment int64                    `json:"totalInvestment"`
	TotalReturn     int64                    `json:"totalReturn"`
	CAGR            float64                  `json:"cagr"`
	XIRR            float64                  `json:"xirr"`
	InvestmentInfos []InvestmentInfoResponse `json:"investmentInfos"`
}

type InvestmentInfoResponse struct {
	SchemeCode string `json:"schemeCode"`
	Amount     int64  `json:"amount"`
	Returns    int64  `json:"returns"`
}
