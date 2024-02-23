# Power Wave

## Installation

To install Power Wave, follow these steps:

1. Clone the repository:

```
$ git clone https://github.com/slatermorgan/powerwave.git
```

2. Navigate to the project directory:

```
$ cd power-wave
```

3. Build the project:

```
$ go build
```

4. Run the application:

```
$ go run main.go
```

## Example curls

1. Get a list of devices for a given customer name:
```
$ curl --location --globoff 'http://localhost:8080/customer/{customerName}/meters'
```

2. Get a device total power consumer at a given time:
```
$ curl --location 'http://localhost:8080/device/{deviceId}?date=2024-02-28'
```