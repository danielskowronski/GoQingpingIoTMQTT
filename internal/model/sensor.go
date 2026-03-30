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
