package raindrops

import (
	"bytes"
	"strconv"
)

type raindrop struct {
	prime int
	sound string
}

func Convert(n int) string {
	var result bytes.Buffer

	drops := []raindrop{
		raindrop{prime: 3, sound: "Pling"},
		raindrop{prime: 5, sound: "Plang"},
		raindrop{prime: 7, sound: "Plong"},
	}

	for _, drop := range drops {
		if n%drop.prime == 0 {
			result.WriteString(drop.sound)
		}
	}

	if result.Len() == 0 {
		return strconv.Itoa(n)
	}
	return result.String()
}
