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
