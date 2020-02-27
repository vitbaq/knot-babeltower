package interactors

import (
	"errors"
	"fmt"
	"github.com/CESARBR/knot-babeltower/pkg/network"
)

// RequestData executes the use case operations to request data from the thing
func (i *ThingInteractor) PublishData(authorization, thingID string, data []network.Data) error {
	var sensorIds []int

	if authorization == "" {
		return errors.New("authorization key not provided")
	}

	thing, err := i.thingProxy.GetThing(authorization, thingID)
	if err != nil {
		i.logger.Error(err)
		return err
	}

	if thing.Schema == nil {
		i.logger.Error(fmt.Errorf("thing %s has no schema yet", thing.ID))
		return err
	}

	for _, dt := range data {
		sensorIds = append(sensorIds, dt.SensorID)
	}

	err = validateSensors(sensorIds, thing.Schema)
	if err != nil {
		i.logger.Error(err)
		return err
	}

	err = i.connectorPublisher.SendPublishData(thingID, data)
	if err != nil {
		i.logger.Error(err)
		return err
	}

	i.logger.Info("publish data message successfully sent")
	return nil
}
