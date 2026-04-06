// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"errors"

	"github.com/danielskowronski/GoQingpingIoTMQTT/internal/model"
)

type RawMessage struct {
	DeviceID         string
	MessageDirection model.MessageDirection
	RawMessage       []byte // encoded string for JSON, bytes for HEX
}

func (rm *RawMessage) Decode() (*Message, error) {
	if rm.IsJSON() {
		return nil, errors.New("JSON decoding not implemented yet")
	} else if rm.IsHEX() {
		return rm.DecodeHEX()
	}
	return nil, errors.New("unknown protocol header")
}
