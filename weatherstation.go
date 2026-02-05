package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type sensor struct {
	ID    int
	Key   string
	Value *float64 // pointer = nullable //pointers default is nil.
}

func main() {
	fmt.Println(".....Weather-Station.....")

	// Sensor list (pointers!)
	sensors := []*sensor{
		{ID: 1, Key: "airTemp"},
		{ID: 2, Key: "airPressure"},
		{ID: 7, Key: "precipitation"},
		{ID: 11, Key: "windSpeed"},
		{ID: 12, Key: "windDirection"},
		{ID: 13, Key: "humidity"},
		{ID: 14, Key: "dewPoint"},
		{ID: 15, Key: "soilMoisture"},
		{ID: 22, Key: "cloudCover"},
	}

	// ID → sensor pointer
	sensorByID := make(map[int]*sensor)
	for _, s := range sensors {
		sensorByID[s.ID] = s
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		switch line {

		case "get":
			for _, s := range sensors {
				if s.Value == nil {
					fmt.Printf("%s:NULL\n", s.Key)
				} else {
					fmt.Printf("%s:%g\n", s.Key, *s.Value)
				}
			}

		case "clear":
			for _, s := range sensors {
				s.Value = nil
			}

		case "exit":
			fmt.Println("Exiting...")
			return

		default:
			parts := strings.Split(line, ",")
			if len(parts) != 2 {
				continue
			}

			id, err := strconv.Atoi(parts[0])
			if err != nil {
				continue
			}

			s, ok := sensorByID[id]
			if !ok {
				continue
			}

			if parts[1] == "NULL" {
				s.Value = nil
				continue
			}

			v, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				continue
			}

			s.Value = &v // pointer assignment
		}
	}
}