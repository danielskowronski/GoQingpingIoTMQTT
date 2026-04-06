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
	Protocol         model.Protocol
	MessageDirection model.MessageDirection
	RawMessage       []byte // encoded string for JSON, bytes for HEX

	decoder RawMessageDecoder
}

type RawMessageDecoder interface {
	Decode() (*ParsedMessage, error)
}

func (rm *RawMessage) Decode() (*ParsedMessage, error) {
	if rm.decoder == nil {
		return nil, errors.New("no decoder configured")
	}
	return rm.decoder.Decode()
}
