// SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>
//
// SPDX-License-Identifier: BSD-3-Clause

package model

// event types, may be re-used by some protocols

const (
	// HEX
	EvtReportTemperatureHigher    int16 = 0x07
	EvtReportTemperatureLower     int16 = 0x08
	EvtReportHumidityHigher       int16 = 0x0A
	EvtReportHumidityLower        int16 = 0x0B
	EvtReportPressureHigher       int16 = 0x0D
	EvtReportPressureLower        int16 = 0x0E
	EvtReportRealTime             int16 = 0x14 // this is only reported
	EvtReportBatteryLow           int16 = 0x17 // this is only reported
	EvtReportTemperatureAuxHigher int16 = 0x29
	EvtReportTemperatureAuxLower  int16 = 0x2A
	EvtReportCO2Higher            int16 = 0x39
	EvtReportCO2Lower             int16 = 0x3A
	EvtReportTVOCHigher           int16 = 0x57
	EvtReportTVOCLower            int16 = 0x58
	EvtReportPM25Higher           int16 = 0x59
	EvtReportPM25Lower            int16 = 0x5A
	EvtReportPM10Higher           int16 = 0x5B
	EvtReportPM10Lower            int16 = 0x5C
	EvtReportNoiseHigher          int16 = 0x5D
	EvtReportNoiseLower           int16 = 0x5E
	EvtReportLightHigher          int16 = 0x5F
	EvtReportLightLower           int16 = 0x60
)
