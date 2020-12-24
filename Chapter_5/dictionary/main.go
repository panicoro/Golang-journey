package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
	// Because the element doesn't exist, we will see "" here.
	fmt.Println(elements["Un"])

	name, ok := elements["Un"]
	// Accessing an element of a map can return
	// two values instead of just one.
	// The first value is the result of the lookup,
	// the second tells us whether or not the lookup was
	fmt.Println(name, ok)

	// First, we try to get the value from the map.
	// Then, if it’s successful, we run the code
	// inside of the block.
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("This element doesn't exist")
	}

	otherElements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	fmt.Println(len(otherElements))

}
