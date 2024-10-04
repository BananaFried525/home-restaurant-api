package helpers

import "fmt"

func PaddingIntToString(number int, length int, with string) string {
	str := fmt.Sprint(number)
	for len(str) < length {
		str = with + str
	}

	return str
}
