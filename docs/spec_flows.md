<!--
SPDX-FileCopyrightText: 2026-present Daniel Skowroński <GoQingpingIoTMQTT@skowronski.cloud>

SPDX-License-Identifier: BSD-3-Clause
-->
# Communication flows - according to spec

## HEX Protocol

### Measurements

#### Standard measurements

```mermaid
sequenceDiagram
    Device->>Server: 0x3B periodic payload with single measurement pack
    Device->>Server: 0x31 periodic payload with series of measurement packs
```

#### Event/alert

```mermaid
sequenceDiagram
    Device->>Server: 0x34 event/alert payload with single measurement pack
```

### Configuration and commands

#### Settings

```mermaid
sequenceDiagram
    Server->>Device: 0x32 list of new settings (may be empty)
    Device->>Server: 0x39 current settings (after apply)
```

#### Initiate firmware upgrade

```mermaid
sequenceDiagram
    Server->>Device: 0x33 init firmware upgrade with specified parameters
```

#### Reconfigure network

```mermaid
sequenceDiagram
    Server->>Device: 0x3A configure new network access parameters
```

## JSON protocol

### Pairing flows

#### BLE flow

*not fully understdood*

```mermaid
sequenceDiagram
    Server->>Device: cmd=1 Server sends BLE connection request
    Server->>Device: cmd=2 Server sends BLE disconnection request

    Server->>Device: cmd=3 Server sends open BLE notification request
    Server->>Device: cmd=4 Server sends close BLE notification request
    Device->>Server: cmd=5 Response for BLE notification from device

    Server->>Device: cmd=6 Server sends BLE data(with response)
    Server->>Device: cmd=7 Server reads BLE data
    Device->>Server: cmd=8 Response for BLE data from device
    Device->>Server: cmd=9 Broadcast data from device

    Server->>Device: cmd=15 Server sends BLE data(without response)
```

#### Device list flow

*not fully understdood, HomeKit?*

```mermaid
sequenceDiagram
    Device->>Server: cmd=10 Device requests device list
    Server->>Device: cmd=11 Server response device list
    Device->>Server: cmd=25 Device Requests Device List(with device name)
    Server->>Device: cmd=26 Server Responses Device List(with device name)
```

### Configuration and commands

#### Temporary rapid reporting

```mermaid
sequenceDiagram
    Server->>Device: cmd=12 Server send setting for temporary report and duration time
```

#### MQTT

```mermaid
sequenceDiagram
    Server->>Device: cmd=14 Reconnect MQTT (no params)
    Server->>Device: cmd=16 Modify MQTT connection setting
```

#### Settings

set:

```mermaid
sequenceDiagram
    Server->>Device: cmd=17 Modify data report interval (or other settings)
    Device->>Server: cmd=28 Device Responses Device Setting
```

just read:

```mermaid
sequenceDiagram
    Server->>Device: cmd=28 Server Requests To Read Device Setting
    Device->>Server: cmd=28 Device Responses Device Setting
```

#### Notifications

```mermaid
sequenceDiagram
    Server->>Device: cmd=32 Notification Rule Settings
    Device->>Server: cmd=28 Device Responses Device Setting
```

#### Alarm

```mermaid
sequenceDiagram
    Server->>Device: cmd=40 Retrieve Device Alarm List
    Device->>Server: cmd=40 Device Reports Alarm List
    Server->>Device: cmd=40 Alarm Settings
```

#### OTA

```mermaid
sequenceDiagram
    Server->>Device: cmd=23 Server Sends OTA Command
    Device->>Server: cmd=24 Device Responses For OTA Command
```

### Measurements

#### Heartbeat

*transmits WiFi info and confirms device is online*

```mermaid
sequenceDiagram
    Device->>Server: cmd=13 Heartbeat package
```

#### Real-time data

```mermaid
sequenceDiagram
    Device->>Server: cmd=12 Sensor data report
```

#### History data

```mermaid
sequenceDiagram
    Device->>Server: cmd=17 Sensor data report
    Server->>Device: cmd=18 Server Response For History Data Report
```

*response is obligatory, as it acts as primitive NTP*

### Completely unknown

- `cmd=19 Device Report Log`
- `cmd=20 Binding Status`
- `cmd=27 Binding Status For Third Part's Device`
