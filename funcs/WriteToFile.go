// Function WriteToFile takes a pointer to a filename and data
// It writes the contents of the data in the file
package funcs

import (
	"os"
)

func WriteToFile(filename *os.File, data string) string {
	_, err := filename.WriteString(data)
	if err != nil {
		return "Internal Server error"
	}
	return ""
}
