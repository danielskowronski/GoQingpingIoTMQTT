package model

type MessageType string

const (
	MessageTypeMeasurement        MessageType = "measurement"
	MessageTypeAck                MessageType = "ack"
	MessageTypeRequestFromServer  MessageType = "request_from_server"
	MessageTypeResponseFromDevice MessageType = "response_from_device"
	MessageTypeOtherFromDevice    MessageType = "other_from_device"
	MessageTypeOtherFromServer    MessageType = "other_from_server"
)

type MessageDirection string

const (
	MessageDirectionDeviceToServer MessageDirection = "device_to_server"
	MessageDirectionServerToDevice MessageDirection = "server_to_device"
	// there's no ambiguity, this field depends on "up" vs "down" topic or user intention
)
