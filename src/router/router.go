package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mutual-fund-tracker/src/domain/controller"
)

func NewRouter(mutualFundTrackerRouter controller.MutualFundTrackerController) (*gin.Engine, error) {
	router := gin.Default()
	mutualFundTrackerRouter.GetTrackerRoutes(router)

	return router, nil
}
