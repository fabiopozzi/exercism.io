package secret

//1 = wink
//10 = double blink
//100 = close your eyes
//1000 = jump

func Handshake(n int) []string {
	secret := map[int]string{
		1: "wink",
		2: "double blink",
		3: "close your eyes",
		4: "jump",
	}
	result := []string(nil)

	if n < 0 {
		return nil
	}

	for i := 1; i < 5; i++ {
		if n%2 != 0 {
			result = append(result, secret[i])
		}
		n = n >> 1
	}

	// invert array (shameless copied from stackoverflow)
	if n != 0 {
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
	}
	return result
}
