//Function MakeChars() takes a string value as argument and returns a slice of the passed string's characters
//The function will split each rune of the passed string and add it into a slice

package funcs

func MakeChars(st string) []string {
	var chars []string //Create an empty string slice

	for i := 0; i < len(st); i++ {
		chars = append(chars, string(st[i]))
	}

	return chars
}
