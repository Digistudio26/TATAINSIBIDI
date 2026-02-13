<<<<<<< HEAD
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


=======
# TATAINSIBIDI - Letter Quest Game

🌟 **Letter Quest: TATAINSIBIDI Edition** 🌟

Welcome to **Letter Quest**, a fun terminal-based game where you match numbers to letters and earn royal blessings! Inspired by African culture, the game includes West African deities and celebrates your victories with an African mask.  

---

## 🎮 How to Play

1. Enter your heroic name.
2. You get **5 chances per round**.
3. Each chance allows **5 number selections** (from 1 to 26).
4. If a number matches its letter value, you score points.
   - **A → 1, B → 2, C → 3, ..., X → 24, Y → 25, Z → 26**
5. **3 or more matches** in a round = **WIN** (score multiplied by 5).
6. Multiples of 5 or number 26 give **royal blessings**.
7. After each round, you can view your stats, retry, or access a free try round.

---

## 🛠 Features

- Interactive terminal gameplay.
- **5 chances per round**, each with 5 selections.
- Color-coded results:
  - ✅ Correct match (Green)
  - ⚠️ Selected but incorrect (Yellow)
  - ❌ Unselected numbers (Default)
- Win rewards include:
  - African mask display
  - Blessings from West African deities: Ogun, Shango, Yemoja, Orunmila, Obatala
- Help instructions accessible from the menu.
- Free try round to practice.

---

## 📊 Game Menu Options

1️⃣ Play round  
2️⃣ View scores  
3️⃣ Help / Instructions  
4️⃣ Free try (bonus round)  
5️⃣ Exit game  

---

## 💻 Running the Game

Make sure you have [Go](https://golang.org/) installed.  

```bash
git clone https://github.com/Digistudio26/TATAINSIBIDI.git
cd TATAINSIBIDI
go run main.go
>>>>>>> dd1e240ca8d8c54263894e83fa647d24ec9ca3c6
