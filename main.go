package main
import ("fmt")

const PI = 3.14

func main() {
	var username string = "Eranga Madhushan"
    var userage int = 23

	var i string = "Hello"
	var j int = -15

	var arr1 = [5]int{1,2,3,4:72} // 5 th element is 72
	arr2 := [5]int{4,5,6,7,8}

	var arr3 = [...]int{1,2,3} // array length is inferred
	arr4 := [...]int{4,5,6,7,8}

	arr5 := [...]string{"Eranga", "Madhushan", "Kumara"}

	// Go slices
	 var s1 []int = arr1[1:4] // slicing array from index 1 to 3
	 s2 := arr2[2:5] // slicing array from index 2 to 4

	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println(arr3)
	fmt.Println(arr4)

	fmt.Println(arr5)

	fmt.Println(s1)
	fmt.Println(s2)

	// -------------------------------------------------------------------
	// slice examples
	myslice1 := []int{}
	fmt.Println(len(myslice1)) // length of the slice
	fmt.Println(cap(myslice1)) // capacity of the slice
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2)) // length of the slice
	fmt.Println(cap(myslice2)) // capacity of the slice
	fmt.Println(myslice2)

	fmt.Println("Current User Name : " + username + " and user age : " + fmt.Sprint(userage))

	inferradName := "Eranga Madhushan"
    inferradAge := 23

    fmt.Println(inferradName)
    fmt.Println(inferradAge)

	

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T\n", j, j)
}