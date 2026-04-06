// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import "github.com/danielskowronski/GoQingpingIoTMQTT/internal/model"

func IsValidCommandJson(cmd int16, direction model.MessageDirection) bool {
	if cmd == model.VerbDeviceLogReport || cmd == model.VerbBindingStatus || cmd == model.VerbBindingStatusThirdPartyDevice {
		return true // FIXME: those are untested and improperly documented in spec, need to investigate
	}
	switch direction {
	case model.MessageDirectionDeviceToServer:
		return cmd == model.VerbBleConnectionRequest ||
			cmd == model.VerbBleDisconnectionRequest ||
			cmd == model.VerbOpenBleNotificationRequest ||
			cmd == model.VerbCloseBleNotificationRequest ||
			cmd == model.VerbBleDataWithResponse ||
			cmd == model.VerbBleDataRead ||
			cmd == model.VerbDeviceListResponse ||
			cmd == model.VerbTemporaryReportSetting ||
			cmd == model.VerbReconnectMqtt ||
			cmd == model.VerbBleDataWithoutResponse ||
			cmd == model.VerbModifyMqttConnectionSetting ||
			cmd == model.VerbModifyDataReportInterval ||
			cmd == model.VerbHistoryDataReportResponse ||
			cmd == model.VerbOtaCommand ||
			cmd == model.VerbDeviceListResponseWithName ||
			cmd == model.VerbReadDeviceSettingRequest
	case model.MessageDirectionServerToDevice:
		return cmd == model.VerbBleNotificationResponse ||
			cmd == model.VerbBleDataResponse ||
			cmd == model.VerbBroadcastData ||
			cmd == model.VerbDeviceListRequest ||
			cmd == model.VerbHeartbeat ||
			cmd == model.VerbSensorDataReportRealtime ||
			cmd == model.VerbSensorDataReportHistory ||
			cmd == model.VerbOtaCommandResponse ||
			cmd == model.VerbDeviceListRequestWithName
	default:
		return false
	}

}
