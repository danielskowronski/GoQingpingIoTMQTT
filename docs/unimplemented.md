<!--
SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>

SPDX-License-Identifier: BSD-3-Clause
-->
# Unimplemented

This doc lists features or parts of spec that cannot be implemented in near future with current understanding of protocol. Payload dumps are welcome.

## HEX protocol

### KEYs that don't immediately make sense

- `0x15` "Timestamp" - this is not shown in any example; doesn't make sense in reports (timestamp is part of historical data values) or commands
- `0x2D` "Historical data reporting (air pressure 0.01 pa)" - this is not shown in any example; I don't have HEX pressure sensor to validate
- `0x2E` "Real-time data reporting (air pressure 0.01 pa)" - this is not shown in any example; I don't have HEX pressure sensor to validate
- `0x32` "AUX temperature and humidity data" - this is not shown in any example; I don't have HEX AUX-capable sensor to validate
- `0x38` "Product ID" - this is not shown in any example; it's unclear whether it's command or report
- `0x66` "Event alert" - this is not shown in any example
- `0xC8` "User defined function" - this is not shown in any example

## JSON protocol

### BLE communication

cannot test it; fields:

- `mac` (String) "BLE MAC Address"
- `buffer` (String) "BLE data content"
- `length` (Integer) "BLE data content length"
- `srv_uuid` (String) "BLE Service UUID"
- `chr_uuid` (String) "BLE Feature UUID"
- `adv_data` (String) "BLE broadcast data"
- `rssi` (Integer) "RSSI"

### other fields

- `timeout` (Integer) "Connecting timeout(in second). Note: After make connection successfully if there is no data communication in 10 seconds, the connection will be closed"; this fields is not shown in any example

### device binding/list flow

cannot test it; fields:

- `dev_list` "Device list"
- `homekit_dev_list`

### notifications (scenes)

### alarms (ringtones)
