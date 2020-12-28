package main

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	var min int = 99999
	for i := 0; i < len(x); i++ {
		if x[i] < min {
			min = x[i]
		}
	}
	/* Another way using range (Python inspired)
	for i, v := range x {
		if i == 0 || v < min {
			min = v
		}
	}
	*/

	println("The minimum element is", min)
}
