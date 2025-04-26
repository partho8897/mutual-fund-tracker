package service

import (
	"github.com/mutual-fund-tracker/src/domain/model"
	"github.com/mutual-fund-tracker/src/domain/model/request"
	"github.com/mutual-fund-tracker/src/domain/model/response"
	"github.com/mutual-fund-tracker/src/mfterror"
)

type MutualFundTrackerService interface {
	ListAllFunds() ([]*model.FundDetails, *mfterror.MFTError)
	SearchFund(fundName string) ([]*model.FundDetails, *mfterror.MFTError)
	GetFundHistoricalData(fundId string) (*model.FundHistoricalData, *mfterror.MFTError)
	GetFundLatestData(fundId string) (*model.FundHistoricalData, *mfterror.MFTError)
	Backtrack(request request.BackTrackRequest) (*response.BackTrackResponse, *mfterror.MFTError)
}
