// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package model

// MeasurementField... contain read-only fields of measurement field in report

const (
	MeasurementFieldValue          string = "value"             // int
	MeasurementFieldStatus         string = "status"            // int, see MeasurementFieldStatus... constants below
	MeasurementFieldLevel          string = "level"             // int
	MeasurementFieldUnit           string = "unit"              // string
	MeasurementFieldStatusDuration string = "status_duration"   // int
	MeasurementFieldStatusStart    string = "status_start_time" // int
)

const (
	MeasurementFieldStatusNormal      int = 0
	MeasurementFieldStatusAbnormal    int = 1
	MeasurementFieldStatusInitialized int = 2
)
