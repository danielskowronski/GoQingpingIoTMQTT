// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"errors"

	"github.com/danielskowronski/GoQingpingIoTMQTT/internal/model"
)

type Message struct {
	DeviceID string
	Protocol model.ProtocolType
	Type     model.MessageType
	Verb     int16                       // command ID or message type
	Fields   map[interface{}]interface{} // JSON will hold string, HEX will hold int

	protocolHandler ProtocolHandler
}

type ProtocolHandler interface {
	Encode() (*RawMessage, error)
}

func (pm *Message) Encode() (*RawMessage, error) {
	if pm.protocolHandler == nil {
		return nil, errors.New("no protocol handler configured")
	}
	return pm.protocolHandler.Encode()
}
