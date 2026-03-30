package protocol

import (
	"errors"

	"github.com/danielskowronski/GoQingpingIoTMQTT/internal/model"
)

type ParsedMessage struct {
	DeviceID string
	Protocol model.Protocol
	Type     model.MessageType
	Command  int16                       // command ID or message type
	Fields   map[interface{}]interface{} // JSON will hold string, HEX will hold int

	encoder   ParsedMessageEncoder
	validator ParsedMessageValidator
}

type ParsedMessageEncoder interface {
	Encode() (*RawMessage, error)
}
type ParsedMessageValidator interface {
	Validate() error
}

func (pm *ParsedMessage) Encode() (*RawMessage, error) {
	if pm.encoder == nil {
		return nil, errors.New("no encoder configured")
	}
	return pm.encoder.Encode()
}

func (pm *ParsedMessage) Validate() error {
	if pm.encoder == nil {
		return errors.New("no encoder configured")
	}
	return pm.validator.Validate()
}
