package main

import "fmt"

func main() {

	//Print basic array in loop

	i := [4]string{"A", "B", "C", "D"}

	for a := 0; a < 4; a++ {

		fmt.Println(i[a])

	}

	//==========================================================================
	//length of an array

	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(len(arr))

	//===========================================================================
	//Array is value type not reference type

	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7}
	//print array before any change
	fmt.Println(arr1)

	arr2 := arr1

	//put a new value at oth position
	arr2[0] = 9

	fmt.Println(arr2)
	//=============================================================================

	//copy array by reference

	arrayref0 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arrayref0)

	arrayref1 := &arrayref0

	arrayref1[5] = 100

	fmt.Println(arr)

	//===========================================================================

	//array comparator

	array0 := [3]int{1, 2, 3}
	array1 := [...]int{4, 5, 6}
	array2 := [3]int{1, 2, 3}

	fmt.Println(array0 == array1)
	fmt.Println(array1 == array2)
	fmt.Println(array2 == array0)

}
