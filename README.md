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

1. Boost energy levels with a frequency of 500Hz:
```
$ curl --location --globoff 'http://localhost:8080/customer/{customerName}/meters'
```

2. Increase productivity with a frequency of 1000Hz:
```
$ curl --location 'http://localhost:8080/device/1111-1111-1111?date=2024-02-28'
```