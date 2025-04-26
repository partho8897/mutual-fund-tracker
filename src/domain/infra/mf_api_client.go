package infra

import (
	"github.com/mutual-fund-tracker/src/domain/model"
	"github.com/mutual-fund-tracker/src/mfterror"
)

type MFApiClient interface {
	ListAllFunds() ([]*model.FundDetails, *mfterror.MFTError)
	SearchFund(fundName string) ([]*model.FundDetails, *mfterror.MFTError)
	GetFundHistoricalData(fundId string) (*model.FundHistoricalData, *mfterror.MFTError)
	GetFundLatestData(fundId string) (*model.FundHistoricalData, *mfterror.MFTError)
}
