// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"encoding/json"
	"errors"

	"github.com/danielskowronski/GoQingpingIoTMQTT/internal/model"
)

func (rm *RawMessage) IsJSON() bool {
	return len(rm.RawMessage) > 0 && rm.RawMessage[0] == model.ProtocolJSONHeader[0]
}
func (rm *RawMessage) DecodeJSON() (*Message, error) {
	msg := &Message{
		DeviceID: rm.DeviceID,
		Protocol: model.ProtocolJSON,
	}
	var m map[string]any
	err := json.Unmarshal(rm.RawMessage, &m)
	if err != nil {
		return nil, err
	}
	cmd, cmd_ok := m[model.ProtocolHEXTypeField].(float64)
	if !cmd_ok {
		return nil, err
	}
	msg.Verb = int16(cmd)

	verbInfo, verbOK := model.VerbInfoMap[msg.Verb]
	if !verbOK {
		switch msg.Verb {
		case model.Verb12:
			if rm.MessageDirection == model.MessageDirectionDeviceToServer {
				msg.Type = model.MessageTypeMeasurement
			} else {
				msg.Type = model.MessageTypeRequestFromServer
			}
		case model.Verb17:
			if rm.MessageDirection == model.MessageDirectionDeviceToServer {
				msg.Type = model.MessageTypeMeasurement
			} else {
				msg.Type = model.MessageTypeRequestFromServer
			}
		case model.Verb28:
			if rm.MessageDirection == model.MessageDirectionDeviceToServer {
				msg.Type = model.MessageTypeResponseFromDevice
			} else {
				msg.Type = model.MessageTypeRequestFromServer
			}
		default:
			return nil, errors.New("unknown command")
		}
	} else {
		msg.Type = verbInfo.Type
		if verbInfo.Protocol != model.ProtocolJSON {
			return nil, errors.New("protocol type attached to command does not match expected protocol")
		}
		if rm.MessageDirection != verbInfo.Direction {
			return nil, errors.New("command direction does not match expected direction")
		}
	}
	for key, value := range m {
		if key == model.ProtocolHEXTypeField {
			continue
		}
		msg.Fields[key] = value
	}

	return msg, nil
}
