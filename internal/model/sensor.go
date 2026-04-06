// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package model

type SensorStatus string

const (
	SensorStatusNormal  SensorStatus = "normal"
	SensorStatusLow     SensorStatus = "low"
	SensorStatusHigh    SensorStatus = "high"
	SensorStatusUnknown SensorStatus = "unknown"
)

type SensorSource string

const (
	SensorSourcePeriodic SensorSource = "periodic"
	SensorSourceEvent    SensorSource = "event"
	SensorSourceRealtime SensorSource = "realtime"
)
