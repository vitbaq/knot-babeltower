package mocks

import (
	"github.com/CESARBR/knot-babeltower/pkg/thing/entities"
	"github.com/stretchr/testify/mock"
	"github.com/CESARBR/knot-babeltower/pkg/network"
)

// FakeConnector represents a mocking type for the connector service
type FakeConnector struct {
	mock.Mock
	SendError error
	RecvError error
}

// SendRegisterDevice provides a mock function to send register device command to connector
func (fc *FakeConnector) SendRegisterDevice(id, name string) (err error) {
	ret := fc.Called(id, name)
	return ret.Error(0)
}

// SendUpdateSchema provides a mock function to send an update schema command to connector
func (fc *FakeConnector) SendUpdateSchema(id string, schemaList []entities.Schema) (err error) {
	ret := fc.Called(id, schemaList)
	return ret.Error(0)
}

// SendPublishData provides a mock function to send a publish data message to connector
func (fc *FakeConnector) SendPublishData(id string, data []network.Data) error {
	ret := fc.Called(id, data)
	return ret.Error(0)
}
