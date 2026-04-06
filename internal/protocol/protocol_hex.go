// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"encoding/binary"
	"errors"

	"github.com/danielskowronski/GoQingpingIoTMQTT/internal/model"
)

func (rm *RawMessage) IsHEX() bool {
	return len(rm.RawMessage) > 2 &&
		rm.RawMessage[0] == byte(model.ProtcolHEXHeader>>8) &&
		rm.RawMessage[1] == byte(model.ProtcolHEXHeader&0xFF)
}
func (rm *RawMessage) DecodeHEX() (*Message, error) {
	msg := &Message{
		DeviceID: rm.DeviceID,
		Protocol: model.ProtocolHEX,
	}

	if len(rm.RawMessage) < 7 {
		return nil, errors.New("raw HEX message is too short")
	}
	cmd := rm.RawMessage[2]
	msg.Verb = int16(cmd)
	verbInfo, ok := model.VerbInfoMap[int16(cmd)]
	if !ok {
		return nil, errors.New("unknown command")
	}
	if msg.Protocol != verbInfo.Protocol {
		return nil, errors.New("protocol type attached to command does not match expected protocol")
	}
	if rm.MessageDirection != verbInfo.Direction {
		return nil, errors.New("command direction does not match expected direction")
	}
	msg.Type = verbInfo.Type

	length := int(binary.LittleEndian.Uint16(rm.RawMessage[3:4]))
	if len(rm.RawMessage) != 2+length { // 2 bytes for checksum
		return nil, errors.New("raw HEX message length does not match expected length")
	}
	pos := 5 // 2 bytes for header, 1 byte for cmd, 2 bytes for length
	for pos < length-3 {
		field_key := rm.RawMessage[pos]
		field_len := binary.LittleEndian.Uint16(rm.RawMessage[pos+1 : pos+3])
		if pos+3+int(field_len) > len(rm.RawMessage) {
			return nil, errors.New("field length exceeds remaining message length")
		}
		field_val := rm.RawMessage[pos+3 : pos+3+int(field_len)]
		msg.Fields[field_key] = field_val

		pos += 3 + int(field_len)
	}
	if pos != length {
		return nil, errors.New("final position does not match expected length")
	}

	return msg, nil
}
