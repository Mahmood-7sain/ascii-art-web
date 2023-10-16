//Function ReturnLineNum() takes a single character as a string and returns its starting line in the banner file

package funcs

func ReturnLineNum(character string) int {
	var char_array []rune    //Slice to store each character as a rune
	var line_Num_array []int //Slice to store the line number corresponding to each rune

	//Append the ascii characters represented by 32 (" ") to 126 ("~")
	for i := 32; i <= 126; i++ {
		char_array = append(char_array, rune(i))
	}

	i := 2 //The starting line of the first character in the banner file
	//Append the line numbers starting with 2 and incremented by 9 util the last line of the last character in the banner file
	for i < 855 {
		line_Num_array = append(line_Num_array, i)
		i = i + 9
	}

	//Loop through the char_array and look for the passed character in the function call
	//If found return the line number from line_Num_array at index i
	for i := 0; i < len(char_array); i++ {
		if character == string(char_array[i]) {
			return line_Num_array[i]
		}
	}

	if character == ""{
		return -2
	}
	return -1 //Return -1 if the character was not found
}
