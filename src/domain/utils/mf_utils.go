package utils

import (
	"github.com/mutual-fund-tracker/src/domain/model"
	"github.com/mutual-fund-tracker/src/domain/model/request"
	"github.com/mutual-fund-tracker/src/domain/model/request/enum"
	"github.com/mutual-fund-tracker/src/domain/model/response"
	"math"
	"time"
)

func GetFundDataAtInvestmentDateWithDateString(data []model.FundData, from string) model.FundData {
	return GetFundDataAtInvestmentDate(data, GetDateFromString(from))
}

func GetFundDataAtInvestmentDate(data []model.FundData, fromDate time.Time) model.FundData {
	low, high := 0, len(data)-1
	var result model.FundData

	for low <= high {
		mid := (low + high) / 2

		fundDate := GetDateFromString(data[mid].Date)

		if fundDate == fromDate {
			result = data[mid]
			break
		} else if fundDate.Before(fromDate) { // Adjusted for descending order
			high = mid - 1
		} else {
			result = data[mid] // Track the closest match (immediately more than `from`)
			low = mid + 1
		}
	}

	// Return the closest match if found
	return result
}

func GetDateFromString(date string) time.Time {
	layout := "02-01-2006" // Layout for dd-mm-yyyy

	parsedDate, _ := time.Parse(layout, date)

	return parsedDate
}

func CalculateCagr(investmentAmount int64, returnAmount float64, from string) float64 {
	fromDate := GetDateFromString(from)
	currentDate := time.Now()
	durationYears := currentDate.Sub(fromDate).Hours() / (24 * 365)

	// Calculate CAGR using the formula: CAGR = ((Final Value / Initial Value)^(1/Years)) - 1
	cagr := returnAmount / float64(investmentAmount)
	cagr = math.Pow(cagr, 1/durationYears) - 1

	return cagr * 100 // Return CAGR as a percentage
}

func CalculateInvestmentFrequencyInMonths(frequency enum.InvestmentFrequency) int {
	var stepMonths int

	switch frequency {
	case enum.InvestmentFrequencyMonthly:
		stepMonths = 1
	case enum.InvestmentFrequencyQuarterly:
		stepMonths = 3
	case enum.InvestmentFrequencyYearly:
		stepMonths = 12
	default:
		return 0
	}

	return stepMonths
}

func AggregateInvestmentResults(investmentInfos []*response.InvestmentInfoResponse) (int64, float64) {
	var totalInvestmentAmount int64
	var totalReturnAmount float64

	for _, investmentInfo := range investmentInfos {
		totalInvestmentAmount += investmentInfo.Amount
		totalReturnAmount += float64(investmentInfo.Returns)
	}

	return totalInvestmentAmount, totalReturnAmount
}

func CreateBackTrackResponse(
	totalInvestmentAmount int64, totalReturnAmount float64, from string,
	investmentInfos []*response.InvestmentInfoResponse,
) *response.BackTrackResponse {
	return &response.BackTrackResponse{
		TotalInvestment: totalInvestmentAmount,
		TotalReturn:     float32(totalReturnAmount),
		CAGR:            float32(CalculateCagr(totalInvestmentAmount, totalReturnAmount, from)),
		XIRR:            0.0, // XIRR calculation can be added later
		InvestmentInfos: investmentInfos,
	}
}

func GetInvestmentInfoResponse(
	fundInfo request.FundInfoRequest, fundInvestment int64, fundReturn float64,
) *response.InvestmentInfoResponse {
	return &response.InvestmentInfoResponse{
		SchemeCode: fundInfo.SchemeCode,
		Amount:     fundInvestment,
		Returns:    float32(fundReturn),
	}
}
