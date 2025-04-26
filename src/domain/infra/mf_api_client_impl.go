package infra

import (
	"github.com/mutual-fund-tracker/src/domain/infra/utils"
	"github.com/mutual-fund-tracker/src/domain/model"
	"github.com/mutual-fund-tracker/src/mfterror"
	"net/http"
	"net/url"
)

var BaseUrl = "https://api.mfapi.in/mf"

type MFApiClientImpl struct {
	httpClient *http.Client
}

func NewMFApiClientImpl() MFApiClient {
	return &MFApiClientImpl{httpClient: http.DefaultClient}
}

func (mfApiClient *MFApiClientImpl) ListAllFunds() ([]*model.FundDetails, *mfterror.MFTError) {
	response, clientErr := mfApiClient.httpClient.Get(BaseUrl)
	if nil != clientErr {
		return nil, mfterror.ERR_CONNECTING_MF_API.WithDetails("Error listing all funds")
	}

	var funds []*model.FundDetails
	if decodingErr := utils.DecodeResponseBody(response, &funds); decodingErr != nil {
		return nil, decodingErr
	}

	return funds, nil
}

func (mfApiClient *MFApiClientImpl) SearchFund(fundName string) ([]*model.FundDetails, *mfterror.MFTError) {
	response, clientErr := mfApiClient.httpClient.Get(BaseUrl + "/search?q=" + url.QueryEscape(fundName))
	if nil != clientErr {
		return nil, mfterror.ERR_CONNECTING_MF_API.WithDetails("Error searching fund")
	}

	var funds []*model.FundDetails
	if decodingErr := utils.DecodeResponseBody(response, &funds); decodingErr != nil {
		return nil, decodingErr
	}

	return funds, nil
}

func (mfApiClient *MFApiClientImpl) GetFundHistoricalData(fundId string) (*model.FundHistoricalData, *mfterror.MFTError) {
	response, clientErr := mfApiClient.httpClient.Get(BaseUrl + "/" + fundId)
	if nil != clientErr {
		return nil, mfterror.ERR_CONNECTING_MF_API.WithDetails("Error getting fund historical data")
	}

	var fundHistoricalData *model.FundHistoricalData
	if decodingErr := utils.DecodeResponseBody(response, &fundHistoricalData); decodingErr != nil {
		return nil, decodingErr
	}

	return fundHistoricalData, nil
}

func (mfApiClient *MFApiClientImpl) GetFundLatestData(fundId string) (*model.FundHistoricalData, *mfterror.MFTError) {
	response, clientErr := mfApiClient.httpClient.Get(BaseUrl + "/" + fundId + "/latest")
	if nil != clientErr {
		return nil, mfterror.ERR_CONNECTING_MF_API.WithDetails("Error getting fund latest data")
	}

	var fundHistoricalData *model.FundHistoricalData
	if decodingErr := utils.DecodeResponseBody(response, &fundHistoricalData); decodingErr != nil {
		return nil, decodingErr
	}

	return fundHistoricalData, nil
}
