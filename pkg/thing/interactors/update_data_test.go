package interactors

import (
	"errors"
	"testing"

	"github.com/CESARBR/knot-babeltower/pkg/thing/entities"
	"github.com/CESARBR/knot-babeltower/pkg/network"
	"github.com/CESARBR/knot-babeltower/pkg/thing/mocks"
	"github.com/stretchr/testify/assert"
)

type SetDataTestCase struct {
	name                        string
	authorization               string
	thingID                     string
	data                        []network.Data
	expectedThing               *entities.Thing
	expectedThingError          error
	expectedUpdateDataResponse  error
	fakeLogger                  *mocks.FakeLogger
	fakeThingProxy              *mocks.FakeThingProxy
	fakePublisher               *mocks.FakePublisher
}

var sdCases = []SetDataTestCase{
	{
		"authorization key not provided",
		"",
		"",
		nil,
		nil,
		nil,
		nil,
		&mocks.FakeLogger{},
		&mocks.FakeThingProxy{},
		&mocks.FakePublisher{},
	},
	{
		"failed to authenticate with provided key",
		"authorization-key",
		"fc3fcf912d0c290a",
		nil,
		nil,
		errors.New("Invalid credentials"),
		nil,
		&mocks.FakeLogger{},
		&mocks.FakeThingProxy{},
		&mocks.FakePublisher{},
	},
	{
		"thing doesn't exists on thing's service",
		"authorization-key",
		"fc3fcf912d0c290a",
		nil,
		nil,
		errors.New("Thing fc3fcf912d0c290a not found"),
		nil,
		&mocks.FakeLogger{},
		&mocks.FakeThingProxy{},
		&mocks.FakePublisher{},
	},
	{
		"thing successfully obtained from the thing's service",
		"authorization-key",
		"fc3fcf912d0c290a",
		[]network.Data{network.Data{SensorID: 2, Value: false}},
		&entities.Thing{
			ID:    "fc3fcf912d0c290a",
			Token: "token",
			Name:  "thing",
			Schema: []entities.Schema{
				{
					SensorID:  2,
					ValueType: 1,
					Unit:      0,
					TypeID:    65521,
					Name:      "Test",
				},
			},
		},
		nil,
		nil,
		&mocks.FakeLogger{},
		&mocks.FakeThingProxy{},
		&mocks.FakePublisher{},
	},
	{
		"thing hasn't schema for the requested sensor",
		"authorization-key",
		"fc3fcf912d0c290a",
		[]network.Data{network.Data{SensorID: 2, Value: false}},
		&entities.Thing{
			ID:    "fc3fcf912d0c290a",
			Token: "token",
			Name:  "thing",
			Schema: []entities.Schema{
				{
					SensorID:  0,
					ValueType: 3,
					Unit:      0,
					TypeID:    65521,
					Name:      "Test",
				},
			},
		},
		nil,
		nil,
		&mocks.FakeLogger{},
		&mocks.FakeThingProxy{},
		&mocks.FakePublisher{},
	},
	{
		"failed to send quest data command to message queue",
		"authorization key",
		"fc3fcf912d0c290a",
		[]network.Data{network.Data{SensorID: 2, Value: false}},
		&entities.Thing{
			ID:    "fc3fcf912d0c290a",
			Token: "token",
			Name:  "thing",
			Schema: []entities.Schema{
				{
					SensorID:  2,
					ValueType: 3,
					Unit:      0,
					TypeID:    65521,
					Name:      "Test",
				},
			},
		},
		nil,
		errors.New("Failed to send request data message"),
		&mocks.FakeLogger{},
		&mocks.FakeThingProxy{},
		&mocks.FakePublisher{},
	},
	{
		"request data command successfully sent",
		"authorization key",
		"fc3fcf912d0c290a",
		[]network.Data{network.Data{SensorID: 2, Value: false}},
		&entities.Thing{
			ID:    "fc3fcf912d0c290a",
			Token: "token",
			Name:  "thing",
			Schema: []entities.Schema{
				{
					SensorID:  2,
					ValueType: 3,
					Unit:      0,
					TypeID:    65521,
					Name:      "Test",
				},
			},
		},
		nil,
		nil,
		&mocks.FakeLogger{},
		&mocks.FakeThingProxy{},
		&mocks.FakePublisher{},
	},
}

func TestSetData(t *testing.T) {
	for _, tc := range sdCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.fakeThingProxy.
				On("GetThing", tc.authorization, tc.thingID).
				Return(tc.expectedThing, tc.expectedThingError).
				Maybe()
			tc.fakePublisher.
				On("SendUpdateData", tc.thingID, tc.data).
				Return(tc.expectedUpdateDataResponse).
				Maybe()
		})

		thingInteractor := NewThingInteractor(tc.fakeLogger, tc.fakePublisher, tc.fakeThingProxy, nil)
		err := thingInteractor.UpdateData(tc.authorization, tc.thingID, tc.data)
		if tc.authorization == "" {
			assert.EqualError(t, err, "authorization key not provided")
		}

		tc.fakeThingProxy.AssertExpectations(t)
		tc.fakePublisher.AssertExpectations(t)
	}
}
