package main

import "fmt"

func main() {

	// type array, object in memory with 8 int items
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("arr %v; arr p - %p\n", arr, &arr)

	// type slice is slice of array
	// slice point to array specific item (first number brackets, 0),
	// and slice have second attribute len is count of used items from slice
	// (second number in brackets, 4).
	// In our example sl slice slices 4 items begin from first item {1,2,3,4}.
	// Slice has third attribute capacity is count of array items
	// between slice point item and last array item.
	sl := arr[0:4]
	fmt.Printf("sl %v; sl p - %p; sl len - %v; sl cap - %v\n",
		sl, sl, len(sl), cap(sl))

	// we can take slice from any array elements, not only first
	// slice for third and fourth array items (by index 2 and 3)
	sl2 := arr[2:4]
	fmt.Printf("sl2 %v; sl2 p - %p; sl2 len - %v; sl2 cap - %v\n",
		sl2, sl2, len(sl2), cap(sl2))

	// we can take slice from slice
	sl3 := sl2[:]
	fmt.Printf("sl3 %v; sl3 p - %p; sl3 len - %v; sl3 cap - %v\n",
		sl3, sl3, len(sl3), cap(sl3))

	// add element to slice
	sl3 = append(sl3, 55)
	fmt.Printf("sl3 %v; sl3 p - %p; sl3 len - %v; sl3 cap - %v\n",
		sl3, sl3, len(sl3), cap(sl3))
	// add next element
	sl3 = append(sl3, 66)
	fmt.Printf("sl3 %v; sl3 p - %p; sl3 len - %v; sl3 cap - %v\n",
		sl3, sl3, len(sl3), cap(sl3))
	// add next element
	sl3 = append(sl3, 77)
	fmt.Printf("sl3 %v; sl3 p - %p; sl3 len - %v; sl3 cap - %v\n",
		sl3, sl3, len(sl3), cap(sl3))
	// add next element
	sl3 = append(sl3, 88)
	fmt.Printf("sl3 %v; sl3 p - %p; sl3 len - %v; sl3 cap - %v\n",
		sl3, sl3, len(sl3), cap(sl3))

	// add next element, out of capacity
	// add next element
	sl3 = append(sl3, 99)
	fmt.Printf("sl3 %v; sl3 p - %p; sl3 len - %v; sl3 cap - %v\n",
		sl3, sl3, len(sl3), cap(sl3))

}
