# Go Qingping IoT MQTT

*work in progress*

Go lang library and utilities for working with Qingping IoT MQTT protocols - JSON and binary ("HEX"). It's a rebuild of Python-based [qingping-iot-mqtt](https://github.com/danielskowronski/qingping-iot-mqtt) made to simplify code that grew too complex.

## Background

Qingping IoT devices working over Wi-Fi (and some home Qingping+ devices) can either push data to Qingping IoT Cloud (see [qingping-iot-cloud](https://github.com/danielskowronski/qingping-iot-cloud) library for details) or be "privatized", i.e. change MQTT endpoint to your own MQTT broker. Each model of Qingping IoT Wi-Fi device use either JSON or binary ("HEX") protocol and this cannot be changed. 

References:

- [MQTT JSON](https://developer.qingping.co/private/communication-protocols/public-mqtt-json)
- [MQTT HEX](https://qingping.feishu.cn/docx/BlYOdJVRQobV0ox6SNZcV8V6nZT)

## Scope

This library and set of utilities are designed for both protocols.

Ultimately, a Home Assistant add-on (which will translate proprietary Qingping payloads to HA-discoverable device states and commands over MQTT) will be created.

For now, main goals are to implement:

1. decoder for raw protocols into simple structs that can be easily dumped to human-readable formats to debug protocol manually
   1. offline - from CLI
   2. online - actively subscribing to MQTT topic from device and logging values to console
   3. online - additionally, logging decoded measurement data to Victoria Metrics
2. encoder for server-to-device commands (mainly device settings)
   1. offline - from CLI flags to raw payload
   2. online - accepting same flags, but publishing messages itself, awaiting for response and decoding them

MQTT broker configuration is not part of this project.

Currently, I only have CO2 sensor (HEX protocol) and Air Monitor Lite (JSON protocol), so coverage may be limited.
