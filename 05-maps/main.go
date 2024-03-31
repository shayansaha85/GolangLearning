package main

import "fmt"

type mappy map[string]string

func main() {

	/*way 1 of creating map*/
	// var colors map[string]string

	/*way 2 of creating map*/
	// colors := make(map[string]string)

	/*way 3 of creating map*/
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	colors_mappy := mappy{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// colors["white"] = "#ffffff"
	// colors["red"] = "#ff0000"

	// delete(colors, "red")

	// fmt.Println(colors)

	printMap(colors)
	fmt.Println()
	colors_mappy.print()

}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println(color, "\t:\t", hex)
	}
}

func (m mappy) print() {
	for color, hex := range m {
		fmt.Println(color, "\t:\t", hex)
	}
}
