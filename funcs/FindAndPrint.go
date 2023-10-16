//Function FindAndPrint() takes a bannerFileName and a slice of individual characters of type string
//The function will create the graphical representation of the characters and return it
//In case of any errors the function will return an empty string and an error msg

package funcs

func FindAndPrint(chars []string, bannerFile string) (string, string) {
	count := 0 //Count used to be added to the line number in order to get all 8 lines
	s := ""    //The value to be returned
	err := ""
	str := ""

	if chars != nil {
		for i := 1; i <= 8; i++ { //Loop to go through the all 8 lines
			for j := 0; j < len(chars); j++ { //Loop to go through all the characters
				//Build s by calling GetString() with the character and the count
				//GetString() will return the string found at (line num + count). line num will be returned by calling ReturnLineNum() in GetString()
				
				if (chars[j] < " " || chars[j] > "~") && chars[j] == "\r" {  //Ignore "/r"
					continue
				} else {
					//get the graphical representation and build s
					//Handle any errors
					str, err = GetString(chars[j], count, bannerFile)
					if err == "" {
						s = s + str
					} else {
						return "", err
					}
				}

			}
			//Add "/n" to s and increment the count
			s = s + "\n"
			count++
		}
	}
	return s, err
}
