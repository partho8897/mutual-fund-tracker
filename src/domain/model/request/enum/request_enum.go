package enum

type InvestmentType string

type InvestmentFrequency string

const (
	InvestmentTypeSIP     InvestmentType = "SIP"
	InvestmentTypeLumpsum InvestmentType = "Lumpsum"
)

const (
	InvestmentFrequencyMonthly   InvestmentFrequency = "Monthly"
	InvestmentFrequencyQuarterly InvestmentFrequency = "Quarterly"
	InvestmentFrequencyYearly    InvestmentFrequency = "Yearly"
)
