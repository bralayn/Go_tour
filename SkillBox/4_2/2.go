package main

import "fmt"

func main() {
	var speed, distance int
	allDistance := 200
	time := 2

	fmt.Print("Input average speed:")
	fmt.Scan(&speed)

	// Calculate distance
	distance = speed * time
	if distance < allDistance {
		fmt.Println("You drove:", distance, "km")
		goto end
	}
	fmt.Println("-- You arrived ----")
end:
}
