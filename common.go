package main

import "fmt"

func mapBasics() {
	var cars = map[string]uint16{"Benz": 1942, "Corola": 1852, "G-Wagon": 1507}

	for car, year := range cars {
		fmt.Println(car, ": ", year)
	}

	//creating map with 'make' method
	var balls = make(map[string]string)
	balls["football"] = "white"
	balls["tennis ball"] = "yellow"
	balls["rugby ball"] = "grey"

	delete(balls, "tennis ball")

	fmt.Println("number of balls=", len(balls))
}
