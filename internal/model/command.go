package model

const (
	// HEX
	CmdDataUploading          int16 = 0x31
	CmdConfigurationSending   int16 = 0x32
	CmdFirmwareUpgrade        int16 = 0x33
	CmdEventReporting         int16 = 0x34
	CmdConfigurationReporting int16 = 0x39
	CmdNetworkAccessSetting   int16 = 0x3A
	CmdRealtimeDataUploading  int16 = 0x3B
	// JSON
	CmdBleConnectionRequest          int16 = 1  // "Server sends BLE connection request"
	CmdBleDisconnectionRequest       int16 = 2  // "Server sends BLE disconnection request"
	CmdOpenBleNotificationRequest    int16 = 3  // "Server sends open BLE notification request"
	CmdCloseBleNotificationRequest   int16 = 4  // "Server sends close BLE notification request"
	CmdBleNotificationResponse       int16 = 5  // "Response for BLE notification from device"
	CmdBleDataWithResponse           int16 = 6  // "Server sends BLE data(with response)"
	CmdBleDataRead                   int16 = 7  // "Server reads BLE data"
	CmdBleDataResponse               int16 = 8  // "Response for BLE data from device"
	CmdBroadcastData                 int16 = 9  // "Broadcast data from device"
	CmdDeviceListRequest             int16 = 10 // "Device requests device list"
	CmdDeviceListResponse            int16 = 11 // "Server response device list"
	CmdTemporaryReportSetting        int16 = 12 // "Server send setting for temporary report and duration time"
	CmdHeartbeat                     int16 = 13 // "Heartbeat package"
	CmdReconnectMqtt                 int16 = 14 // "Reconnect MQTT"
	CmdBleDataWithoutResponse        int16 = 15 // "Server sends BLE data(without response)"
	CmdModifyMqttConnectionSetting   int16 = 16 // "Modify MQTT connection setting"
	CmdModifyDataReportInterval      int16 = 17 // "Modify data report interval"
	CmdSensorDataReportRealtime      int16 = 12 // "Sensor data report - real-time data"
	CmdSensorDataReportHistory       int16 = 17 // "Sensor data report - history data"
	CmdHistoryDataReportResponse     int16 = 18 // "Server response for history data report in type 17"
	CmdDeviceLogReport               int16 = 19 // "Device report log"
	CmdBindingStatus                 int16 = 20 // "Binding status"
	CmdOtaCommand                    int16 = 23 // "Server sends OTA command"
	CmdOtaCommandResponse            int16 = 24 // "Devices response for OTA command"
	CmdDeviceListRequestWithName     int16 = 25 // "Device requests device list(with device name)"
	CmdDeviceListResponseWithName    int16 = 26 // "Server responses device list(with device name)"
	CmdBindingStatusThirdPartyDevice int16 = 27 // "Binding status for third part's device"
	CmdReadDeviceSettingRequest      int16 = 28 // "Request to read device setting"
)
