package main

import "fmt"

func main() {

	/*way 1 of creating map*/
	// var colors map[string]string

	/*way 2 of creating map*/
	colors := make(map[string]string)

	/*way 3 of creating map*/
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	colors["white"] = "#ffffff"
	colors["red"] = "#ff0000"

	fmt.Println(colors)

	delete(colors, "red")

	fmt.Println(colors)
}
