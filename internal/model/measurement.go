// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Measurement struct {
	SensorType     SensorType
	Timestamp      time.Time
	Value          decimal.Decimal
	SensorTypeInfo SensorTypeInfo
}
