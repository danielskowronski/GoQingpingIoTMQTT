package protocol

import "github.com/danielskowronski/GoQingpingIoTMQTT/internal/model"

func IsValidCommandJson(cmd int16, direction model.MessageDirection) bool {
	if cmd == model.CmdDeviceLogReport || cmd == model.CmdBindingStatus || cmd == model.CmdBindingStatusThirdPartyDevice {
		return true // FIXME: those are untested and improperly documented in spec, need to investigate
	}
	switch direction {
	case model.MessageDirectionDeviceToServer:
		return cmd == model.CmdBleConnectionRequest ||
			cmd == model.CmdBleDisconnectionRequest ||
			cmd == model.CmdOpenBleNotificationRequest ||
			cmd == model.CmdCloseBleNotificationRequest ||
			cmd == model.CmdBleDataWithResponse ||
			cmd == model.CmdBleDataRead ||
			cmd == model.CmdDeviceListResponse ||
			cmd == model.CmdTemporaryReportSetting ||
			cmd == model.CmdReconnectMqtt ||
			cmd == model.CmdBleDataWithoutResponse ||
			cmd == model.CmdModifyMqttConnectionSetting ||
			cmd == model.CmdModifyDataReportInterval ||
			cmd == model.CmdHistoryDataReportResponse ||
			cmd == model.CmdOtaCommand ||
			cmd == model.CmdDeviceListResponseWithName ||
			cmd == model.CmdReadDeviceSettingRequest
	case model.MessageDirectionServerToDevice:
		return cmd == model.CmdBleNotificationResponse ||
			cmd == model.CmdBleDataResponse ||
			cmd == model.CmdBroadcastData ||
			cmd == model.CmdDeviceListRequest ||
			cmd == model.CmdHeartbeat ||
			cmd == model.CmdSensorDataReportRealtime ||
			cmd == model.CmdSensorDataReportHistory ||
			cmd == model.CmdOtaCommandResponse ||
			cmd == model.CmdDeviceListRequestWithName
	default:
		return false
	}

}
