//Main ascii-art-web program
//Creates a server on port 8080 and listens for any requests
//Only serves the path "/" using a home handler function
//Handles client and server side errors using an error handler function
//Implements GET and POST HTTP Methods
//Executes the appropriate HTML file based on the request
//"index.html": Home and result page. Includes a form for data entry and textarea for the result
//"error.html": Error page. Shows the appropriate status and status code for the errors (400, 404, 500)
//Calls several functions from ascii-art-web/funcs to create the graphical representation of the passed string

package main

import (
	"ascii-art-web-export-file/funcs"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// Struct that will hold the form's data after the POST request
type Details struct {
	Str       string //The string to change
	Templatee string //The banner to use
}

// Struct to save the result after processing
// The data will be sent to the html files (index.html/error.html)
type Result struct {
	Status string //The message to be sent in case of an error (ex.Not Found)
	Code   int    //The status code (200,400,404,500)
	Output string //The result of changing the string
	Str    string //The string to place in the text input
}

// The html template that will be used to execut the files
var tpl *template.Template

func main() {
	var err error
	tpl, err = template.ParseGlob("templates/*.html") //Parsing all HTML files located in the templates Dir
	if err != nil {
		log.Fatal(err) //Handling errors
	}

	http.HandleFunc("/", HomeHandler)                                                          //Handling the path "/": Home page
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles")))) //Serving css files
	http.HandleFunc("/ascii-art", AsciiHandler)                                                //Handling the output path

	fmt.Println("Listening on port 8080")
	er := http.ListenAndServe(":8080", nil) //Creating a server at port 8080
	if er != nil {
		log.Fatal(er) //Handle any error
	}

}

// Handles the GET request on the route "/". Executes index.html with empty output
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	result := Result{} //Holds the status, code, and output (The Data)

	//Set the code to be 200 and status to "OK"
	result.Code = 200
	result.Status = "OK"

	if r.Method == http.MethodGet { //If the http method is GET, then execute the index.html file
		if r.URL.Path != "/" { //In case the URL path is anything other than "/", then call the ErrorHandler()
			//Set appropriate status and status code
			result.Code = 404
			result.Status = "Not Found"
			ErrorHandler(w, r, &result)
			return
		}

		result.Output = "" //Set the output to be empty
		result.Str = ""    //Set empty input

		//Execute the template and handle any errors
		er := tpl.ExecuteTemplate(w, "index.html", &result)
		if er != nil {
			result.Status = "Internal Server Error" //Any error with the template execution will be an internal error
			result.Code = 500
			ErrorHandler(w, r, &result)
			return
		}
	} else {
		//Set appropriate status and status code
		result.Code = 400
		result.Status = "Bad Request"
		ErrorHandler(w, r, &result)
		return
	}
}

// Handles the POST request from the form submition
func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	res := Result{}
	res.Code = 200
	res.Status = "OK"

	//If it is a POST method request
	if r.Method == http.MethodPost {
		if r.URL.Path != "/ascii-art" { //In case the URL path is anything other than "/ascii-art", then call the ErrorHandler()
			//Set appropriate status and status code
			res.Code = 400
			res.Status = "Bad Request"
			ErrorHandler(w, r, &res)
			return
		}
		//Parse the form and process the data and send it back
		err := r.ParseForm()
		if err != nil {
			res.Code = http.StatusInternalServerError
			res.Status = "Internal server error"
			ErrorHandler(w, r, &res)
			return
		}

		//Save the string to change and the banner to use
		details := Details{
			Str:       r.FormValue("str"),
			Templatee: r.FormValue("banner"),
		}
		st := details.Str
		templ := details.Templatee
		res.Str = st

		outputFile, err := os.Create("outputFile/output.txt")
		if err != nil {
			res.Status = "Internal Server Error"
			res.Code = 500
			ErrorHandler(w, r, &res)
			return
		}

		//Call the GetArt() function and handle any errors
		res.Output, res.Status = GetArt(st, templ)

		Werr := funcs.WriteToFile(outputFile, res.Output)
		if Werr != "" {
			res.Status = Werr
			res.Code = 500
			ErrorHandler(w, r, &res)
			return
		}

		if res.Status != "" {
			//Based on the error status save the appropriate code
			if res.Status == "Not Found" {
				res.Code = 404
			} else if res.Status == "Bad Request" {
				res.Code = 400
			} else {
				res.Code = http.StatusInternalServerError

			}
			//send the status code and call ErrorHandler to execute the error template
			ErrorHandler(w, r, &res)
			return
		}

		//In case there is no errors execute index.html with the result
		er := tpl.ExecuteTemplate(w, "index.html", &res)
		//Handle template execution error
		if er != nil {
			res.Status = "Internal Server Error"
			res.Code = 500
			ErrorHandler(w, r, &res)
			return
		}
	} else {
		//Set appropriate status and status code
		res.Code = 400
		res.Status = "Bad Request"
		ErrorHandler(w, r, &res)
		return
	}
}

// ErrorHandler() will send the appropriate error status code and will execute error.html
func ErrorHandler(w http.ResponseWriter, r *http.Request, res *Result) {
	w.WriteHeader(res.Code)
	err := tpl.ExecuteTemplate(w, "error.html", res)
	if err != nil {
		fmt.Println("Error with error.html")
		os.Exit(2)
	}
}

// GetArt() will return the ascii-art of the passed string
// In case of ay errors it will return an empty string and the error msg
func GetArt(st, templ string) (string, string) {
	var chars []string //String slice to store each character of the strings separately
	res := ""
	err := ""
	str := ""
	if st != "" {
		strings := strings.Split(st, "\n") //Splits the string if there exists a new line "\n" in it. Returns slice of strings
		for _, c := range strings {        //loop through the strings
			//Create a slice containing each character in the string passed
			//Find and print the graphical representaion of each of the characters
			if c == "\n" {
				res = res + "\n"
			} else {
				chars = funcs.MakeChars(c)
				str, err = funcs.FindAndPrint(chars, templ)
				if err == "" {
					res = res + str
				} else {
					return "", err
				}
			}
		}
	}
	return res, err
}
