# Diagrams

## Sequence Diagrams for API Calls

### hue#GetSensors returns (Sensors)

```mermaid
sequenceDiagram
    autonumber
    participant Homecontrol Client
    participant Homecontrol GRPC
    participant Philips Hue Bridge
    Homecontrol Client->>+Homecontrol GRPC: GetSensors()
    Homecontrol GRPC->>+Philips Hue Bridge: GetSensors()
    Philips Hue Bridge->>-Homecontrol GRPC: returns (Sensors)
    Homecontrol GRPC-->>-Homecontrol Client: returns (Sensors)
```

### hue#GetGroups returns (Groups)

```mermaid
sequenceDiagram
    autonumber
    participant Homecontrol Client
    participant Homecontrol GRPC
    participant Philips Hue Bridge
    Homecontrol Client->>+Homecontrol GRPC: GetGroups()
    Homecontrol GRPC->>+Philips Hue Bridge: GetGroups()
    Philips Hue Bridge->>-Homecontrol GRPC: returns (Groups)
    Homecontrol GRPC-->>-Homecontrol Client: returns (Groups)
```

### hue#SwitchOn returns (Sensors)

```mermaid
sequenceDiagram
    autonumber
    participant Homecontrol Client
    participant Homecontrol GRPC
    participant Philips Hue Bridge
    Homecontrol Client->>+Homecontrol GRPC: SwitchOn()
    Homecontrol GRPC->>+Philips Hue Bridge: SwitchOn()
    Philips Hue Bridge->>-Homecontrol GRPC: returns (Success)
    Homecontrol GRPC-->>-Homecontrol Client: returns (Success)
```


### hue#SwitchOff returns (Sensors)

```mermaid
sequenceDiagram
    autonumber
    participant Homecontrol Client
    participant Homecontrol GRPC
    participant Philips Hue Bridge
    Homecontrol Client->>+Homecontrol GRPC: SwitchOff()
    Homecontrol GRPC->>+Philips Hue Bridge: SwitchOff()
    Philips Hue Bridge->>-Homecontrol GRPC: returns (Success)
    Homecontrol GRPC-->>-Homecontrol Client: returns (Success)
```

### hue#Blink

```mermaid
sequenceDiagram
    autonumber
    participant Homecontrol Client
    participant Homecontrol GRPC
    participant Philips Hue Bridge
    Homecontrol Client->>+Homecontrol GRPC: Blink()
    Homecontrol GRPC->>+Philips Hue Bridge: Blink()
    Philips Hue Bridge->>-Homecontrol GRPC: returns
    Homecontrol GRPC-->>-Homecontrol Client: returns
```