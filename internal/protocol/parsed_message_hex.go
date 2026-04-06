// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import "github.com/danielskowronski/GoQingpingIoTMQTT/internal/model"

func IsValidCommandHex(cmd int16, direction model.MessageDirection) bool {
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
