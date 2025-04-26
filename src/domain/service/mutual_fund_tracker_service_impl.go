package service

import (
	"github.com/mutual-fund-tracker/src/domain/infra"
	"github.com/mutual-fund-tracker/src/domain/model"
	"github.com/mutual-fund-tracker/src/domain/model/request"
	"github.com/mutual-fund-tracker/src/domain/model/response"
	"github.com/mutual-fund-tracker/src/mfterror"
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

func (service MutualFundTrackerServiceImpl) GetFundHistoricalData(fundId string) (*model.FundHistoricalData, *mfterror.MFTError) {
	return service.client.GetFundHistoricalData(fundId)
}

func (service MutualFundTrackerServiceImpl) GetFundLatestData(fundId string) (*model.FundHistoricalData, *mfterror.MFTError) {
	return service.client.GetFundLatestData(fundId)
}

func (service MutualFundTrackerServiceImpl) Backtrack(request request.BackTrackRequest) (*response.BackTrackResponse, *mfterror.MFTError) {
	//TODO implement me
	panic("implement me")
}
