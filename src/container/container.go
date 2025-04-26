package container

import (
	"github.com/mutual-fund-tracker/src/domain/controller"
	"github.com/mutual-fund-tracker/src/domain/infra"
	"github.com/mutual-fund-tracker/src/domain/service"
	"github.com/mutual-fund-tracker/src/router"
	"go.uber.org/dig"
)

func BuildContainer() (*dig.Container, error) {
	container := dig.New()

	// Register infra
	err := getError(container.Provide(infra.NewMFApiClientImpl), nil)

	// Register service
	err = getError(container.Provide(service.NewMutualFundTrackerServiceImpl), err)

	// Register controller
	err = getError(container.Provide(controller.NewMutualFundTrackerController), err)

	// Register router
	err = getError(container.Provide(router.NewRouter), err)

	return container, nil
}

func getError(previousError error, newErr error) error {
	if previousError != nil {
		return previousError
	} else {
		return newErr
	}
}
