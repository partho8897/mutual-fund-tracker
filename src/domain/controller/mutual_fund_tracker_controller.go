package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mutual-fund-tracker/src/domain/model/request"
	"github.com/mutual-fund-tracker/src/domain/service"
	"strings"
)

type MutualFundTrackerController interface {
	GetTrackerRoutes(r *gin.Engine) *gin.Engine
}

type MutualFundTrackerControllerImpl struct {
	service service.MutualFundTrackerService
}

func NewMutualFundTrackerController(service service.MutualFundTrackerService) MutualFundTrackerController {
	return MutualFundTrackerControllerImpl{service: service}
}

func (controller MutualFundTrackerControllerImpl) GetTrackerRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/mftracker/v1/list", controller.ListAllFunds)
	r.GET("/mftracker/v1/search", controller.SearchFund)

	r.GET("/mftracker/v1/fund/:fundId", controller.GetFundHistoricalData)
	r.GET("/mftracker/v1/fund/:fundId/latest", controller.GetFundLatestData)

	r.POST("/mftracker/v1/backtrack", controller.Backtrack)

	return r
}

func (controller MutualFundTrackerControllerImpl) ListAllFunds(ctx *gin.Context) {
	funds, mftError := controller.service.ListAllFunds()
	if mftError != nil {
		ctx.JSON(mftError.GetHTTPErrorCode(), gin.H{"error": mftError.Error()})
		return
	}

	ctx.JSON(200, funds)
}

func (controller MutualFundTrackerControllerImpl) SearchFund(ctx *gin.Context) {
	fundName := ctx.Query("fundName")
	if "" == strings.TrimSpace(fundName) {
		ctx.JSON(400, gin.H{"error": "fundName is required"})
		return
	}

	funds, mftError := controller.service.SearchFund(fundName)
	if mftError != nil {
		ctx.JSON(mftError.GetHTTPErrorCode(), gin.H{"error": mftError.Error()})
		return
	}

	ctx.JSON(200, funds)
}

func (controller MutualFundTrackerControllerImpl) GetFundHistoricalData(ctx *gin.Context) {
	fundId := ctx.Param("fundId")
	if "" == strings.TrimSpace(fundId) {
		ctx.JSON(400, gin.H{"error": "fundId is required"})
		return
	}

	fund, mftError := controller.service.GetFundHistoricalData(fundId)
	if mftError != nil {
		ctx.JSON(mftError.GetHTTPErrorCode(), gin.H{"error": mftError.Error()})
		return
	}

	ctx.JSON(200, fund)
}

func (controller MutualFundTrackerControllerImpl) GetFundLatestData(ctx *gin.Context) {
	fundId := ctx.Param("fundId")
	if "" == strings.TrimSpace(fundId) {
		ctx.JSON(400, gin.H{"error": "fundId is required"})
		return
	}

	fund, mftError := controller.service.GetFundLatestData(fundId)
	if mftError != nil {
		ctx.JSON(mftError.GetHTTPErrorCode(), gin.H{"error": mftError.Error()})
		return
	}

	ctx.JSON(200, fund)
}

func (controller MutualFundTrackerControllerImpl) Backtrack(ctx *gin.Context) {
	var backtrackRequest request.BackTrackRequest
	err := ctx.BindJSON(&backtrackRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body."})
		return
	}

	response, mftError := controller.service.Backtrack(backtrackRequest)
	if mftError != nil {
		ctx.JSON(mftError.GetHTTPErrorCode(), gin.H{"error": mftError.Error()})
		return
	}

	ctx.JSON(200, response)
}
