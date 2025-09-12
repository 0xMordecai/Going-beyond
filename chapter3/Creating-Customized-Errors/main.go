package main

func main() {

}
func Demo() {
	// -->> Creating Customized Errors
	/*
		The errors.New function is very straightforward.

		-->			err := errors.New("Syntax error in the code")
	*/

	/*
	   The fmt.Errorf function, like many of the functions in the fmt
	   function, allows for formatting within the string.

	   -->			err := fmt.Errorf("Syntax error in the code at line %d", line)
	*/
}

// Implementing the Error() Interface
// let’s say you are writing a program used for communications and want to
// create a custom error to represent an error during communications.
type CommsError struct{}

func (m CommsError) Error() string {
	return "An error happened during data transfer."
}

// You want to provide information about where the error came from. To do
// this, you can create a custom error to provide the information.
