// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package model

type ProtocolType string

const (
	ProtocolJSON       ProtocolType = "json"
	ProtocolJSONHeader string       = "{"

	ProtocolHEX          ProtocolType = "hex"
	ProtcolHEXHeader     uint16       = 0x4347
	ProtocolHEXTypeField string       = "type"
)
