package main

import "fmt"
import "math"

func main() {
	fmt.Println(convert_to_secret(7))
	//convert_to_secret(2)
	//convert_to_secret(4)
	//convert_to_secret(8)
	//convert_to_secret(3)
	//convert_to_secret(19)
	//convert_to_secret(31)
	//convert_to_secret(0)
	//convert_to_secret(32)
	//convert_to_secret(33)

}

//1 = wink
//10 = double blink
//100 = close your eyes
//1000 = jump

func convert_to_secret(n int) [4]string {
	secret := map[int]string{
		1: "double blink",
		2: "close your eyes",
		3: "jump",
	}
	var result [4]string
	if n%2 != 0 {
		result[0] = "wink"
		n = n - 1
	}

	for i := 1; i < 4; i++ {
		tmp := int(math.Pow(2.0, float64(i)))
		if n%tmp == 0 && n != 0 {
			result[i] = secret[i]
			n = n - tmp
		}
	}
	//for i := 0; i < 4; i++ {
	//fmt.Println(result[i])
	//}

	// invert list (from stackoverflow)
	if n%16 != 0 {
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
	}
	return result
}
