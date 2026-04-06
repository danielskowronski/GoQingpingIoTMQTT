<!--
SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>

SPDX-License-Identifier: BSD-3-Clause
-->
# Struct schema - first proposal

- single measurement - struct with:
  - timestamp - proper datetime
  - sensor type - enum
  - sensor value - decimal
  - sensor unit - enum
  - status - enum of (OK, ALERT_LOW, ALERT_HIGH)
  - source - enum of (PERIODIC, REAL_TIME, ALERT)
- measurement pack - iterable of single measurement
- command field:
  - field - enum
  - value - any
- command field of type X (int, float, string?):
  - field - enum
  - value - that specific type
- command - simple struct of:
  - command type - enum
  - list of command fields
- protocol message:
  - protocol - enum of (JSON, HEX)
  - message type - enum of (MEASUREMENT, ACK_FROM_SERVER, REQUEST_FROM_SERVER, RESPONSE_FROM_DEVICE, OTHER_FROM_SERVER, OTHER_FROM_DEVICE) where last 2 are for unimplemented 
  - device ID (from MQTT queue)
  - direction - enum of (DEVICE_TO_SERVER, SERVER_TO_DEVICE)
  - command code
  - fields - mapping of string to raw extracted payload (json - simple strings, hex - validated and extracted fields)
- protocol - JSON and HEX implementing:
  - decodeProtocolMessage - accepts byte stream and outputs object with structure split and only command/type decoded
  - protocolMessageToMeasurementPack - only for MEASUREMENT, returns iterable of structs + response message if needed
  - ackToProtocolMessage - takes no params and prepares protocol message with current timestamp - produces - ACK_FROM_SERVER
  - commandFieldsToCommand - takes list of command fields, checks whether they belong to single command in specific protocol, selects that command and assembles list of fields
  - commandToProcolMessage - takes command and prepares all fields of protocol message - produces REQUEST_FROM_SERVER or OTHER_FROM_SERVER
  - encodeProtocolMessage - converts struct to json or hex, this should be opposite of decodeProtocolMessage
  - protocolMessageToCommandResponse - only for RESPONSE_FROM_DEVICE
  - protocolMessageDump - for all messages, including unhandled OTHER_FROM_DEVICE
- device - constructed from device model enum, MQTT credentials, device ID and MQTT paths (up,down), implements MQTT subscriber and method to publish, has list of valid sensors (some data is duplicated like CO2 vs pressure and device must interpret values) and list of valid commands; some are taken from shared list for protocol, some are model-specific; implements:
  - loop for MQTT subscription - switch/case selects which function will handle message:
    - measurement -> calls registerable callback method, e.g. just log
    - response_from_device -> calls internal response handler
    - other -> logs on console
  - command send - accepts list of command params, verifies initially with capabilities of device, gets expected message code from command enum, puts on some list, renders raw command, publishes over MQTT, awaits for response for some time, when response comes back, decodes it fully and returns to requestor
  - internal response handler - checks if it was expected and if so, wakes requestor
