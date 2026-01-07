package main

import (
	"fmt"

	"github.com/PackMeister/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	// Or by
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

}
