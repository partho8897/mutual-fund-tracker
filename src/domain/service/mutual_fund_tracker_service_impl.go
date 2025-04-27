package service

import (
	"fmt"
	"github.com/mutual-fund-tracker/src/domain/infra"
	"github.com/mutual-fund-tracker/src/domain/model"
	"github.com/mutual-fund-tracker/src/domain/model/request"
	"github.com/mutual-fund-tracker/src/domain/model/request/enum"
	"github.com/mutual-fund-tracker/src/domain/model/response"
	"github.com/mutual-fund-tracker/src/domain/utils"
	"github.com/mutual-fund-tracker/src/mfterror"
	"strconv"
	"sync"
	"time"
)

type MutualFundTrackerServiceImpl struct {
	client infra.MFApiClient
}

func NewMutualFundTrackerServiceImpl(client infra.MFApiClient) MutualFundTrackerService {
	return &MutualFundTrackerServiceImpl{
		client: client,
	}
}

func (service MutualFundTrackerServiceImpl) ListAllFunds() ([]*model.FundDetails, *mfterror.MFTError) {
	return service.client.ListAllFunds()
}

func (service MutualFundTrackerServiceImpl) SearchFund(fundName string) ([]*model.FundDetails, *mfterror.MFTError) {
	return service.client.SearchFund(fundName)
}

func (service MutualFundTrackerServiceImpl) GetFundHistoricalData(fundId string) (
	*model.FundHistoricalData, *mfterror.MFTError,
) {
	historicalData, fundHistoryErr := service.client.GetFundHistoricalData(fundId)
	if fundHistoryErr != nil {
		return nil, fundHistoryErr
	}

	if historicalData == nil || historicalData.FundMetaData.SchemeCode == 0 {
		return nil, mfterror.ERR_BAD_REQUEST.WithDetails(fmt.Sprintf("Invalid fund id: %s.", fundId))
	}

	return historicalData, nil
}

func (service MutualFundTrackerServiceImpl) GetFundLatestData(fundId string) (
	*model.FundHistoricalData, *mfterror.MFTError,
) {
	return service.client.GetFundLatestData(fundId)
}

func (service MutualFundTrackerServiceImpl) Backtrack(request request.BackTrackRequest) (
	*response.BackTrackResponse, *mfterror.MFTError,
) {
	fundHistoricalDatasMap, err := service.getFundHistoricalDatas(request)
	if err != nil {
		return nil, err
	}

	switch request.InvestmentType {
	case enum.InvestmentTypeLumpsum:
		return service.backtrackLumpSumInvestment(fundHistoricalDatasMap, request)
	case enum.InvestmentTypeSIP:
		return service.backtrackSIPInvestment(fundHistoricalDatasMap, request)
	default:
		return nil, mfterror.ERR_BAD_REQUEST.WithDetails("Invalid investment type")
	}
}

func (service MutualFundTrackerServiceImpl) getFundHistoricalDatas(request request.BackTrackRequest) (
	map[string]*model.FundHistoricalData, *mfterror.MFTError,
) {
	fundHistoricalDatas := make(map[string]*model.FundHistoricalData)
	fundHistoricalDataChan := make(chan *model.FundHistoricalData, len(request.FundInfos))
	errorChan := make(chan *mfterror.MFTError, len(request.FundInfos))

	for _, fundInfo := range request.FundInfos {
		go func(schemeCode string) {
			fundHistoricalData, err := service.GetFundHistoricalData(schemeCode)
			if err != nil {
				errorChan <- err
				return
			}
			fundHistoricalDataChan <- fundHistoricalData
		}(fundInfo.SchemeCode)
	}

	for range request.FundInfos {
		select {
		case err := <-errorChan:
			return nil, err
		case fundHistoricalData := <-fundHistoricalDataChan:
			fundHistoricalDatas[fmt.Sprintf("%d", fundHistoricalData.FundMetaData.SchemeCode)] = fundHistoricalData
		}
	}

	return fundHistoricalDatas, nil
}

func (service MutualFundTrackerServiceImpl) backtrackLumpSumInvestment(
	fundHistoricalDatasMap map[string]*model.FundHistoricalData, backtrackRequest request.BackTrackRequest,
) (*response.BackTrackResponse, *mfterror.MFTError) {
	investmentInfos := service.calculateLumpSumInvestments(fundHistoricalDatasMap, backtrackRequest)
	totalInvestmentAmount, totalReturnAmount := utils.AggregateInvestmentResults(investmentInfos)

	return utils.CreateBackTrackResponse(
		totalInvestmentAmount, totalReturnAmount, backtrackRequest.From, investmentInfos,
	), nil
}

func (service MutualFundTrackerServiceImpl) backtrackSIPInvestment(
	fundHistoricalDatasMap map[string]*model.FundHistoricalData, backtrackRequest request.BackTrackRequest,
) (*response.BackTrackResponse, *mfterror.MFTError) {
	investmentInfos, totalInvestmentAmount, totalReturnAmount := service.calculateSIPInvestments(
		fundHistoricalDatasMap, backtrackRequest,
	)

	return utils.CreateBackTrackResponse(
		totalInvestmentAmount, totalReturnAmount, backtrackRequest.From, investmentInfos,
	), nil
}

func (service MutualFundTrackerServiceImpl) calculateLumpSumInvestments(
	fundHistoricalDatasMap map[string]*model.FundHistoricalData, backtrackRequest request.BackTrackRequest,
) []*response.InvestmentInfoResponse {
	var mu sync.Mutex
	wg := sync.WaitGroup{}
	investmentInfos := make([]*response.InvestmentInfoResponse, 0)

	for _, fundInfo := range backtrackRequest.FundInfos {
		wg.Add(1)
		go func(fundInfo request.FundInfoRequest) {
			defer wg.Done()

			investmentInfo := service.calculateLumpSumForFund(fundHistoricalDatasMap, fundInfo, backtrackRequest.From)

			mu.Lock()
			investmentInfos = append(investmentInfos, investmentInfo)
			mu.Unlock()
		}(fundInfo)
	}

	wg.Wait()
	return investmentInfos
}

func (service MutualFundTrackerServiceImpl) calculateLumpSumForFund(
	fundHistoricalDatasMap map[string]*model.FundHistoricalData, fundInfo request.FundInfoRequest, from string,
) *response.InvestmentInfoResponse {
	dataAtInvestmentDate := utils.GetFundDataAtInvestmentDateWithDateString(
		fundHistoricalDatasMap[fundInfo.SchemeCode].FundData, from,
	)
	currentData := fundHistoricalDatasMap[fundInfo.SchemeCode].FundData[0]

	investmentDateNav, _ := strconv.ParseFloat(dataAtInvestmentDate.NAV, 64)
	navAtInvestment := float64(fundInfo.Amount) / investmentDateNav
	currentNAV, _ := strconv.ParseFloat(currentData.NAV, 64)

	returnAmount := currentNAV * navAtInvestment

	return utils.GetInvestmentInfoResponse(fundInfo, fundInfo.Amount, returnAmount, from)
}

func (service MutualFundTrackerServiceImpl) calculateSIPInvestments(
	fundHistoricalDatasMap map[string]*model.FundHistoricalData, backtrackRequest request.BackTrackRequest,
) ([]*response.InvestmentInfoResponse, int64, float64) {
	var mu sync.Mutex
	wg := sync.WaitGroup{}
	investmentInfos := make([]*response.InvestmentInfoResponse, 0)
	var totalInvestmentAmount int64
	var totalReturnAmount float64

	for _, fundInfo := range backtrackRequest.FundInfos {
		wg.Add(1)
		go func(fundInfo request.FundInfoRequest) {
			defer wg.Done()

			fundInvestment, fundReturn := service.calculateSIPForFund(
				fundHistoricalDatasMap, fundInfo, backtrackRequest,
			)

			mu.Lock()
			totalInvestmentAmount += fundInvestment
			totalReturnAmount += fundReturn
			investmentInfos = append(
				investmentInfos,
				utils.GetInvestmentInfoResponse(fundInfo, fundInvestment, fundReturn, backtrackRequest.From),
			)
			mu.Unlock()
		}(fundInfo)
	}

	wg.Wait()

	return investmentInfos, totalInvestmentAmount, totalReturnAmount
}

func (service MutualFundTrackerServiceImpl) calculateSIPForFund(
	fundHistoricalDatasMap map[string]*model.FundHistoricalData, fundInfo request.FundInfoRequest,
	backtrackRequest request.BackTrackRequest,
) (int64, float64) {
	sipAmount := fundInfo.Amount
	fromDate := utils.GetDateFromString(backtrackRequest.From)
	toDate := time.Now()
	stepMonths := utils.CalculateInvestmentFrequencyInMonths(backtrackRequest.InvestmentFrequency)

	var fundTotalInvestment int64
	var fundTotalReturn float64

	for date := fromDate; !date.After(toDate); date = date.AddDate(0, stepMonths, 0) {
		dataAtInvestmentDate := utils.GetFundDataAtInvestmentDate(
			fundHistoricalDatasMap[fundInfo.SchemeCode].FundData, date,
		)
		currentData := fundHistoricalDatasMap[fundInfo.SchemeCode].FundData[0]

		investmentDateNav, _ := strconv.ParseFloat(dataAtInvestmentDate.NAV, 32)
		navAtInvestment := float64(sipAmount) / investmentDateNav
		currentNAV, _ := strconv.ParseFloat(currentData.NAV, 32)

		returnAmount := currentNAV * navAtInvestment
		fundTotalInvestment += sipAmount
		fundTotalReturn += returnAmount
	}

	return fundTotalInvestment, fundTotalReturn
}
