# Weather Station

A simple Go program that processes incoming weather station telemetry, reconstructs full sensor state from partial updates, and outputs the current known state on demand.

This project simulates a single weather station sending frequent incremental updates, with occasional full snapshots, while keeping the server-side implementation minimal and reliable.

---

## Overview

Weather stations send meteorological data every minute, but only include values that have changed since the previous message to minimize payload size.

This application:

- Maintains the latest known state of each sensor
- Accepts partial updates
- Reconstructs and outputs the full sensor state on request
- Uses only in-memory state (no external storage or tooling)

---

## Supported Sensors

Only the following sensor IDs are recognized. All others are ignored.

| ID | Key            |
|----|----------------|
| 1  | airTemp        |
| 2  | airPressure    |
| 7  | precipitation  |
| 11 | windSpeed      |
| 12 | windDirection  |
| 13 | humidity       |
| 14 | dewPoint       |
| 15 | soilMoisture   |
| 22 | cloudCover     |

All values are floating-point numbers. Missing data is represented as `NULL`.

---

## Program Behavior

### Startup


