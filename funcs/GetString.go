//Function GetString() receives a string s:(single character) and an integer count, and a file name and returns a string value
//The function will first get the starting line number of the passed character by calling ReturnLineNum(s)
//A scanner will be used to read the contents of the banner file provided
//The function will look for the string value found at (lineNum + count)
//The string value will be returned

package funcs

import (
	"bufio"
	"io"
	"os"
)

func GetString(s string, count int, fileName string) (string, string) {
	error := ""
	fileName = fileName + ".txt"
	st := ""
	line := ReturnLineNum(s) //Get the starting line number of s

	//-1 means invalid character
	if line == -1 {
		error = "Bad Request"
		return st, error
	} else if line == -2 { //-2 means a "/n"
		return "\n", error
	}
	path := "../ascii-art-web-export-file/banners/" + fileName //Saving the path of the file to read from
	file, err := os.Open(path)                                 //Open the banner file and handle errors when needed
	if err != nil {
		error = "Internal Server Error"
		return st, error
	}

	scanner := bufio.NewScanner(file)
	scan := 0 //Used to indicate the number of scans done

	for scanner.Scan() { //While the scanner is able to scan the banner file
		scan++
		if scan == line+count { //If lineNum + count match the # of scans ==> Correct line was found
			st = scanner.Text() //Save the contents of the line in st
		}
	}
	_, err = file.Seek(0, io.SeekStart) //Return the scanner to the beggining of the file
	if err != nil {
		error = err.Error()
		return "", error
	}

	return st, error
}
