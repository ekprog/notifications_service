package bootstrap

import (
	"go.uber.org/dig"
	"microservice_clean_design/app"
	"microservice_clean_design/delivery"
	"microservice_clean_design/domain"
	"microservice_clean_design/interactors"
	"microservice_clean_design/repos"
)

func initDependencies(di *dig.Container) error {

	di.Provide(repos.NewNotificationsDBRepo, dig.As(new(domain.NotificationsRepository)))
	di.Provide(interactors.NewNotificationsUCase, dig.As(new(domain.NotificationInteractor)))
	di.Provide(delivery.NewNotificationsDeliveryService)

	// DELIVERY
	deliveryInit := func(d *delivery.NotificationsDeliveryService) error {
		if err := app.InitDelivery(d); err == nil {
			return err
		}
		return nil
	}

	err := di.Invoke(deliveryInit)
	if err != nil {
		return err
	}

	return nil
}
