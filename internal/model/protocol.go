// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package model

type Protocol string

const (
	ProtocolMQTT Protocol = "mqtt"
	ProtocolHEX  Protocol = "hex"

	ProtocolHEXTypeField string = "type"
)
