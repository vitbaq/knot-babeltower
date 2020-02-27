package interactors

import (
	"errors"
	"testing"

	"github.com/CESARBR/knot-babeltower/pkg/thing/entities"
	"github.com/CESARBR/knot-babeltower/pkg/network"
	"github.com/CESARBR/knot-babeltower/pkg/thing/mocks"
	"github.com/stretchr/testify/assert"
)

type PubDataTestCase struct {
	name                         string
	authorization                string
	thingID                      string
	data                         []network.Data
	expectedThing                *entities.Thing
	expectedThingError           error
	expectedPublishDataConnector error
	fakeLogger                   *mocks.FakeLogger
	fakeThingProxy               *mocks.FakeThingProxy
	fakeConnector                *mocks.FakeConnector
}

var pubCases = []PubDataTestCase{
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
		&mocks.FakeConnector{},
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
		&mocks.FakeConnector{},
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
		&mocks.FakeConnector{},
	},
}

func TestPubData(t *testing.T) {
	for _, tc := range pubCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.fakeThingProxy.
				On("GetThing", tc.authorization, tc.thingID).
				Return(tc.expectedThing, tc.expectedThingError).
				Maybe()
			tc.fakeConnector.
				On("SendPublishData", tc.thingID, tc.data).
				Return(tc.expectedPublishDataConnector).
				Maybe()
		})

		thingInteractor := NewThingInteractor(tc.fakeLogger, nil, tc.fakeThingProxy, tc.fakeConnector)
		err := thingInteractor.UpdateData(tc.authorization, tc.thingID, tc.data)
		if tc.authorization == "" {
			assert.EqualError(t, err, "authorization key not provided")
		}

		tc.fakeThingProxy.AssertExpectations(t)
	}
}
