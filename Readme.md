# Weather CLI Application

A command-line interface tool that displays current weather and hourly forecast data using the WeatherAPI.com service.

## Features

- Current weather conditions
- Hourly forecast for the next 24 hours
- Color-coded rain probability (red text for >40% chance)
- Support for global locations

## Prerequisites

- Go 1.16 or higher
- WeatherAPI.com API key

## Installation

```bash
git clone https://github.com/Unnati-Gupta24/go-cli.git
cd go-cli
go mod tidy
```

## Usage

Run the application:

```bash
go run main.go
```

You'll be prompted to enter:
1. City name
2. Country name

## API Key

The application uses WeatherAPI.com. Get your API key at [WeatherAPI.com](https://www.weatherapi.com).

## Dependencies

- github.com/fatih/color - For colored terminal output

## Output Format

- Current weather: `[City], [Country]: [Temperature]C, [Condition]`
- Hourly forecast: `[Time] - [Temperature]C, [Rain Chance]%, [Condition]`

## Error Handling

The application handles:
- Invalid API responses
- Network errors
- Invalid location inputs
